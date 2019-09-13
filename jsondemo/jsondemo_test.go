package jsondemo

import (
	"encoding/json"
	"fmt"
	"github.com/gaonl/gopl/domain"
	"log"
	"strings"
	"testing"
)

const jsonMap = `{
	"Id": 1,
	"Name": "wskj",
	"NickNames": ["apple","orange"],
	"Password": "123465",
	"RegisterDateTime": 1568276700
}`

const jsonArray = `[
	{
		"Id": 1,
		"Name": "wskj",
		"NickNames": ["apple","orange"]
		"Password": "123465",
		"RegisterDateTime": 1568276700
	},{
		"Id": 2,
		"Name": "qkl",
		"NickNames": ["bitcoin","block"]
		"Password": "private key",
		"RegisterDateTime": 1568276700
	}
]`

func TestStruct(t *testing.T) {
	//结构体类型
	var user *domain.User //指针
	//err := json.NewDecoder(strings.NewReader(jsonMap)).Decode(&user) //取地址，因为要修改user指针的值
	err := json.Unmarshal([]byte(jsonMap), &user) //取地址，因为要修改user指针的值
	if err != nil {
		t.Error(err)
	}
	userActual := domain.User{
		Id:               1,
		Name:             "wskj",
		NickNames:        [2]string{"apple", "orange"},
		Password:         "123466",
		RegisterDateTime: 1568276700,
	}
	if *user != userActual {
		t.Errorf("----->TestStruct fail")
	}
}
func TestMap(t *testing.T) {
	//map类型
	var m map[string]interface{} //map类型其实底层是指针
	//err := json.NewDecoder(strings.NewReader(jsonMap)).Decode(&m) //取地址，因为要修改map指针的值
	err := json.Unmarshal([]byte(jsonMap), &m) //取地址，因为要修改map指针的值
	if err != nil {
		t.Error(err)
	}

	if m["Id"] != 1 {
		t.Errorf("----->Id %d match fail", m["Id"])
	}
	if m["Name"] != "wskj" {
		t.Error("----->Name match fail")
	}
	if m["NickNames"] != [2]string{"apple", "orange"} {
		t.Error("----->NickNames match fail")
	}
	if m["Password"] != "123466" {
		t.Error("----->Password match fail")
	}
	if m["RegisterDateTime"] != 1568276700 {
		t.Error("----->RegisterDateTime match fail")
	}
}
func TestArray(t *testing.T) {
	//数组类型
	var s []string //对象头
	//err := json.Unmarshal([]byte(jsonArray), &s) //取对象头的地址，让Unmarshal修改对象头
	err := json.NewDecoder(strings.NewReader(jsonArray)).Decode(&s) //取对象头的地址，让Unmarshal修改对象头
	if err != nil {
		log.Fatalln(err)
	}
	for i, e := range s {
		fmt.Println(i, e)
	}
}
