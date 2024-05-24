package main

import (
	"github.com/stepanleas/notification-service/container"
	"go.uber.org/fx"
)

func main() {
	appContainer := fx.New(container.ApplicationModule)

	appContainer.Run()
}
