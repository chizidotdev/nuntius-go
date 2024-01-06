package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PORT               string `mapstructure:"PORT"`
	PostgresUrl        string `mapstructure:"POSTGRES_URL"`
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	AuthCallbackUrl    string `mapstructure:"AUTH_CALLBACK_URL"`
	BaseUrl            string `mapstructure:"BASE_URL"`
}

var EnvVars Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read config file:", err)
	}
	viper.AutomaticEnv()

	EnvVars.PORT = viper.GetString("PORT")
	EnvVars.PostgresUrl = viper.GetString("POSTGRES_URL")
	EnvVars.GoogleClientID = viper.GetString("GOOGLE_CLIENT_ID")
	EnvVars.GoogleClientSecret = viper.GetString("GOOGLE_CLIENT_SECRET")
	EnvVars.AuthCallbackUrl = viper.GetString("AUTH_CALLBACK_URL")
	EnvVars.BaseUrl = viper.GetString("BASE_URL")
}
