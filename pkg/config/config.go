package config

import "strings"

type Config struct {
	SecretKey     string `mapstructure:"SECRET_KEY"`
	DatabasePath  string `mapstructure:"DATABASE_LOCATION"`
	RootFolder    string `mapstructure:"ROOT_FOLDER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
}

func (c *Config) Apply(appPath string) {
	c.DatabasePath = strings.ReplaceAll(c.DatabasePath, "{{appPath}}", appPath)
	c.RootFolder = strings.ReplaceAll(c.RootFolder, "{{appPath}}", appPath)
}
