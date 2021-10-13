package configuration

import (
	"fmt"
	"log"

	"github.com/hxhieu/codehq-ts/output"
	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool     `mapstructure:"CODEHQ_TS_DEBUG"`
	TimesheetDsn string   `mapstructure:"CODEHQ_TS_TIMESHEET_DSN"`
	PimpDsn      string   `mapstructure:"CODEHQ_TS_PIMP_DSN"`
	DateFormat   string   `mapstructure:"CODEHQ_TS_DATE_FORMAT"`
	TimeFormat   string   `mapstructure:"CODEHQ_TS_TIME_FORMAT"`
	AllowedCodes []string `mapstructure:"CODEHQ_TS_ALLOWED_CODES"`
}

var config *Config

func init() {
	// Default keys
	viper.SetDefault("CODEHQ_TS_DEBUG", false)
	viper.SetDefault("CODEHQ_TS_TIMESHEET_DSN", "") // dd/MM/yyyy
	viper.SetDefault("CODEHQ_TS_PIMP_DSN", "")
	viper.SetDefault("CODEHQ_TS_DATE_FORMAT", "02/01/2006") // dd/MM/yyyy
	viper.SetDefault("CODEHQ_TS_TIME_FORMAT", "1504")       // HHmm
	viper.SetDefault("CODEHQ_TS_ALLOWED_CODES", []string{
		"PLAN",
		"ANLYS",
		"CODIG",
		"TESTG",
		"REVEW",
		"RETRO",
		"GROOM",
		"ADMIN",
		"AFTER",
		"CONSY",
		"DOCUM",
		"EDUCN",
		"INSTL",
		"MEETG",
		"NETWK",
		"SALES",
		"STAFF",
		"TRAIN",
		"TTIME",
	})

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
		err := fmt.Errorf("failed to deserialise the configuration: %s", err)
		output.ConsoleErrorJson(&err)
	}
}

func Get() *Config {
	return config
}
