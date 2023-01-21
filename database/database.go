package database

import (
	"context"
	"fmt"
	"yanita_inventario/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *config.Config) (*sqlx.DB, error) {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	return sqlx.ConnectContext(ctx, "mysql", connection)
}
