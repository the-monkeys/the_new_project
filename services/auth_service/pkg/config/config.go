package config

import "github.com/spf13/viper"

type Config struct {
	AuthAddr     string `mapstructure:"AUTH_SERVICE"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`

	SMTPAddress  string `mapstructure:"SMTP_ADDRESS"`
	SMTPMail     string `mapstructure:"SMTP_MAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	SMTPHost     string `mapstructure:"SMTP_HOST"`
}

func LoadConfig() (config Config, err error) {
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
