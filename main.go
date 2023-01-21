package main

import (
	"context"
	"fmt"

	"yanita_inventario/config"
	"yanita_inventario/database"
	"yanita_inventario/internal/api"
	"yanita_inventario/internal/repository"
	"yanita_inventario/internal/service"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	application := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)

	application.Run()

}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *config.Config, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
