package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	viper *viper.Viper
}

func init() {
	Config.viper = viper.GetViper()

	viper.SetConfigType("yaml")
	viper.SetConfigName("trading")
	viper.AddConfigPath("/etc/trading/")
	viper.AddConfigPath("$HOME/.trading")
	viper.AddConfigPath(".")

	viper.SetDefault("Port", "8080")
	viper.SetDefault("BaseURL", "/v1")
	viper.SetDefault("LogLevel", "trace")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

var Config config

func Get(key string) interface{} { return Config.Get(key) }
func (c *config) Get(key string) interface{} {
	return c.viper.Get(key)
}

func GetBool(key string) bool { return Config.GetBool(key) }
func (c *config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func GetFloat64(key string) float64 { return Config.GetFloat64(key) }
func (c *config) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

func GetInt(key string) int { return Config.GetInt(key) }
func (c *config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

func GetString(key string) string { return Config.GetString(key) }
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}

func GetStringMap(key string) map[string]interface{} { return Config.GetStringMap(key) }
func (c *config) GetStringMap(key string) map[string]interface{} {
	return c.viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string { return Config.GetStringMapString(key) }
func (c *config) GetStringMapString(key string) map[string]string {
	return c.viper.GetStringMapString(key)
}

func GetStringSlice(key string) []string { return Config.GetStringSlice(key) }
func (c *config) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func GetTime(key string) time.Time { return Config.GetTime(key) }
func (c *config) GetTime(key string) time.Time {
	return c.viper.GetTime(key)
}

func GetDuration(key string) time.Duration { return Config.GetDuration(key) }
func (c *config) GetDuration(key string) time.Duration {
	return c.viper.GetDuration(key)
}

func IsSet(key string) bool { return Config.IsSet(key) }
func (c *config) IsSet(key string) bool {
	return c.viper.IsSet(key)
}
