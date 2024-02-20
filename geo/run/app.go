package run

import (
	"context"
	"os"
	"projects/LDmitryLD/hugoproxy-microservices/geo/config"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/db"
	cache "projects/LDmitryLD/hugoproxy-microservices/geo/internal/infrastructure/cahe"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/infrastructure/component"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/infrastructure/server"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/modules"
	"projects/LDmitryLD/hugoproxy-microservices/geo/internal/storages"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Runner interface {
	Run() error
}

type App struct {
	conf   config.AppConf
	logger *zap.Logger
	rpc    server.Server
	Sig    chan os.Signal
}

func NewApp(conf config.AppConf, logger *zap.Logger) *App {
	return &App{conf: conf, logger: logger, Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		sigInt := <-a.Sig
		a.logger.Info("signal interrupt recieved", zap.Stringer("os_signal", sigInt))
		cancel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.rpc.Serve(ctx)
		if err != nil {
			a.logger.Error("app: server error", zap.Error(err))
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) Bootstrap(options ...interface{}) Runner {

	_, sqlAdapter, err := db.NewSqlDB(a.conf.DB, a.logger)
	if err != nil {
		a.logger.Fatal("error init db:", zap.Error(err))
	}

	cacheClient, err := cache.NewRedisClient(a.conf.Cache.Host, a.conf.Cache.Port, a.logger)
	if err != nil {
		a.logger.Fatal("error init cache:", zap.Error(err))
	}

	components := component.NewComponents(a.conf, a.logger)

	newStorages := storages.NewStorages(sqlAdapter, cacheClient, a.logger)

	services := modules.NewServices(newStorages, components)

	a.rpc, err = server.GetServerRPC(a.conf.RPCServer, services.Geo, a.logger)
	if err != nil {
		a.logger.Fatal("error init rpc server:", zap.Error(err))
	}

	// a.rpc = server.NewServerGRPCTest(a.logger)

	return a
}
