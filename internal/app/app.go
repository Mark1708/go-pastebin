package app

import (
	"context"
	"flag"
	"net"
	"net/http"
	"sync"

	"github.com/Mark1708/go-pastebin/internal/config"
	"github.com/Mark1708/go-pastebin/internal/router"
	"github.com/Mark1708/go-pastebin/pkg/closer"
	"github.com/Mark1708/go-pastebin/pkg/logger"
	"go.uber.org/zap"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		logger.Log.With(zap.Error(err)).Error("failed to initialize deps")
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			logger.Log.With(zap.Error(err)).Fatal("failed to run HTTP server")
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envPath := flag.String("env-file", "", "path to file with env")
	flag.Parse()

	if len(*envPath) > 0 {
		err := config.Load(*envPath)
		if err != nil {
			logger.Log.With(zap.Error(err)).Error("env file not found")
		}
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	cfg := a.serviceProvider.HTTPConfig()

	r := router.New(a.serviceProvider.PasteHandler(ctx), a.serviceProvider.HealthHandler())

	a.httpServer = &http.Server{
		Addr:              cfg.Address(),
		Handler:           r,
		ReadTimeout:       cfg.ReadTimeout(),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout(),
		WriteTimeout:      cfg.WriteTimeout(),
		IdleTimeout:       cfg.IdleTimeout(),
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Log.Infof("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()
	if err != nil {
		logger.Log.With(zap.Error(err)).Error("error of running http server")
		return err
	}

	return nil
}
