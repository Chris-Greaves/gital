package core

import (
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

var (
	KeyDirectories  = "directories"
	KeyScanDelay    = "scan_delay"
	KeyDatabasePath = "database_location"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	// paths to look for the config file in
	viper.AddConfigPath("$HOME/.gital")  // Linux or Mac home directory
	viper.AddConfigPath("$HOME\\.gital") // Windows home directoy
	viper.AddConfigPath("/etc/gital")    // Alternate Linux config location
	viper.AddConfigPath(".")             // Local config, useful for testing

	viper.SetDefault(KeyScanDelay, time.Second*30)
	viper.SetDefault(KeyDatabasePath, userHomeDir()+string(os.PathSeparator)+".gital")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("Config file changed", slog.String("filepath", e.Name))
	})
	viper.WatchConfig()

	return &Config{viper.GetViper()}, nil
}

func (c *Config) IsSet(key string) bool                { return c.v.IsSet(key) }
func (c *Config) GetString(key string) string          { return c.v.GetString(key) }
func (c *Config) GetStringSlice(key string) []string   { return c.v.GetStringSlice(key) }
func (c *Config) GetInt(key string) int                { return c.v.GetInt(key) }
func (c *Config) GetDuration(key string) time.Duration { return c.v.GetDuration(key) }

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
