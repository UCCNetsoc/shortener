package config

import (
	"strings"

	"github.com/spf13/viper"
)

// InitConfig initialises environment variables
func InitConfig() {
	initDefaults()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initDefaults() {
	viper.SetDefault("api.token", "")
	viper.SetDefault("db.host", "")
	viper.SetDefault("db.user", "")
	viper.SetDefault("db.port", "")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.database", "")
}
