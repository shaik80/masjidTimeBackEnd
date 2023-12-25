package config

import (
	"sync"

	"github.com/spf13/viper"
)

// AppConfig represents the structure of the config.yaml file
type AppConfig struct {
	App      AppConfigData  `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// AppConfigData represents application-level configuration
type AppConfigData struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
}

// LoggingConfig represents the logging configuration
type LoggingConfig struct {
	Level string      `mapstructure:"level"`
	File  LoggingFile `mapstructure:"file"`
}

// LoggingFile represents the file logging configuration
type LoggingFile struct {
	Enabled  bool   `mapstructure:"enabled"`
	Filename string `mapstructure:"filename"`
}

var (
	appConfig     AppConfig
	appConfigOnce sync.Once
)

// Load initializes the configuration once
func Load() error {
	var err error
	appConfigOnce.Do(func() {
		err = loadConfig()
	})
	return err
}

// loadConfig loads the configuration from the file
func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set default values
	viper.SetDefault("app.name", "MyFiberApp")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("app.debug", true)
	viper.SetDefault("server.address", "localhost")
	viper.SetDefault("server.port", 3000)
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.dbname", "mydatabase")
	viper.SetDefault("database.maxIdleConns", 10)
	viper.SetDefault("database.maxOpenConns", 100)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.file.enabled", true)
	viper.SetDefault("logging.file.filename", "app.log")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return err
	}

	return nil
}

// GetConfig returns the loaded configuration
func GetConfig() AppConfig {
	return appConfig
}

// GetString gets a string configuration value by key
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt gets an integer configuration value by key
func GetInt(key string) int {
	return viper.GetInt(key)
}
