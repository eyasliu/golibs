# 日志库

```go
import "github.com/eyasliu/golibs/log"

func main() {
  // logger 是 *logrus.Logger 类型
  logger, err := log.Default(&log.Options{
		Level:        "info",
		Formatter:    "json",
		Write:        true,
		Path:         "logs",
		FileName:     "songstore",
		MaxAge:       7,
		RotationTime: time.Duration(1*24) * time.Hour,
		Debug:        true,
  })

  if err != nil {
    panic(err)
  }
  
  log.Info("your msg")
}
```

配置项

```go
// Options 日志配置
type Options struct {
	Level         string        // 级别
	Formatter     string        // 输出格式 json | text
	Write         bool          // 是否输出到文件
	Path          string        // 如果要输出到文件，指定目录路径
	FileName      string        // 日志文件名
	MaxAge        int           // 只保存多少天的日志
	RotationCount uint          // 要保存多少个日志文件 只在maxAge 为 -1 时用
	RotationTime  time.Duration // 每隔多久记录一个文件
	Debug         bool          // 是否输出打日志的文件名和行号
}
```