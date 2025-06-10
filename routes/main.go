package routes

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/dhruvvadoliya1/movie-app-backend/config"
	"github.com/dhruvvadoliya1/movie-app-backend/constants"
	controller "github.com/dhruvvadoliya1/movie-app-backend/controllers/api/v1"
	"github.com/dhruvvadoliya1/movie-app-backend/middlewares"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/events"
	pMetrics "github.com/dhruvvadoliya1/movie-app-backend/pkg/prometheus"
	"github.com/dhruvvadoliya1/movie-app-backend/pkg/watermill"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

// Setup func
func Setup(app *fiber.App, goqu *goqu.Database, logger *zap.Logger, config config.AppConfig, events *events.Events, pMetrics *pMetrics.PrometheusMetrics, pub *watermill.WatermillPublisher) error {
	mu.Lock()

	app.Use(middlewares.LogHandler(logger, pMetrics))

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./assets/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	router := app.Group("/api")
	v1 := router.Group("/v1")

	middlewares := middlewares.NewMiddleware(config, logger)

	err := setupAuthController(v1, goqu, logger, middlewares, config)
	if err != nil {
		return err
	}

	err = setupUserController(v1, goqu, logger, middlewares, events, pub)
	if err != nil {
		return err
	}


	err = setupMovieController(v1, goqu, logger)
	if err != nil {
		return err
	}

	err = healthCheckController(app, goqu, logger)
	if err != nil {
		return err
	}

	err = metricsController(app, goqu, logger, pMetrics)
	if err != nil {
		return err
	}

	mu.Unlock()
	return nil
}

func setupAuthController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, config config.AppConfig) error {
	authController, err := controller.NewAuthController(goqu, logger, config)
	if err != nil {
		return err
	}
	v1.Post("/login", authController.DoAuth)

	if config.Kratos.IsEnabled {
		kratos := v1.Group("/kratos")
		kratos.Get("/auth", middlewares.Authenticated, authController.DoKratosAuth)
	}
	return nil
}

func setupUserController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger, middlewares middlewares.Middleware, events *events.Events, pub *watermill.WatermillPublisher) error {
	userController, err := controller.NewUserController(goqu, logger, events, pub)
	if err != nil {
		return err
	}

	userRouter := v1.Group("/users")
	userRouter.Post("/", userController.CreateUser)
	userRouter.Get(fmt.Sprintf("/:%s", constants.ParamUid), middlewares.Authenticated, userController.GetUser)
	return nil
}

func setupMovieController(v1 fiber.Router, goqu *goqu.Database, logger *zap.Logger) error {

	movieController, err := controller.NewMovieController(goqu, logger)
	if err != nil {
		return err
	}

	movieRouter := v1.Group("/movies")
	movieRouter.Post("/", movieController.CreateMovie)
	movieRouter.Get("/", movieController.GetMovies)
	movieRouter.Get(fmt.Sprintf("/:%s", constants.ParamMid), movieController.GetMovie)
	movieRouter.Put(fmt.Sprintf("/:%s", constants.ParamMid), movieController.UpdateMovie)
	movieRouter.Delete(fmt.Sprintf("/:%s", constants.ParamMid), movieController.DeleteMovie)
	return nil
}


func healthCheckController(app *fiber.App, goqu *goqu.Database, logger *zap.Logger) error {
	healthController, err := controller.NewHealthController(goqu, logger)
	if err != nil {
		return err
	}

	healthz := app.Group("/healthz")
	healthz.Get("/", healthController.Overall)
	healthz.Get("/db", healthController.Db)
	return nil
}

func metricsController(app *fiber.App, db *goqu.Database, logger *zap.Logger, pMetrics *pMetrics.PrometheusMetrics) error {
	metricsController, err := controller.InitMetricsController(db, logger, pMetrics)
	if err != nil {
		return nil
	}

	app.Get("/metrics", metricsController.Metrics)
	return nil
}
