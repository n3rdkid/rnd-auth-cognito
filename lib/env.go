package lib

import (
	"github.com/spf13/viper"
)

type Env struct {
	CognitoAppClientId string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	CognitoRegion      string `mapstructure:"COGNITO_REGION"`
}

// Will be populated by NewEnv
var env = Env{}

func GetEnv() Env {
	return env
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")
	viper.SetDefault("TIMEZONE", "UTC")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return &env
}
