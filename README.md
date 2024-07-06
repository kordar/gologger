# 日志组件

- 安装

```go
go get github.com/kordar/gologger v0.0.8
```

该组件定义一系列日志操作相关接口，提供统一全局API，注入不同组件完成日志组件全局切换。
默认通过fmt包实现简单日志打印，其他组件可参考实际实现

- [`nazalog`](https://github.com/kordar/nazalog)日志组件
- [`logrus`](https://github.com/sirupsen/logrus)日志组件





