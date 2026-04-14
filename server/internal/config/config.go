package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Thumbnail ThumbnailConfig
}

type ServerConfig struct {
	Port          int    `mapstructure:"port"`
	UploadDir     string `mapstructure:"upload_dir"`
	MaxUploadSize string `mapstructure:"max_upload_size"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type ThumbnailConfig struct {
	MaxWidth  int `mapstructure:"max_width"`
	MaxHeight int `mapstructure:"max_height"`
	Quality   int `mapstructure:"quality"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./server")

	viper.SetEnvPrefix("PAWPRINTS")
	viper.AutomaticEnv()

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.upload_dir", "./uploads")
	viper.SetDefault("server.max_upload_size", "10MB")
	viper.SetDefault("database.path", "./data/pawprints.db")
	viper.SetDefault("thumbnail.max_width", 400)
	viper.SetDefault("thumbnail.max_height", 400)
	viper.SetDefault("thumbnail.quality", 80)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
