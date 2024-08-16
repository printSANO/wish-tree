package main

import (
	"github.com/printSANO/wish-tree/config"
	"github.com/printSANO/wish-tree/server"
)

// 앱에서 가장 먼저 실행되는 함수
func main() {
	port := config.GetEnvVarAsString("PORT", "8080")
	server.Start(port)
}
