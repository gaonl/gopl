package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	TMP_BASE_PATH string = "F:\\my\\learn_go\\src\\server\\"
)

var handlerFunc = func(w http.ResponseWriter, r *http.Request) {
	/*
	r.ParseForm()                       //解析参数，默认是不会解析的
	fmt.Println("query form: ", r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("requestUri: ", r.RequestURI)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	 */

	t, err := template.ParseFiles(TMP_BASE_PATH + "index.gtpl")
	if err != nil {
		log.Println(err)
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	t.Execute(w, nil)
}

func loginFunC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles(TMP_BASE_PATH + "login.gtpl")
		if err != nil {
			log.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		data := map[string]string{"username": "wskj", "password": "123456"}
		t.Execute(w, data)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func uploadFunc (w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {

	}else{

	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/login", loginFunC)
	http.HandleFunc("/upload", uploadFunc)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
