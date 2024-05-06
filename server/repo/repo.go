package repo

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log/slog"
)

type TaskXRepository struct {
	DB     *sqlx.DB
	Config *Config
}

type Config struct {
	RootPassword string
}

func NewConfig() *Config {

	viper.AutomaticEnv()

	return &Config{
		RootPassword: viper.GetString("DB_ROOT_PASSWORD"),
	}
}

func NewTaskXRepository(lc fx.Lifecycle, cfg *Config) *TaskXRepository {

	db, err := sqlx.Connect("mysql", fmt.Sprintf("root:%s@(taskx-db:3306)/taskx-db", cfg.RootPassword))
	if err != nil {
		return nil
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			slog.Info("Gracefully closing DB connection")

			_ = db.Close()

			return nil
		},
	})

	return &TaskXRepository{
		DB:     db,
		Config: cfg,
	}
}
