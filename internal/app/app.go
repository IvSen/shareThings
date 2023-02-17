package app

import (
	"context"
	"errors"
	"fmt"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/IvSen/shareThings/internal/composites"

	"github.com/IvSen/shareThings/pkg/metric"

	"github.com/IvSen/shareThings/pkg/logging"

	"github.com/IvSen/shareThings/pkg/config"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"golang.org/x/sync/errgroup"

	"net"
	"net/http"
)

type App struct {
	logger logging.Logger
	cfg    *config.Config

	router     *httprouter.Router
	httpServer *http.Server

	//pgClient *gorm.DB
}

func NewApp(ctx context.Context, config *config.Config, logger logging.Logger) (App, error) {
	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info(ctx, "swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info("heartbeat metric initializing")
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	logger.Info("postgresql composite initializing")
	pgClient, err := composites.NewPgClientComposite(config)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	//pgClient.AutoMigrate(&daoUser.User{})
	//pgClient.AutoMigrate(&daoAlbum.Album{})
	//pgClient.AutoMigrate(&daoPhoto.Photo{})

	logger.Println("cache composite initializing")
	// TODO: вынести в конфиг размер кеша
	cacheComposite, err := composites.NewCacheComposite(104857600) // 100MB
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	logger.Info("jwt composite initializing")
	jwtComposite, err := composites.NewJWTComposite(cacheComposite, logger)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	logger.Info("auth composite initializing")
	authComposite, err := composites.NewAuthComposite(pgClient, jwtComposite)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}
	authComposite.AuthHandler.Register(router)

	logger.Info("user composite initializing")
	userComposite, err := composites.NewUserComposite(pgClient, jwtComposite)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}
	userComposite.UserHandler.Register(router)

	logger.Info("gander composite initializing")
	genderComposite, err := composites.NewGenderComposite(pgClient, jwtComposite)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}
	genderComposite.GenderHandler.Register(router)

	logger.Info("child composite initializing")
	childComposite, err := composites.NewChildComposite(pgClient, jwtComposite)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}
	childComposite.ChildHandler.Register(router)

	//auth.NewHandler()

	return App{
		cfg: config,
		//pgClient: pgClient.Db,
		router: router,
		logger: logger,
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
