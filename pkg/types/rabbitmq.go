package types

// rabbitMQ配置结构体
type RabbitMQConfig struct {
	Host     string // ip
	Port     string // 端口
	User     string // 用户名
	Password string // 密码
	VHost    string // 区域
}
