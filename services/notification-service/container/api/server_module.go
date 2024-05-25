package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stepanleas/notification-service/bootstrap"
	"go.uber.org/fx"
)

var serverModule = fx.Module("api-server",
	fx.Provide(newHttpServer),
	fx.Invoke(httpServerHandler),
	fx.Provide(newRouter),
)

func newRouter() *gin.Engine {
	return gin.Default()
}

func newHttpServer(r *gin.Engine, app bootstrap.Application) *http.Server {
	srv := &http.Server{
		Addr:              app.Env.ServerAddress,
		Handler:           r,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	return srv
}

func httpServerHandler(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.ListenAndServe()
				if err != nil {
					panic(err)
				}
			}()

			fmt.Println("Starting HTTP server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			go func() {
				err := srv.Shutdown(ctx)
				if err != nil {
					panic(err)
				}
			}()

			fmt.Println("Shutting HTTP server")
			return nil
		},
	})
}
