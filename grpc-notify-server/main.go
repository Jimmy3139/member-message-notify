package main

import (
	"grpc-notify-server/config"
	"grpc-notify-server/pkg/mqttcli"
	"grpc-notify-server/proto"
	"grpc-notify-server/usecase/notify"
	"log"
	"net"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	configuration *config.Configurations
	mqttCli       *mqttcli.MqttCli
)

// init 初始化
func init() {
	//configInit 初始化設定檔
	configInit()
	mqttCli = mqttcli.InitMqttCli(configuration)
}
func configInit() {
	//初始化設定檔
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetConfigType("yml")

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	// 创建NotifyService实例时传入config和mq
	notifyService := notify.NewNotifyService(configuration, mqttCli)
	proto.RegisterNotificationServiceServer(server, notifyService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
