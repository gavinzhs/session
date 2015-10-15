package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func change(str *string) {
	log.Println(*str)
	*str = "321"
}

type Human struct {
	name string
	int
}

func say(str string) {
	log.Printf("到底谁先进来:%s", str)
	for i := 0; i < 5; i++ {
		log.Printf("sche之前:%s", str)
		runtime.Gosched()
		log.Printf("sche之后:%s", str)
		log.Println(str)
	}
}

type Me struct {
	sync.Mutex // 互斥锁 不能直接解锁 直接解锁会panic 加锁后再次加锁会死锁   正式利用这个机制来做的死锁
	name       string
}

type MeRW struct {
	sync.RWMutex // 读写锁 可以多个读锁 并发读   但是如果写锁进入 会优先级高一些
	name         string
}

func (me *MeRW) ROperation(num int) {
	me.RLock()
	defer me.RUnlock()
	log.Printf("name:%s", me.name)
	time.Sleep(5 * time.Second)
	log.Printf("读操作完毕:%d", num)
}

func (me *MeRW) WOperation(modify string) {
	me.Lock()
	defer me.Unlock()
	log.Printf("修改前name:%s", me.name)
	me.name = modify
	time.Sleep(5 * time.Second)
	log.Println("写操作完毕")
}

func TestMutex(t *testing.T) {
	me := &MeRW{name: "123"}
	log.Printf("name:%s", me.name)
	go me.WOperation("after")
	time.Sleep(time.Second)
	go me.ROperation(1)
	time.Sleep(time.Second)
	go me.ROperation(2)
	time.Sleep(time.Second)
	go me.WOperation("after1")
	time.Sleep(time.Second * 20)
}

func TestTmp(t *testing.T) {
	redisPool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		Wait:        true,
		IdleTimeout: 4 * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", "a1!"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	c := redisPool.Get()
	defer c.Close()

	if i, err := redis.Int(c.Do("GET", "test1")); err != nil {
		log.Printf("err :%+v", err)
		return
	} else {
		log.Printf("value:%d", i)
	}
}
