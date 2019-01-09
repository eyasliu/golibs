package log

import "testing"

func TestNewLogger(t *testing.T) {
	option := Options{
		Level:     "debug",
		Debug:     true,
		Formatter: "json",
		Write:     true,           // 是否输出到文件
		Path:      "./",           // 如果要输出到文件，指定目录路径
		FileName:  "test_log.log", // 日志文件名
	}
	Default(&option)

	Info("test log msg")
}
