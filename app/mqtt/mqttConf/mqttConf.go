package mqttConf

import "mqttGo/pkg/viper"

var (
	MqttHost = viper.Viper.GetString("mqtt.host")
	MqttPort = viper.Viper.GetString("mqtt.port")
)
