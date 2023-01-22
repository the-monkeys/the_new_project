package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ArticleServerPort string `mapstructure:"ARTICLE_SERVICE_PORT"`
	OSAddress         string `mapstructure:"OPENSEARCH_ADDRESS"`
	OSUsername        string `mapstructure:"OSUSERNAME"`
	OSPassword        string `mapstructure:"OSPASSWORD"`
	ArticleSvcUrl     string `mapstructure:"STORY_SVC_URL"`
}

func LoadArtNPostConfig() (config Config, err error) {
	viper.AddConfigPath("/etc/the_monkey")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		logrus.Errorf("cannot load the config file, error: %+v", err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		logrus.Errorf("cannot unmarshal the config file, error: %+v", err)
		return
	}

	return
}
