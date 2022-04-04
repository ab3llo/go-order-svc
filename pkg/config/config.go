package config

import "github.com/spf13/viper"

type Config struct {
	Port             string `mapstructure:"PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USER"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	JwtSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
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
