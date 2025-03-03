package initial

import (
	"com/chat/service/internal/config"
	"com/chat/service/internal/database"
	"com/chat/service/pkg/app"
	"com/chat/service/pkg/logger"
)

// Shutdown Close releasing resources after service exit
func Shutdown(servers []app.GoServer) []app.Close {
	var closes []app.Close

	// close server
	for _, s := range servers {
		closes = append(closes, s.Stop)
	}

	// close database
	closes = append(closes, func() error {
		return database.CloseMysql()
	})

	// close redis
	if config.GetConfig().App.CacheType == "redis" {
		closes = append(closes, func() error {
			return database.CloseRedis()
		})
	}


	// close logger
	closes = append(closes, func() error {
		return logger.Sync()
	})

	return closes
}
