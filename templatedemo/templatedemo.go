package main

import (
	"html/template"
	"log"
	"os"
)

var TEMPLATE_STRING string = `

结构体：
user_name:{{.user.Name}}
user_age:{{.user.Age}}

map:
map_name:{{.map.UserName}}
map_age:{{.map.UserAge}}

slice:
{{range $index, $element := .slice}}
{{$index}}: {{$element.Name}}-{{$element.Age}}{{end}}

string:
{{.Title}}

条件分支：
若 pipeline 的值不为 empty，条件匹配，执行相应分支。
empty 空值为：false、0、任意nil指针或者nil接口，任意长度为0的数组、切片、字典

{{if .slice}} slice is not empty {{else}} slice is empty {{end}}
{{if .map}} map is not empty {{else}} map is empty {{end}}

`

type User struct {
	Name string
	Age int
}

func main(){
	t := template.New("test_template") //创建一个模板
	t,err := t.Parse(TEMPLATE_STRING)
	if err !=nil {
		log.Fatalln(err)
	}

	data := make(map[string]interface{})

	//结构体
	data["user"] = User{"ws",100}
	data["map"] = map[string]string{"UserName":"wskj","UserAge":"20"}
	data["slice"] = []User{{"ws1",101},{"ws2",102},{"ws3",103}}
	data["Title"] = "This is TITLE"

	t.Execute(os.Stdout, data) //执行模板的merger操作
}