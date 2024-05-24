package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stepanleas/notification-service/api/controller"
	"github.com/stepanleas/notification-service/pkg/logger"
	"go.uber.org/fx"
)

var healthCheckModule = fx.Module("api-healthcheck-module",
	fx.Provide(provideHealthCheckController),
	fx.Invoke(healthcheckRoutes),
)

func provideHealthCheckController(logger *logger.LogrusLogger) controller.HealthCheckController {
	return controller.NewHealthCheckController(logger)
}

func healthcheckRoutes(
	router *gin.Engine,
	ctrl controller.HealthCheckController,
) {
	router.GET("/notification-health", ctrl.HealthCheck)
}
