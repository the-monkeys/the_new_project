package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"ARTICLE_SERVICE_PORT"`
	OSAddress string `mapstructure:"OPENSEARCH_ADDRESS"`
	Username  string `mapstructure:"USERNAME"`
	Password  string `mapstructure:"PASSWORD"`
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

// OSClient *opensearch.Client
