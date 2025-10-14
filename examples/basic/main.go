package main

import (
	"fmt"
	"time"

	"github.com/idio1981/go-pn-public/api"
	"github.com/idio1981/go-pn-public/logger"
)

func main() {
	// 设置日志级别
	logger.SetLevel(logger.LevelDebug)

	// 创建服务器配置
	config := &api.ServerConfig{
		Port: 8080,
	}

	// 创建服务器实例
	server := api.NewServer(config)

	// 启动服务器
	if err := server.Start(); err != nil {
		logger.Error("Failed to start server: %v", err)
		return
	}

	// 获取日志记录器
	log := server.GetLogger()
	log.Info("Server is running on port %d", server.GetPort())

	// 模拟运行一段时间
	time.Sleep(5 * time.Second)

	// 停止服务器
	if err := server.Stop(); err != nil {
		logger.Error("Failed to stop server: %v", err)
		return
	}

	fmt.Println("Server stopped successfully")
}
