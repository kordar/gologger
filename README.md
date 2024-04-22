# 日志组件

提供统一接口，底层桥接不同实现类，完成不同日志系统的接入使用。所有日志桥接类继承`logger`接口，实现统一日志入口，提供`exports`导出静态日志方法。



- 依赖包

```go
github.com/kordar/gologger 
```



## 使用`nazalog`日志

- 初始化

```go
import (
	logger "github.com/kordar/gologger"
	"github.com/kordar/nazalog"
)

func InitLogger() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.Level = nazalog.LevelInfo
		option.Filename = "./logs/progress.log" 
		option.IsRotateDaily = true
		option.LevelFlag = true
	})
	defer nazalog.Sync()
	logger.InitGlobal(logger.NewNazalogAdapt())
}
```

- 使用

```go
import (
	logger "github.com/kordar/gologger"
)

logger.Info("this is nazalog!")
```



## 使用`logrus`日志

```go
import (
	logger "github.com/kordar/gologger"
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrusLog := logrus.New()
	logrusLog.SetLevel(logrus.InfoLevel)
	src, err := os.OpenFile(os.DevNull, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	logrusLog.SetOutput(writer)
	Log = logger.NewLogrusAdapt(logrusLog)
}
```



