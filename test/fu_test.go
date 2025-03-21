package test

import (
	"com/chat/service/pkg/gin/middleware"
	"github.com/jinzhu/copier"
	"testing"
)

func Test(t *testing.T) {
	middleware.AddWhiteRouter(map[string]string{
		"/api/v1/user": "get",
	})
}

func TestCopier(t *testing.T) {
	type Ts struct {
		Name string `json:"name"`
	}
	type T1 struct {
		Name int `json:"name"`
	}
	ty := Ts{}
	t1 := T1{Name: 98}
	err := copier.Copy(&ty, &t1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ty.Name)
}
