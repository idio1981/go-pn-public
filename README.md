# go-pn-public

一个Go语言的服务器框架，提供命令行工具和可复用的API组件。

## 功能特性

- 🚀 简单的服务器启动和停止命令
- 📝 彩色日志输出系统
- 🔧 可配置的服务器参数
- 📦 可复用的API组件
- 🛠️ 基于Cobra的命令行界面

## 安装

```bash
go get github.com/idio1981/go-pn-public
```

## 使用方法

### 作为命令行工具

```bash
# 启动服务器
./your-app start -p 8080

# 后台模式启动
./your-app start -p 8080 -d

# 停止服务器
./your-app stop
```

### 作为Go模块引用

```go
package main

import (
    "fmt"
    "github.com/idio1981/go-pn-public/api"
    "github.com/idio1981/go-pn-public/logger"
)

func main() {
    // 创建服务器配置
    config := &api.ServerConfig{
        Port: 8080,
    }
    
    // 创建服务器实例
    server := api.NewServer(config)
    
    // 设置日志级别
    server.SetLogLevel(logger.LevelDebug)
    
    // 启动服务器
    if err := server.Start(); err != nil {
        logger.Error("Failed to start server: %v", err)
        return
    }
    
    // 获取日志记录器
    log := server.GetLogger()
    log.Info("Server is running on port %d", server.GetPort())
}
```

### 单独使用日志组件

```go
package main

import (
    "github.com/idio1981/go-pn-public/logger"
)

func main() {
    // 设置日志级别
    logger.SetLevel(logger.LevelDebug)
    
    // 使用不同级别的日志
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")
    logger.Success("This is a success message")
    logger.Debug("This is a debug message")
}
```

## API 文档

### Server API

- `NewServer(config *ServerConfig) *Server` - 创建新的服务器实例
- `Start() error` - 启动服务器
- `Stop() error` - 停止服务器
- `GetLogger() *logger.Logger` - 获取日志记录器
- `SetLogLevel(level int)` - 设置日志级别
- `GetPort() int` - 获取服务器端口
- `SetPort(port int)` - 设置服务器端口

### Logger API

- `New() *Logger` - 创建新的日志记录器
- `SetLevel(level int)` - 设置日志级别
- `Info(format string, v ...interface{})` - 输出信息日志
- `Warn(format string, v ...interface{})` - 输出警告日志
- `Error(format string, v ...interface{})` - 输出错误日志
- `Success(format string, v ...interface{})` - 输出成功日志
- `Debug(format string, v ...interface{})` - 输出调试日志
- `Panic(format string, v ...interface{})` - 输出致命错误并退出

### 日志级别

```go
const (
    LevelError = iota
    LevelWarning
    LevelSuccess
    LevelInfo
    LevelDebug
)
```

## 开发

### 构建

```bash
go build -o your-app
```

### 运行测试

```bash
go test ./...
```

### 依赖管理

```bash
go mod tidy
```

## 版本要求

- Go 1.21+

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

## 更新日志

### v1.0.0
- 初始版本发布
- 支持基本的服务器启动/停止功能
- 提供彩色日志输出
- 支持命令行参数配置