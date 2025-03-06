package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"mqttGo/app/mqtt/mqttConf"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	var connStr = fmt.Sprintf("tcp://%s:%s", mqttConf.MqttHost, mqttConf.MqttPort)
	opts := mqtt.NewClientOptions().AddBroker(connStr).SetClientID("emqx_test_client")

	opts.SetUsername("tongwz")
	opts.SetPassword("123456")
	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(10 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 发布消息
	i := 1
	for {
		token := c.Publish("testtopic/tong", 0, false, "童伟珍测试"+fmt.Sprint(i))
		token.Wait()
		time.Sleep(5 * time.Second)
		i++
	}

	time.Sleep(6 * time.Second)

	// 取消订阅
	if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	select {}
	// 断开连接
	// c.Disconnect(250)
	// time.Sleep(1 * time.Second)
}
