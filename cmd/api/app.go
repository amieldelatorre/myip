package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amieldelatorre/myip/handler"
	"github.com/amieldelatorre/myip/service"
	"github.com/amieldelatorre/myip/utils"
)

type App struct {
	Server *http.Server
	Logger utils.CustomJsonLogger
}

func NewApp() App {
	logger := utils.NewCustomJsonLogger(os.Stdout, slog.LevelDebug)

	mux := http.NewServeMux()

	middleware := handler.NewMiddleware(logger)
	ipInfoService := service.NewIpInfoService(logger)
	ipInfoHandler := handler.NewIpInfoHandler(logger, ipInfoService)

	RegisterRoutes(mux, middleware, ipInfoHandler)

	app := App{
		Logger: logger,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
	return app
}

func (a *App) Exit() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	a.Logger.Info(ctx, "Exiting application...")

	a.Logger.Info(ctx, "Shutting down server")
	err := a.Server.Shutdown(ctx)
	if err != nil {
		a.Logger.Error(ctx, "Error shutting down server", "error", err)
	}

	a.Logger.Info(ctx, "Application has been shutdown, bye bye !")
}

func (a *App) Run() {
	ctx := context.Background()
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		a.Logger.Info(ctx, "Attempting to start application...")
		a.Logger.Info(ctx, fmt.Sprintf("Starting application on port %s", a.Server.Addr))
		err := a.Server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.Logger.Error(ctx, "Something went wrong with the server", "error", err)
		}
	}()

	sig := <-stopChan

	a.Logger.Info(ctx, fmt.Sprintf("Received signal '%+v', attempting to shutdown", sig))
	a.Exit()
}
