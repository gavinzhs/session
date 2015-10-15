package main

import (
	"fmt"
	//	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	//    "time"
	"io/ioutil"
	"regexp"
)

type test struct {
	desc string
	ch   chan int
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

type A struct {
	One int
	Two int
}

type WXAccessToken struct {
	access_token string
	expires_in   int
}

type testChan struct {
	C <-chan int
}

var REG_ORG_MEMBER = regexp.MustCompile(`^(\d+)\.(\d+)$`)

func main1() {

	//    add := func(a ...interface{}) int {
	//        if len(a) == 0{
	//            return 0
	//        }
	//        return a[0].(A).One + a[0].(A).Two
	//    }
	//
	//    a := &A{1, 2}
	//    b := &A{1, 2}
	//    slice := []interface{}{}
	//    slice = append(slice, a)
	//    slice = append(slice, b)
	//
	//    log.Println(add(slice...))

	//	m := make(map[string]A)
	//	m["a"] = A{One: 1, Two: 2}
	//	log.Println(m["a"])
	//	a := m["a"]
	//	a.One = 2
	//	log.Println(m["a"])
	//	if v, ok := m["a"]; ok {
	//		v.One = 3
	//	}
	//	log.Println(m["a"])

	//    m := REG_ORG_MEMBER.FindAllStringSubmatch("104593.100087", -1)
	//    log.Println(m)

	//    src := `郑赛
	//    你好`
	//    log.Println(url.QueryUnescape(src))
	//    log.Println(url.QueryUnescape(url.QueryEscape(src)))
	//
	//    testHttp()
	//    time.Sleep(time.Second * 10)

	//    m := make(map[int]int)
	//    m[1] = 1
	//    m[2] = 2
	//    m[3] = 3
	//    m[4] = 4
	//    m[5] = 5
	//    log.Println(m)

	//    c := make(chan int)
	//    t := &testChan{
	//        C : c,
	//    }
	//    go func(){
	//        t.C <- 5
	//    }()
	//
	//    log.Println(<- t.C)

	//    v, ok := m[1]
	//    if ok {
	//        log.Println("v:", v)
	//    }else{
	//        log.Println("no v")
	//    }

	//    if v, ok := m[2]; ok{
	//        log.Println("v:", v)
	//    }else{
	//        log.Println("no v")
	//    }

	//	log.Println("start")
	//
	//    var a A
	//    a.One = int32(1)
	//    a.Two = int32(2)
	//
	//    buf := new(bytes.Buffer)
	//    log.Printf("序列化A需要%d位的buf!", binary.Size(a))
	//    binary.Write(buf, binary.LittleEndian, a)
	//    log.Print("after write, buf is %q", buf.Bytes())
	//
	//    var aa A
	//    log.Println("aa before : ", aa)
	//    binary.Read(buf, binary.LittleEndian, &aa)
	//    log.Println("aa after : ", aa)
	//
	//    i := int32(1000)
	//    log.Println("size:", binary.Size(i))
	//    bufInt := new(bytes.Buffer)
	//    binary.Write(bufInt, binary.LittleEndian, i)
	//    log.Println("写完后是多少:", bufInt.Bytes())
	//
	//    var ii int32
	//    binary.Read(bufInt, binary.BigEndian, &ii)
	//    log.Println("反序列化是多少:", ii)

	//    jsonStr := "{'access_token':'b17gaRI0xvbyRoOoa2vTxRkL0xUwL1Gvn5mW5S7mXP13XuiVWGOHdJH8SNo5N_UdvFrCgNl0nUYElbWH05QMvryyGgrXGtyJiMg4cWbL8p0','expires_in':7200}"
	//    log.Println("jsonstr:", jsonStr)
	//    log.Println("json:", string([]byte(jsonStr)))
	//    var token *WXAccessToken
	//    err := json.Unmarshal([]byte(jsonStr), &token)
	//    if err == nil{
	//        log.Printf("token err:%v;;;", err)
	//    }
	//    log.Printf("token:%v;;;", token)

	//    var jsonBlob = []byte(`{"access_token":"b17gaRI0xvbyRoOoa2vTxRkL0xUwL1Gvn5mW5S7mXP13XuiVWGOHdJH8SNo5N_UdvFrCgNl0nUYElbWH05QMvryyGgrXGtyJiMg4cWbL8p0","expires_in":302}`)
	//    type Animal struct {
	//        Access_token  string
	//        Expires_in int
	//    }
	//    var animals Animal
	//    err := json.Unmarshal(jsonBlob, &animals)
	//    if err != nil {
	//        fmt.Println("error:", err)
	//    }
	//    fmt.Printf("%+v", animals)
	//
	//    for {
	//        time.Sleep(time.Second * time.Duration(animals.Expires_in-300))
	//        log.Println("once")
	//    }

	//        tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6666") //获取一个TCP地址信息,TCPAddr
	//        checkError(err)
	//        conn, err := net.Dial("tcp", "127.0.0.1:6666") //创建一个TCP连接:TCPConn
	//        checkError(err)
	////        _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) //发送HTTP请求头
	////        checkError(err)
	////        result, err := ioutil.ReadAll(conn) //获得返回数据
	////        checkError(err)
	//
	//
	//
	//    for {
	//        log.Println("进来了")
	////        b := make([]byte, 4)
	////        _, err = io.ReadFull(conn, b)
	////        checkError(err)
	////
	////        buf := bytes.NewBuffer(b)
	////        var n uint32
	////        err = binary.Read(buf, binary.BigEndian, &n)
	////        checkError(err)
	////
	////        b = make([]byte, n)
	////        _, err = io.ReadFull(conn, b)
	//                result, err := ioutil.ReadAll(conn) //获得返回数据
	//        checkError(err)
	//        log.Println("hold")
	//
	//        log.Println("这里是取到的数据:", string(result))
	//    }
	//
	//
	//
	////    b := make([]byte, 1024)
	////    log.Println("到这里吧:", string(b))
	////            _, err = io.ReadFull(conn, b)
	////    log.Println("没有到这里吧:", string(b))
	////
	////        os.Exit(1)
	//
	//    time.Sleep(time.Minute * 2)

	//    a := "a"
	//    print("字符串的长度:", len(a))
	//    print("字符串的长度byte:", len([]byte(a)))
	////    b := 1
	////    print("数字的长度:", "")
	////
	////    print("数字的长度byte:", len([]byte(b)))
	//    c := '\n'
	//    print("汉子的长度:", len(c))
	//    print("hanzi的长度byte:", len([]byte(c)))

	//    conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	//    for {     // read in input from stdin
	////        log.Println("到了这里面")
	//        reader := bufio.NewReader(os.Stdin)
	//        fmt.Print("Text to send: ")
	//        text, _ := reader.ReadString('\n')
	////          // send to socket
	////        fmt.Fprintf(conn, text + "\n")
	//        conn.Write([]byte(text))
	//            // listen for reply
	////        message, _ := bufio.NewReader(conn).ReadString('\n')
	//        b := make([]byte, 2048)
	//        message := make([]byte, 3)
	//        for {
	//            num, err := conn.Read(message)
	//
	//            if err == io.EOF{
	//                log.Printf("读完了 : %v", err)
	//                break
	//            }
	//
	//            if err != nil{
	//                log.Printf("读取错误 : %v", err)
	//            }
	//
	//            log.Print("数量:", num)
	//
	//            b = append(b, message...)
	//            fmt.Print("Message from server: "+string(b))
	//        }
	//        //        log.Println(string(msg))
	//
	//    }

	//    cmd1 := exec.Command("ps", "aux")
	//    cmd2 := exec.Command("grep", "go")
	//
	//    stdout1, err := cmd1.StdoutPipe()
	//    if err != nil{
	//        log.Printf("cmd1.StdoutPipe() : %+v", err)
	//    }
	//
	//    if err := cmd1.Start(); err != nil{
	//        log.Printf("cmd1.Start() : %+v", err)
	//    }
	//
	//    stdint2, err := cmd2.StdinPipe()
	//    if err != nil{
	//        log.Printf("cmd2.StdinPipe() : %+v", err)
	//    }
	//
	//    outputbuf := bufio.NewReader(stdout1)
	//    outputbuf.WriteTo(stdint2)
	//
	//    var output2 bytes.Buffer
	//    cmd2.Stdout = &output2
	//
	//    if err := cmd2.Start(); err != nil{
	//        log.Printf("cmd2.Start() : %+v", err)
	//    }
	//
	//    err = stdint2.Close()
	//    if err != nil{
	//        log.Printf("stdint2.Close() : %+v", err)
	//    }
	//
	//    if err := cmd2.Wait(); err != nil{
	//        log.Printf("cmd2.Wait() : %+v", err)
	//    }
	//
	//    log.Println(output2.String())

	//	martiniRoute()
	//    fmt.Printf("pid : %d; ppid : %d", os.Getpid(), os.Getppid())

	//    cmd := exec.Command("echo", "nihaoya!")
	//    log.Printf("cmd : %+v", cmd)
	//
	//    stdout0, err := cmd.StdoutPipe()
	//    if err != nil{
	//        log.Println("创建输出管道")
	//    }
	//
	//    if err := cmd.Start(); err != nil{
	//        log.Printf("cmd.start() err : %+v", err)
	//    }
	//
	////    var outputBufo bytes.Buffer
	////    output0 := make([]byte, 1)
	////    for{
	////        n, err := stdout0.Read(output0)
	////        if err != nil{
	////            if err == io.EOF{
	////                break
	////            }else{
	////                log.Printf("stdout0.Read(output0) err : %+v", err)
	////            }
	////        }
	////        outputBufo.Write(output0[:n])
	////    }
	////
	////    fmt.Printf("读取出来了%d个byte,都是什么: %s", outputBufo.Len(), outputBufo.String())
	//
	//
	//    outputBufo := bufio.NewReader(stdout0)
	//    output0, _, err := outputBufo.ReadLine()
	//    if err != nil{
	//        log.Printf("stdout0.Read(output0) err : %+v", err)
	//    }
	//    fmt.Printf("读取出来了%d个byte,都是什么: %s", len(output0), string(output0))

}

func testHttp() {
	query := `郑赛`
	str := "http://openapi.baidu.com/public/2.0/bmt/translate"
	q := url.Values{}
	q.Add("client_id", "ke0FZEFG4LrgD1Awow1pCdTu")
	q.Add("from", "auto")
	q.Add("to", "auto")
	q.Add("q", query)
	//    resp, _ := http.PostForm("http://openapi.baidu.com/public/2.0/bmt/translate",
	//    url.Values{"key": {"Value"}, "id": {"123"}})
	//    log.Println(resp)

	////    log.Println(url.QueryEscape(query))
	////    url := fmt.Sprintf("http://openapi.baidu.com/public/2.0/bmt/translate?client_id=ke0FZEFG4LrgD1Awow1pCdTu&q=%s&from=auto&to=auto", url.QueryEscape(query))
	//
	//

	////	req, err := http.NewRequest("POST", url, nil)
	////	if err != nil {
	////		log.Printf("new request err : %v", err)
	////        return
	////	}
	////
	////	res, err := http.DefaultClient.Do(req)
	////	if err != nil {
	////		log.Printf("new response err : %v", err)
	////        return
	////	}
	////    res, err := http.PostForm(url, q)
	//
	res, err := http.PostForm(str, q)

	log.Println("res.status : ", res.Status)
	log.Println("res.statuscode : ", res.StatusCode)
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ioutil.readall err : %v", err)
		return
	}

	log.Println("contentLength:", res.ContentLength)
	log.Println("Content-Encoding:", res.Header.Get("Content-Encoding"))
	log.Println("Content-Encoding:", res.Header.Get("Date"))
	log.Println("Etag:", res.Header.Get("Etag"))
	log.Println("body : ", len(string(b)))
	log.Println("body : ", string(b))
}
