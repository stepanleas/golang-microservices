package api

import (
	"go.uber.org/fx"
)

var ApiModule = fx.Module("api",
	serverModule,
	healthCheckModule,
)
