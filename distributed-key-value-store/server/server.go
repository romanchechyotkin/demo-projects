package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"kv-store/store"
)

type handler struct {
	store store.Store
}

func newHandler(store store.Store) *handler {
	return &handler{store: store}
}

func RunServer(ctx context.Context, store store.Store) error {
	r := mux.NewRouter()

	h := newHandler(store)
	h.registerRoutes(r)

	srv := &http.Server{
		Addr: ":5000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()
	<-ctx.Done()

	log.Println("shutting down server gracefully")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	longShutdown := make(chan struct{}, 1)

	go func() {
		time.Sleep(3 * time.Second)
		longShutdown <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("server shutdown: %w", ctx.Err())
	case <-longShutdown:
		err := os.Remove("read.log")
		log.Println(err)
		log.Println("finished")
	}

	return nil
}

func (h *handler) registerRoutes(r *mux.Router) {
	r.HandleFunc("/v1/key/{key}", h.addRecord).Methods("POST")
	r.HandleFunc("/v1/key/{key}", h.getRecord).Methods("GET")
	r.HandleFunc("/v1/key/{key}", h.deleteRecord).Methods("DELETE")
}	

func (h *handler) addRecord(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.store.AddRecord(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	h.store.WritePostLog(key, string(value))

	w.WriteHeader(http.StatusCreated) 
	w.Write(value)
}

func (h *handler) getRecord(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := h.store.GetRecord(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Write([]byte(value))
}

func (h *handler) deleteRecord(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)
	vars := mux.Vars(r)
	key := vars["key"]

	err := h.store.DeleteRecord(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	h.store.WriteDeleteLog(key)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("deleted"))
}