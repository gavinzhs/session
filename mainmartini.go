package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	//	"strings"
	//	"time"
)

func main2() {
	//	m := martini.Classic()
	//	m.Get("/m", testHandler)
	//	m.Get("/m/main", testMainHandler)
	////		m.Run()
	//	http.Handle("/m/", m)
	//
	//	voip := martini.Classic()
	//	voip.Get("/voip/auth", testMainM2Handler)
	//	http.Handle("/voip/", voip)
	//
	//	http.ListenAndServe("localhost:3000", nil)

	//    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777") //获取一个tcpAddr
	//    checkMainError(err)
	//    listener, err := net.ListenTCP("tcp", tcpAddr) //监听一个端口
	//    checkMainError(err)

	//    listener, err := net.Listen("tcp", ":6666")
	//    checkMainError(err)
	//    for {
	//        log.Print("start")
	//        conn, err := listener.Accept()
	//        log.Println("hold")
	//        if err != nil {
	//            continue
	//        }
	//        daytime := time.Now().String()
	////        read := []byte{}
	////        _, err = conn.Read(read)
	////        if err != nil{
	////            log.Println("read err :", err)
	////        }
	//
	////        b := make([]byte, 1024)
	////        _, err = io.ReadFull(conn, b)
	////        log.Println("receive:", string(b))
	//
	////        result, err := ioutil.ReadAll(conn) //获得返回数据
	////        fmt.Println(string(result))
	//
	//        go func(){
	//            time.Sleep(time.Second * 5)
	//            log.Print("发出了")
	//            conn.Write([]byte(daytime))
	//        }()
	//        //        conn.Write([]byte("你好呀"))
	////        conn.Write([]byte("I'm fine!"))
	//
	//
	////        conn.Write([]byte(daytime))
	//        conn.Close()

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port
	for {

		conn, _ := ln.Accept()
		// run loop forever (or until ctrl-c)

		go func(conn net.Conn) {
			for {
				// will listen for message to process ending in newline (\n)
				//            message, _ := bufio.NewReader(conn).ReadString('\n')
				// output message received

				b := make([]byte, 2048)
				message := make([]byte, 3)

				for {
					num, err := conn.Read(message)

					if err == io.EOF {
						log.Printf("读完了 : %v", err)
						return
					}

					if err != nil {
						log.Printf("读取错误 : %v", err)
					}

					log.Print("数量:", num)

					b = append(b, message...)
					fmt.Print("Message from server: " + string(b))
				}

				//				fmt.Print("Message Received:", string(message))
				//
				//				// sample process for string received
				//				newmessage := strings.ToUpper(string(message))
				//
				//				if newmessage == "" {
				//					log.Print("收到数据为空    算了")
				//					continue
				//				}
				//				// send new string back to client
				//				for {
				//					time.Sleep(time.Second * 10)
				//					conn.Write([]byte(newmessage))
				//				}
			}
		}(conn)

	}

}

func checkMainError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func testHandler() string {
	return "hello world"
}

func testM1Handler() string {
	return "hello world m1"
}

func testMainHandler() string {
	return "hello world man !!!"
}

func testMainM2Handler() string {
	return "hello world man !!! M2"
}
