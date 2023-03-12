package config

import "github.com/spf13/viper"

type Config struct {
	UserSrvPort       string `mapstructure:"USER_SERVICE"`
	DBUrl             string `mapstructure:"DB_URL"`
	BlogAndPostSvcURL string `mapstructure:"BLOG_SERVICE"`
}

func LoadUserConfig() (config Config, err error) {
	viper.AddConfigPath("/the_monkeys/etc")
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
