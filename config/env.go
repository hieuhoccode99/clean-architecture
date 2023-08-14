package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"DB_HOST"`
	DBPort     string `env:"DB_PORT" envDefault:"DB_PORT"`
	DBUser     string `env:"DB_USER" envDefault:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"DB_NAME"`
	Port       string `env:"PORT" envDefault:"PORT"`

	MaxOpenConns    int `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`     // số lượng kết nối mở tối đa đến database
	MaxIdleConns    int `env:"DB_MAX_IDLE_CONNS" envDefault:"25"`     // số lượng kết nối tối đa ứng dụng có thể giữ khi không sử dụng
	ConnMaxLifetime int `env:"DB_CONN_MAX_LIFETIME" envDefault:"600"` // thời gian tối đa một kết nối có thể tồn tại
}

var cfg Config

func SetEnv() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to read environment variables: %v", err)
		return
	}
}

func GetEnv() Config {
	return cfg
}
