package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hexisa_go_nal_todo/internal/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	once sync.Once
	db   *bun.DB
)

func GetDB() *bun.DB {
	return db
}

func setDB(d *bun.DB) {
	db = d
}

func NewDB(ctx context.Context, cfg config.DB) func() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			panic(err)
		}

		bundb := bun.NewDB(db, mysqldialect.New())

		bundb.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(false),
			bundebug.FromEnv("BUNDEBUG")),
		)

		setDB(bundb)
	})

	return func() { _ = db.Close() }
}
