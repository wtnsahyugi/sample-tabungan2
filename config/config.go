package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	DatabaseHost      string `env:"DATABASE_HOST,default=localhost"`
	DatabasePort      string `env:"DATABASE_PORT,default=5433"`
	DatabaseUsername  string `env:"DATABASE_USERNAME,required"`
	DatabasePassword  string `env:"DATABASE_PASSWORD,required"`
	DatabaseName      string `env:"DATABASE_NAME,required"`
	RedisPort         string `env:"REDIS_PORT,default=:6378"`
	WorkerNamespace   string `env:"WORKER_NAMESPACE,default=tabungan-api"`
	WorkerConcurrency uint   `env:"WORKER_CONCURRENCY,default=10"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
