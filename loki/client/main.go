package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	timeServerURL, ok := os.LookupEnv("ADDR")
	if !ok {
		logger.Error("missing env variable ADDR")
		os.Exit(1)
	}

	getTimeURL := fmt.Sprintf("%s/time", timeServerURL)

	logger.Debug("application started")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("got request", slog.String("method", r.Method), slog.String("uri", r.RequestURI))

		logger.Debug("made request", slog.String("url", getTimeURL))

		resp, err := http.DefaultClient.Get(getTimeURL)
		if err != nil {
			logger.Error("failed to get response", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Time server (%s) is not available: %v", timeServerURL, err)
			return
		}
		defer resp.Body.Close()

		respTime, _ := io.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World at %s", string(respTime))
	})

	logger.Info("http server started", slog.String("port", "8001"))
	if err := http.ListenAndServe(":8001", nil); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
