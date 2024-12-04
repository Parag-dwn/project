package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Parag-dwn/student-api/internal/config"
	"github.com/Parag-dwn/student-api/internal/http/handlers/student"
	"github.com/Parag-dwn/student-api/internal/storage/sqlite"
)

func main() {

	//Load Config

	cfg := config.MustLoad()

	// database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initalized ", slog.String(("env"), cfg.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GETById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("server Started %s", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("SErver Started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
