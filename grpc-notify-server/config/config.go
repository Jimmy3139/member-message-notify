package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Mqtt MqttConfigurations
}

type MqttConfigurations struct {
	Broker               string
	ClientID             string
	KeepAlive            int64
	TopicCorrection      string
	ActionMemberRegister string
}

func NewConfig(logger *logrus.Logger, configPath string) *Configurations {
	var config Configurations
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		logger.Fatalf("Unable to decode into struct, %v", err)
	}
	return &config
}
