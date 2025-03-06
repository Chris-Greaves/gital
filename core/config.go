package core

import (
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	// paths to look for the config file in
	viper.AddConfigPath("$HOME/.gital")  // Linux or Mac home directory
	viper.AddConfigPath("$HOME\\.gital") // Windows home directoy
	viper.AddConfigPath("/etc/gital")    // Alternate Linux config location
	viper.AddConfigPath(".")             // Local config, useful for testing

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("Config file changed", slog.String("filepath", e.Name))
	})
	viper.WatchConfig()

	return &Config{viper.GetViper()}, nil
}

func (c *Config) IsSet(key string) bool              { return c.v.IsSet(key) }
func (c *Config) GetString(key string) string        { return c.v.GetString(key) }
func (c *Config) GetStringSlice(key string) []string { return c.v.GetStringSlice(key) }
func (c *Config) GetInt(key string) int              { return c.v.GetInt(key) }
