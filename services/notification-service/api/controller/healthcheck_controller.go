package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stepanleas/notification-service/pkg/logger"
)

type HealthCheckController struct {
	logger *logger.LogrusLogger
}

func NewHealthCheckController(logger *logger.LogrusLogger) HealthCheckController {
	return HealthCheckController{logger: logger}
}

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func (ctrl HealthCheckController) HealthCheck(c *gin.Context) {
	ctrl.logger.Info("Print a message here to kibana!")
	c.JSON(http.StatusOK, HealthCheckResponse{Message: "Notification service healthcheck is working!"})
}
