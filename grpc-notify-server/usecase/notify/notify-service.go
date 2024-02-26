package notify

import (
	"context"
	"grpc-notify-server/config"
	mq "grpc-notify-server/pkg/mqttcli"
	"grpc-notify-server/proto"
)

type NotifyService struct {
	config *config.Configurations
	mq     *mq.MqttCli
}

func NewNotifyService(config *config.Configurations, mq *mq.MqttCli) *NotifyService {
	return &NotifyService{config: config, mq: mq}
}

func (s *NotifyService) SendNotification(context.Context, *proto.NotificationRequest) (*proto.NotificationResponse, error) {

	s.mq.PublishMemberRegisterInfo("notification message")

	return &proto.NotificationResponse{Success: true}, nil
}
