package app

import (
	"context"
	"time"

	"github.com/Mark1708/go-pastebin/internal/config"
	hh "github.com/Mark1708/go-pastebin/internal/handler/healthcheck"
	ph "github.com/Mark1708/go-pastebin/internal/handler/paste"
	pr "github.com/Mark1708/go-pastebin/internal/repository/paste"
	ps "github.com/Mark1708/go-pastebin/internal/service/paste"
	"github.com/Mark1708/go-pastebin/pkg/closer"
	"github.com/Mark1708/go-pastebin/pkg/db"
	"github.com/Mark1708/go-pastebin/pkg/db/pg"
	"github.com/Mark1708/go-pastebin/pkg/db/transaction"
	"github.com/Mark1708/go-pastebin/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig

	// db
	pgPoolConfig *pgxpool.Config
	dbClient     db.Client
	txManager    db.TxManager

	// paste
	pasteRepository pr.Repository
	pasteService    ps.Service
	pasteHandler    ph.Handler
	healthHandler   hh.Handler
}

func newServiceProvider() *serviceProvider {
	cfg, err := loggerConfig()
	if err == nil {
		logger.SetZapLogger(cfg)
	}
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() *pgxpool.Config {
	if s.pgPoolConfig == nil {
		if s.pgConfig == nil {
			cfg, err := config.NewPGConfig()
			if err != nil {
				logger.Log.With(zap.Error(err)).Fatal("failed to get pg config")
			}

			s.pgConfig = cfg
		}

		const defaultMaxConns = int32(4)
		const defaultMinConns = int32(0)
		const defaultMaxConnLifetime = time.Hour
		const defaultMaxConnIdleTime = time.Minute * 30
		const defaultHealthCheckPeriod = time.Minute
		const defaultConnectTimeout = time.Second * 5

		var err error
		s.pgPoolConfig, err = pgxpool.ParseConfig(s.pgConfig.DSN())
		if err != nil {
			logger.Log.With(zap.Error(err)).Fatal("failed to create a config, error")
		}

		s.pgPoolConfig.MaxConns = defaultMaxConns
		s.pgPoolConfig.MinConns = defaultMinConns
		s.pgPoolConfig.MaxConnLifetime = defaultMaxConnLifetime
		s.pgPoolConfig.MaxConnIdleTime = defaultMaxConnIdleTime
		s.pgPoolConfig.HealthCheckPeriod = defaultHealthCheckPeriod
		s.pgPoolConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
	}

	return s.pgPoolConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			logger.Log.With(zap.Error(err)).Fatal("failed to get http config")
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := pg.New(ctx, s.PGConfig())
		if err != nil {
			logger.Log.With(zap.Error(err)).Fatal("failed to create db client")
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			logger.Log.With(zap.Error(err)).Fatal("ping db error")
		}

		closer.Add(client.Close)
		s.dbClient = client
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) PasteRepository(ctx context.Context) pr.Repository {
	if s.pasteRepository == nil {
		s.pasteRepository = pr.NewRepository(
			s.DBClient(ctx),
			s.TxManager(ctx),
		)
	}
	return s.pasteRepository
}

func (s *serviceProvider) PasteService(ctx context.Context) ps.Service {
	if s.pasteService == nil {
		s.pasteService = ps.NewService(s.PasteRepository(ctx))
	}
	return s.pasteService
}

func (s *serviceProvider) PasteHandler(ctx context.Context) ph.Handler {
	if s.pasteHandler == nil {
		s.pasteHandler = ph.NewHandler(s.PasteService(ctx))
	}
	return s.pasteHandler
}

func (s *serviceProvider) HealthHandler() hh.Handler {
	if s.healthHandler == nil {
		s.healthHandler = hh.NewHandler()
	}
	return s.healthHandler
}

func loggerConfig() (config.LoggerConfig, error) {
	cfg, err := config.NewLoggerConfig()
	if err != nil {
		logger.Log.With(zap.Error(err)).Warn("failed to get logger config")
		return nil, err
	}
	return cfg, nil
}
