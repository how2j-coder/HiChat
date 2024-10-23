package test

import (
	"HiChat/utils"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Level int

func (l Level) String() string {
	return "this Level String"
}

func TestLeve(t *testing.T) {
	l := Level(5)
	v := reflect.ValueOf(l)
	fmt.Println(v.Int())
}

func TestSnowflake(t *testing.T) {
	machineID := 1 // 根据实际情况设置机器ID
	sf := utils.NewSnowflake(int64(machineID))
	// 生成10个唯一ID并输出
	for i := 0; i < 1000; i++ {
		go func() {
			id := sf.GenerateID()
			fmt.Println(id)
		}()
	}
	time.Sleep(10 * time.Second)
}
