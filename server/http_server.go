package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type BasePath struct {
	basePath    string
	baseTmpPath string
}

func (bp BasePath) Path(filename string) string {
	return bp.basePath + filename
}

func (bp BasePath) TmpPath(filename string) string {
	return bp.baseTmpPath + filename
}

var TMPLATE_PATH BasePath = BasePath{
	basePath:    "F:\\my\\learn_go\\src\\server\\",
	baseTmpPath: "F:\\my\\learn_go\\tmp\\",
}

var handlerFunc = func(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(TMPLATE_PATH.Path("index.gtpl"))
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
		t, err := template.ParseFiles(TMPLATE_PATH.Path("login.gtpl"))
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

func uploadFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(TMPLATE_PATH.Path("upload.gtpl"))
		if err != nil {
			log.Println(err)
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(10240)
		appendName := r.FormValue("appendName")
		fUpload, h, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
			return
		}
		defer fUpload.Close()
		fileLocal, err := os.OpenFile(TMPLATE_PATH.TmpPath(h.Filename+"_"+appendName), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		defer fileLocal.Close()
		io.Copy(fileLocal, fUpload)
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
