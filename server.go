package main

import (
	"os"

	"github.com/idio1981/go-pn-public/cmd"
	"github.com/idio1981/go-pn-public/logger"
	"github.com/spf13/cobra"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int `json:"port"`
}

// Server 服务器结构体
type Server struct {
	config *ServerConfig
	logger *logger.Logger
}

// NewServer 创建新的服务器实例
func NewServer(config *ServerConfig) *Server {
	return &Server{
		config: config,
		logger: logger.New(),
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	logger.Info("Starting server on port %d", s.config.Port)
	
	// 这里可以添加实际的服务器启动逻辑
	// 例如：HTTP服务器、gRPC服务器等
	
	return nil
}

// Stop 停止服务器
func (s *Server) Stop() error {
	logger.Info("Stopping server")
	
	// 这里可以添加实际的服务器停止逻辑
	
	return nil
}

// GetLogger 获取日志记录器
func (s *Server) GetLogger() *logger.Logger {
	return s.logger
}

// SetLogLevel 设置日志级别
func (s *Server) SetLogLevel(level int) {
	logger.SetLevel(level)
}

func main() {
	// 设置命令行参数
	cmd.Execute(func(cobraCmd *cobra.Command, args []string) {
		config := &ServerConfig{
			Port: cmd.ARGV_LISTEN_PORT,
		}
		
		server := NewServer(config)
		
		if err := server.Start(); err != nil {
			logger.Error("Failed to start server: %v", err)
			os.Exit(1)
		}
		
		// 保持服务器运行
		select {}
	})
}