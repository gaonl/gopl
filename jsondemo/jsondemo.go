package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name string
	Age  int
}

const jsonMap = `{
	"Name": "wskj",
	"Age": 30
}`

const jsonArray = `["a","b","c"]`

func main() {
	//结构体类型
	/*var user *User                                                   //指针，或者值都可以（var user User）
	err := json.NewDecoder(strings.NewReader(jsonMap)).Decode(&user) //取地址，因为要修改user指针的值，只要这边取地址就行
	//err := json.Unmarshal([]byte(jsonMap), &user)                    //取地址，因为要修改user指针的值
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)*/

	//map类型
	var m map[string]interface{}                                  //map类型其实底层是指针
	err := json.NewDecoder(strings.NewReader(jsonMap)).Decode(&m) //取地址，因为要修改map指针的值
	//err := json.Unmarshal([]byte(jsonMap), &m)                       //取地址，因为要修改map指针的值
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(m["Name"])
	fmt.Println(m["Age"])

	//数组类型
	//var s []string //对象头
	//err := json.Unmarshal([]byte(jsonArray), &s) //取对象头的地址，让Unmarshal修改对象头
	//err := json.NewDecoder(strings.NewReader(jsonArray)).Decode(&s) //取对象头的地址，让Unmarshal修改对象头
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for i, e := range s {
	//	fmt.Println(i, e)
	//}
}
