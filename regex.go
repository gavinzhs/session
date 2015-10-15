package main

/*
   这里是验证正则表达式的地方
*/

import (
	"log"
	"regexp"
)

const (
	CONTENT = `<p><img src="http://mrocker-test.qiniudn.com/55603e01f92cff029e000001" style="" title="屏幕快照 2015-05-21 上午11.07.09.png"/></p><p><img src="http://mrocker-test.qiniudn.com/55603e01f92cff029e000002" style="" title="屏幕快照 2015-05-23 下午3.38.56.png"/></p><p><img src="http://mrocker-test.qiniudn.com/5560321ef92cff04da000002" title="555ddc0de9f0a30c08000003.png" alt="555ddc0de9f0a30c08000003.png" style="width: 100%;"/><br/></p>`
)

//type test struct {
//	Id bson.ObjectId
//}

func regex() {

	//    a := "123"
	//    b := a
	//    b = "456"
	//    log.Print(a, b)
	//
	//    log.Print(strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))

	//    bytes := []byte{1,2}
	//    bytes = strconv.AppendBool(bytes, true)
	//    log.Print(bytes)
	//    log.Print([]byte("t"))
	//    log.Print(string(rune(116)))
	//    log.Print(string(116))
	//    log.Print(strconv.Itoa(116))
	//    log.Print(strconv.FormatBool(true))
	//    log.Print(strconv.ParseBool("2"))
	//    print("hello world;")
	//    reg := regexp.MustCompile(`<img src="http://mrocker-test.qiniudn.com/.{24}`)
	reg := regexp.MustCompile(`<img src="http://mrocker-test.qiniudn.com/.{24}`)
	imgs := reg.FindAllString(CONTENT, -1)
	log.Print(imgs)

	ids := []string{}
	for _, img := range imgs {
		id := getLastStr(img, 24)
		ids = append(ids, id)

	}

	log.Print(ids)
	//
	//
	//
	//
	//    i := 0
	//    c := cron.New()
	//    spec := "*/1 * * * * ?"
	//    c.AddFunc(spec, func() {
	//        i++
	//        log.Println("cron running:", i)
	//    })
	//    c.Start()

	//    select{}
}

func getLastStr(str string, length int) string {
	rs := []rune(str)
	rl := len(rs)

	if rl <= length {
		return str
	}

	start := rl - length

	return string(rs[start:])
}
