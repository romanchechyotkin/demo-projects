package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Debug("application started")

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("got request", slog.String("method", r.Method), slog.String("uri", r.RequestURI))
		now := time.Now()

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%d\n", now.Unix())
	})

	logger.Info("http server started", slog.String("port", "8000"))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
