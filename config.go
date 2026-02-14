package main

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
	Path string `yaml:"base_path" mapstructure:"base_path"`
	User string `yaml:"user" mapstructure:"user"`
	Pwd  string `yaml:"pwd" mapstructure:"pwd"`
}

type DriveConfig struct {
	UID         string `yaml:"uid" mapstructure:"uid"`
	CID         string `yaml:"cid" mapstructure:"cid"`
	SEID        string `yaml:"seid" mapstructure:"seid"`
	KID         string `yaml:"kid" mapstructure:"kid"`
	Rate        int    `yaml:"rate" mapstructure:"rate"`
	CacheExpire int    `yaml:"cache_expire" mapstructure:"cache_expire"`
}

type Config struct {
	Server ServerConfig `yaml:"server" mapstructure:"server"`
	Drive  DriveConfig  `yaml:"drive" mapstructure:"drive"`
}

func loadConfig(configPath string) (*Config, error) {
	var err error

	if configPath == "" {
		configPath, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	if _, err = os.Stat(path.Join(configPath, ".env")); err == nil {
		if err = godotenv.Load(); err != nil {
			return nil, err
		}
		slog.Debug(".env loaded")
	}

	v := viper.New()

	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8090)
	v.SetDefault("server.path", "/dav")
	v.SetDefault("server.user", "user")
	v.SetDefault("server.pwd", "password")
	v.SetDefault("drive.uid", "")
	v.SetDefault("drive.cid", "")
	v.SetDefault("drive.seid", "")
	v.SetDefault("drive.kid", "")
	v.SetDefault("drive.rate", 3)
	v.SetDefault("drive.cache_expire", 1)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, err
		}
	}

	conf := &Config{}

	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
