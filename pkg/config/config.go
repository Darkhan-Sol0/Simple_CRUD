package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type PgCfg struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	Database string `env:"PG_DATABASE"`
	Username string `env:"PG_USERNAME"`
	Password string `env:"PG_PASSWORD"`
}

type JwtCfg struct {
	JWTKey string `env:"JWT_KEY"`
}

var (
	jwtConf *JwtCfg
	pgConf  *PgCfg
	once    sync.Once
)

func parsConf() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: error loading .env file: ", err)
	}

	pgConf = &PgCfg{}
	if err = cleanenv.ReadConfig(".env", pgConf); err != nil {
		log.Println("Warning: error reading PostgreSQL config: ", err)
		cleanenv.GetDescription(pgConf, nil)
	}

	jwtConf = &JwtCfg{}
	if err = cleanenv.ReadConfig(".env", jwtConf); err != nil {
		log.Println("Warning: error reading JWT config: ", err)
		cleanenv.GetDescription(jwtConf, nil)
	}
}

func GetPgEnv() *PgCfg {
	once.Do(func() {
		parsConf()
	})
	return pgConf
}

func GetJwtEnv() *JwtCfg {
	once.Do(func() {
		parsConf()
	})
	return jwtConf
}
