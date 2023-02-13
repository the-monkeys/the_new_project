package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthSvcUrl        string `mapstructure:"AUTH_SVC_URL"`
	ArticleSvcUrl     string `mapstructure:"STORY_SVC_URL"`
	UserSvcUrl        string `mapstructure:"USER_SVC_URL"`
	BlogAndPostSvcURL string `mapstructure:"BLOGANDPOSTS_SVC_URL"`
}

func LoadGatewayConfig() (cfg Config, err error) {
	viper.AddConfigPath("/the_monkey/etc")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		logrus.Errorf("cannot load the config file, error: %+v", err)
		return
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		logrus.Errorf("cannot unmarshal the config file, error: %+v", err)
		return
	}

	return
}
