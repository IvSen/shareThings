package config

import (
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `yaml:"is-debug" env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `yaml:"is-development" env:"IS_DEV" env-default:"false"`
	HTTP          struct {
		IP           string        `yaml:"ip" env:"HTTP-IP"`
		Port         int           `yaml:"port" env:"HTTP-PORT"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP-READ-TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP-WRITE-TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed-methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed-origins"`
			AllowCredentials   bool     `yaml:"allow-credentials"`
			AllowedHeaders     []string `yaml:"allowed-headers"`
			OptionsPassthrough bool     `yaml:"options-passthrough"`
			ExposedHeaders     []string `yaml:"exposed-headers"`
			Debug              bool     `yaml:"debug"`
		} `yaml:"cors"`
	} `yaml:"http"`
	AppConfig struct {
		LogLevel string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
	} `yaml:"app"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
	// TODO: перенести в раздел Security
	JWT struct {
		Secret string `yaml:"secret" env-required:"true"`
	} `yaml:"JWT"`
	Security struct {
		Salt string `yaml:"salt" env-required:"true"`
	} `yaml:"security"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "configs/config.local.yaml", "this is app config file")
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "IvSen - shareThings"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
