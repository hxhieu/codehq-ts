package configuration

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	TimesheetDsn string `mapstructure:"CODEHQ_TS_TIMESHEET_DSN"`
	PimpDsn      string `mapstructure:"CODEHQ_TS_PIMP_DSN"`
	Debug        bool   `mapstructure:"CODEHQ_TS_DEBUG"`
}

var config *Config

func init() {
	// Default keys
	viper.SetDefault("CODEHQ_TS_TIMESHEET_DSN", "")
	viper.SetDefault("CODEHQ_TS_PIMP_DSN", "")
	viper.SetDefault("CODEHQ_TS_DEBUG", false)

	// The read the config file, if any
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		// Parse env as the fallback
		log.Println("No .env file found, using environment variables.")
		viper.AutomaticEnv()
	}

	// Deserialise
	config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse the configuration: %s", err))
	}
}

func Get() *Config {
	return config
}
