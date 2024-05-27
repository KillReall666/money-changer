package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/handlers/calculation"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {

	cfg, err := config.New("INFO")
	if err != nil {
		panic("failed to load cfg: " + err.Error())
	}
	log := logger.New(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	serv := service.New(cfg, log)

	r := chi.NewRouter()
	r.Post("/calculation", calculation.NewCalculationHandler(serv, log).GetCalculation)

	srv := &http.Server{
		Addr:    cfg.Address + cfg.Port,
		Handler: r,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.LogFatal("starting http server error: ", err)
		}
	}()

	log.LogInfo("starting http server to serve metrics on port: " + cfg.Port)

	<-ctx.Done()

	log.LogInfo("graceful shutdown of http server")

}
