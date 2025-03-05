package test

import (
	"com/chat/service/pkg/gin/middleware"
	"testing"
)

func Test(t *testing.T)  {
	middleware.AddWhiteRouter(map[string]string{
		"/api/v1/user": "get",
	})
}
