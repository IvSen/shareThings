package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/IvSen/shareThings/pkg/client/postgresql"

	"github.com/IvSen/shareThings/pkg/metric"

	"github.com/IvSen/shareThings/pkg/logging"

	"github.com/IvSen/shareThings/pkg/config"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"net"
	"net/http"
)

type App struct {
	logger logging.Logger
	cfg    *config.Config

	router     *httprouter.Router
	httpServer *http.Server

	pgClient *gorm.DB
}

func NewApp(ctx context.Context, config *config.Config, logger logging.Logger) (App, error) {
	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("heartbeat metric initializing")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	logger.Info("postgresql initializing")
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(pgConfig)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	//pgClient.AutoMigrate(&daoUser.User{})
	//pgClient.AutoMigrate(&daoAlbum.Album{})
	//pgClient.AutoMigrate(&daoPhoto.Photo{})

	return App{
		cfg:      config,
		pgClient: pgClient,
		router:   router,
		logger:   logger,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return a.startHTTP()
	})
	return group.Wait()
}

func (a *App) startHTTP() error {
	a.logger.WithFields(map[string]interface{}{
		"IP":   a.cfg.HTTP.IP,
		"Port": a.cfg.HTTP.Port,
	}).Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		a.logger.WithError(err).Fatal("failed to create listener")
	}

	a.logger.WithFields(map[string]interface{}{
		"AllowedMethods":     a.cfg.HTTP.CORS.AllowedMethods,
		"AllowedOrigins":     a.cfg.HTTP.CORS.AllowedOrigins,
		"AllowCredentials":   a.cfg.HTTP.CORS.AllowCredentials,
		"AllowedHeaders":     a.cfg.HTTP.CORS.AllowedHeaders,
		"OptionsPassthrough": a.cfg.HTTP.CORS.OptionsPassthrough,
		"ExposedHeaders":     a.cfg.HTTP.CORS.ExposedHeaders,
		"Debug":              a.cfg.HTTP.CORS.Debug,
	}).Info("cors.Options")
	c := cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warning("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}
	return err
}
