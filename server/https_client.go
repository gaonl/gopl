package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyRoundTrip struct {
}

func (myTrip *MyRoundTrip) RoundTrip(*http.Request) (*http.Response, error) {
	f, err := os.Open("C:\\Users\\Administrator\\Desktop\\p2p\\p2p_douyu.html")
	if err != nil {
		return nil, err
	}

	header := make(http.Header)
	header["Content-Encoding"] = []string{"gzip"}

	resp := &http.Response{
		Header: header,
		Body:   f,
	}
	return resp, nil
}

func main() {
	//response, err := http.Get("https://www.taobao.com/")
	v, _ := http.DefaultTransport.(*http.Transport)
	my := *v
	my.DisableCompression = true
	client := &http.Client{
		//Transport: &MyRoundTrip{},

		Transport: &my,
	}
	response, err := client.Get("https://www.taobao.com/")
	if err != nil {
		log.Fatalln(err)
	}

	body := response.Body
	defer body.Close()

	fmt.Println(response.Header["Content-Encoding"])
	io.Copy(os.Stdout, body)
}
