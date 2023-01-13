package config

import "github.com/spf13/viper"

type Config struct {
	UserSrvPort string `mapstructure:"USER_SERVICE_PORT"`
	DBUrl       string `mapstructure:"DB_URL"`
}

func LoadUserConfig() (config Config, err error) {
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
