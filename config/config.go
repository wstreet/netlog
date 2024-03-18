package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Config struct {
	ConfigPath string `default:"./config"`
	Env        string `default:"dev"`
	Sqlite     Sqlite
	Log        LogConfig
	Server     ServerConfig
}

type LogConfig struct {
	Path  string
	Level string
}

type ServerConfig struct {
	Mode string
	Port string
}

type Sqlite struct {
	Filename string
}

var c *Config

func GetConfig() Config {
	return *c
}

func init() {
	c = NewConfig()
}

func NewConfig() *Config {
	c := &Config{}
	defaults.SetDefaults(c)

	viper.SetConfigFile(c.configPathFile())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.Unmarshal(&c)
	return c
}

func (c Config) configPathFile() string {
	if !filepath.IsAbs(c.ConfigPath) {
		prefix, _ := os.Getwd()
		c.ConfigPath = filepath.Join(prefix, c.ConfigPath)
	}
	return filepath.Join(c.ConfigPath, fmt.Sprintf("config.%s.toml", c.Env))
}
