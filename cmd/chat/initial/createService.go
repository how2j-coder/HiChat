package initial

import (
	"com/chat/service/internal/config"
	"com/chat/service/internal/server"
	"com/chat/service/pkg/app"
	"strconv"
)

// CreateService 创建一个服务
func CreateService() []app.GoServer {
	var cfg = config.GetConfig()
	var servers []app.GoServer

	httpAddr := ":" + strconv.Itoa(cfg.HTTP.Port)
	httpServer := server.NewHTTPServer(httpAddr,
		server.WithHTTPIsProd(cfg.App.Env == "prod"),
	)
	servers = append(servers, httpServer)
	return servers
}
