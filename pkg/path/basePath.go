package path

import (
	"os"
	"path/filepath"
)

// 获取根目录
var BasePath = myBasePath()

// export MQTT_GO=/Users/a1/work/mqtt-go
func myBasePath() string {
	var path string
	var ok bool
	if path, ok = os.LookupEnv("MQTT_GO"); ok {
		return path
	}

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("根目录地址查询失败：" + err.Error())
	}
	return path
}
