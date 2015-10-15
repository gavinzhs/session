package main

import (
	"io"
	"net/http"
	"testing"
)

func TestHttps(t *testing.T) {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world!")
	})

	//        err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)//todo 需要生成这两个证书
	//        if err != nil {
	//            log.Fatal("ListenAndServeTLS:", err.Error())
	//        }
}
