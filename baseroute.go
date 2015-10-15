package main

/*
   这里是处理原生的go路由
*/

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数,默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到 w 的是输出到客户端的
}

func baseroute() {
	//	initDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	host := os.Getenv("HOST")
	log.Println("host:", host)
	if host == "" {
		host = "localhost"
		log.Println("是空得 我需要重新设置host值")
	}

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/first/", sayhelloName)
	http.HandleFunc("/second/", sayhelloName)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("我擦,有人访问我了")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "hello")
}

type Person struct {
	Name  string
	Phone string
}

func initDB() {
	session, err := mgo.Dial("test:test@182.92.166.138/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("user")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
