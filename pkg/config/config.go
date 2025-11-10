package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string        `yaml:"env" env-required:"true"`
	GRPC    GRPCConfig    `yaml:"grpc"`
	Storage StorageConfig `yaml:"storage"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type StorageConfig struct {
	Driver   string        `yaml:"driver"`
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	User     string        `yaml:"user"`
	Password string        `yaml:"password"`
	DbName   string        `yaml:"dbname"`
	DSN      string        `yaml:"dsn"`
	Timeout  time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is not set")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() (path string) {
	flag.StringVar(&path, "config", "", "Path to config file")
	flag.Parse()

	if path == "" {
		panic("config path is not set")
	}

	return
}
