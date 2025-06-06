package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Client interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Close()
}

type Cfg struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
	Listen       struct {
		BindIP string `env:"LISTEN_BIND_IP" env-default:"127.0.0.1"`
		Port   string `env:"LISTEN_PORT" env-default:"8080"`
	}
	Postgresql struct {
		Host     string `env:"POSTGRESQL_HOST"`
		Port     string `env:"POSTGRESQL_PORT"`
		Database string `env:"POSTGRESQL_DATABASE"`
		Username string `env:"POSTGRESQL_USERNAME"`
		Password string `env:"POSTGRESQL_PASSWORD"`
	}
}

var instance *Cfg
var once sync.Once

func GetConfigEnv() *Cfg {
	once.Do(func() {
		instance = &Cfg{}
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error open .env file")
		}
		err = cleanenv.ReadConfig(".env", instance)
		if err != nil {
			cleanenv.GetDescription(instance, nil)
		}
	})
	return instance
}

func ConnectDB(ctx context.Context) (pool *pgxpool.Pool, err error) {
	cfg := GetConfigEnv()
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Postgresql.Username, cfg.Postgresql.Password, cfg.Postgresql.Host, cfg.Postgresql.Port, cfg.Postgresql.Database)
	pool, err = pgxpool.New(ctx, dns)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %v", err)
	}
	return pool, nil
}
