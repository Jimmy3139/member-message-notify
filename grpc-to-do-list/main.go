package main

import (
	"grpc-to-do-list/config"
	//"grpc-to-do-list/database"
	route "grpc-to-do-list/routers"
	"strings"

	"github.com/gin-gonic/gin"
	gelfFormatter "github.com/seatgeek/logrus-gelf-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger *logrus.Logger = logrus.New()
var configuration *config.Configurations

// configInit 初始化設定檔
func configInit() {
	//初始化設定檔
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Fatalf("Unable to decode into struct, %v", err)
	}
}

// init 初始化
func init() {

	configInit()

	logInit()

}

// logInit 初始化log
func logInit() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.Formatter = new(gelfFormatter.GelfFormatter)

	// hook := graylog.NewAsyncGraylogHook(configuration.Gelf.Host, map[string]interface{}{"ServiceName": configuration.Gelf.LogSource})

	// logger.AddHook(hook)
	logger.Printf("zone_ballard service start")

	logger.SetLevel(configuration.Gelf.Level)
}

func main() {

	//db := database.New(configuration.MysqlDB)
	//repo := repository.NewRoleRepository(db, config)

	// Initialize Gin router
	gin.SetMode(configuration.Gin.Mode)
	router := gin.Default()
	groupApi := router.Group("/api/v1/")

	route.NewApiRoute(groupApi, logger, configuration).RouteApi()

	// Start the server
	router.Run(":80")
}
