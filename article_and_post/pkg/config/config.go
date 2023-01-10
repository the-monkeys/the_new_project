package config

import "github.com/spf13/viper"

type Config struct {
	ArticleServerPort string `mapstructure:"ARTICLE_SERVICE_PORT"`
	OSAddress         string `mapstructure:"OPENSEARCH_ADDRESS"`
	OSUsername        string `mapstructure:"OSUSERNAME"`
	OSPassword        string `mapstructure:"OSPASSWORD"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("/etc/the_monkey")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
