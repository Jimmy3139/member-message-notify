package mqttcli

import (
	"encoding/json"
	"grpc-notify-server/config"
	dtos "grpc-notify-server/dtos/mqtt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type MqttCli struct {
	cli    mqtt.Client
	config *config.Configurations
	logger *logrus.Logger
}

// InitMqttCli 初始化MQTT
func InitMqttCli(conf *config.Configurations) *MqttCli {
	// mqtt.DEBUG = l
	opts := mqtt.NewClientOptions().AddBroker(conf.Mqtt.Broker).SetClientID(conf.Mqtt.ClientID)
	opts.SetKeepAlive(time.Duration(conf.Mqtt.KeepAlive) * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &MqttCli{cli: c}
}

// 推送會員註冊通知
func (m *MqttCli) PublishMemberRegisterInfo(message interface{}) (mqtt.Token, error) {
	msg := dtos.MQTTResultDto{
		ActionName: m.config.Mqtt.ActionMemberRegister,
		Message:    message,
	}
	msgString, err := m.getJSONString(msg)
	if err != nil {
		m.logger.Errorln("PublishMaintainInfo convert msg to json error:", err)
		return nil, err
	}

	token := m.cli.Publish(m.config.Mqtt.TopicCorrection, 1, false, msgString)

	if token.Error() != nil {
		m.logger.Errorln("PublishMaintainInfo Push error:", token.Error())
		return nil, token.Error()
	}

	return token, nil
}

// getJSONString 取得JSONString
func (m *MqttCli) getJSONString(obj interface{}) (string, error) {
	objByte, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	m.logger.Trace("PublishFootBallMatchFoul Push:", string(objByte))

	return string(objByte), nil
}
