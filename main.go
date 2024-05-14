package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ngobrut/halo-sus-api/config"
	"github.com/ngobrut/halo-sus-api/database"
	http_handler "github.com/ngobrut/halo-sus-api/internal/handler"
	"github.com/ngobrut/halo-sus-api/internal/repository"
	"github.com/ngobrut/halo-sus-api/internal/usecase"
	"github.com/sirupsen/logrus"
)

const (
	addr = ":8080"
)

func exec() error {
	cnf := config.New()
	logger := logrus.New()

	db, err := database.NewDBClient(cnf, logger)
	if err != nil {
		return err
	}

	repo := repository.New(cnf, db)
	uc := usecase.New(cnf, db, repo)
	handler := http_handler.InitHTTPHandler(cnf, uc)

	httpServer := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 90 * time.Second,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// implement graceful shutdown
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Printf("[graceful-shutdown-time-out] \n%v\n", err.Error())
			}
		}()

		defer cancel()

		log.Println("graceful shutdown.....")

		// trigger graceful shutdown
		err = httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("[graceful-shutdown-error] \n%v\n", err.Error())
		}

		serverStopCtx()
	}()

	// run server
	log.Printf("[http-server-online] %v\n", "http://localhost:8080")

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("[http-server-failed] \n%v\n", err.Error())
		return err
	}

	<-serverCtx.Done()

	return nil
}

func main() {
	if err := exec(); err != nil {
		log.Fatal("[app-failed]", err)
	}
}
