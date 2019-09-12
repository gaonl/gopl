package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func uploadFile(fileName, url string) error {
	//http.DefaultTransport()
	buf := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(buf)

	idx := strings.LastIndex(fileName, "\\")
	fName := string([]byte(fileName)[idx:])
	fileFieldWriter, err := multipartWriter.CreateFormFile("uploadfile", fName)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(fileFieldWriter, f)
	f.Close()

	textFieldWriter, err := multipartWriter.CreateFormField("appendName")
	if err != nil {
		log.Fatalln(err)
	}
	textFieldWriter.Write([]byte("version1.1.1"))

	//一定要先在这边关掉
	multipartWriter.Close()

	http.Post(url, multipartWriter.FormDataContentType(), buf)

	return nil
}

func main() {
	uploadFile("C:\\Users\\Administrator\\Desktop\\p2p\\p2p需求20180119.txt", "http://127.0.0.1:8181/upload")
}
