package config

import "github.com/spf13/viper"

type KafkaConnCfg struct {
	Url   string `mapstructure:"kafka_url"`
	Topic string `mapstructure:"topic"`
}

func LoadConfig(path string) (config KafkaConnCfg, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
