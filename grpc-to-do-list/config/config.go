package config

import "github.com/sirupsen/logrus"

// Configurations exported
type Configurations struct {
	Gin       GinConfigurations
	Redis     RedisConfigurations
	MachineID MachineIDConfigurations
	Gelf      GelfConfigurations
	MysqlDB   DBCollentSetting
}

// GinConfigurations exported
type GinConfigurations struct {
	Mode string
}

type DBCollentSetting struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// RedisConfigurations exported
type RedisConfigurations struct {
	ConnectionString string
	Password         string
}

type MachineIDConfigurations struct {
	Id int64
}

// GelfConfigurations exported
type GelfConfigurations struct {
	Host      string
	LogSource string
	Level     logrus.Level
}
