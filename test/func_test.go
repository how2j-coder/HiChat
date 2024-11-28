package test

import (
	"HiChat/utils"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strconv"
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
	for i := 0; i < 2; i++ {
		go func() {
			id := sf.GenerateID()
			fmt.Println(strconv.FormatInt(sf.GenerateID(), 10))
			fmt.Println(id)
		}()
	}
	time.Sleep(10 * time.Second)
}

type CommonTest struct {
	Name string
}

func (t *CommonTest) log() {
	fmt.Println(t.Name)
}

type User struct {
	CommonTest
	Age string
}

func TestCommon(t *testing.T) {
	user := User{Age: "123123", CommonTest: CommonTest{Name: "how2j"}}
	user.log()
}

type Test struct {
	ID   int
	Name string
}

func TestJSONData(t *testing.T) {
	// 创建一些 User 实例
	users := []*Test{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	// 将 []*models.User 转换为 JSON 字符串
	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// 输出 JSON 字符串
	fmt.Println(string(jsonData))
}

func TestTypeDefault(t *testing.T) {
	test := func(status ...bool) {
		fmt.Println(status)
	}
	test()
}

// 盐值
func TestSlat(t *testing.T) {
	uuidValue := uuid.New()
	fmt.Println(uuidValue.String())

}





