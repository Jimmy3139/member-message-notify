package conf

// import (
// 	"os"

// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	Database  Database
// 	MachineID MachineIDSetting
// }
// type Database struct {
// 	MySQL DBSetting
// }

// type DBSetting struct {
// 	Master DBCollentSetting
// 	Slave  DBCollentSetting
// }
// type DBCollentSetting struct {
// 	Host     string `yaml:"host"`
// 	Port     string `yaml:"port"`
// 	User     string `yaml:"user"`
// 	Password string `yaml:"password"`
// 	Name     string `yaml:"name"`
// }

// type MachineIDSetting struct {
// 	Id int64 `yaml:"id"`
// }

// func Load() (conf *Config, err error) {
// 	viper.SetConfigFile(GetPath())
// 	err = viper.ReadInConfig()

// 	if err != nil {
// 		return nil, err
// 	}
// 	err = viper.Unmarshal(&conf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return conf, nil
// }
// func GetPath() string {
// 	workdir, _ := os.Getwd()
// 	switch os.Getenv("MODE") {
// 	case "development":
// 		return workdir + "/configs/config.dev.yml"
// 	default:
// 		return workdir + "/configs/config.yml"
// 	}
// }
