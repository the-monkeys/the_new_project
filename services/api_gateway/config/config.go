package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Address struct {
	APIGatewayHTTPS string `mapstructure:"API_GATEWAY_HTTPS"`
	APIGatewayHTTP  string `mapstructure:"API_GATEWAY_HTTP"`
	AuthService     string `mapstructure:"AUTH_SERVICE"`
	StoryService    string `mapstructure:"STORY_SERVICE"`
	UserService     string `mapstructure:"USER_SERVICE"`
	BlogService     string `mapstructure:"BLOG_SERVICE"`
}

func LoadGatewayConfig() (cfg Address, err error) {
	viper.AddConfigPath("/the_monkeys/etc")
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
