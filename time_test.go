package main

import (
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	log.Println("到了这里了")
	log.Println(time.Date(2015, time.October, 1, 1, 11, 0, 0, time.Local).AddDate(0, 1, 0).Format("2006-01-02aaa15:04:05bbb07:00"))
	log.Println(time.ParseDuration("72h3m0.5s"))
	log.Println(time.Now().AddDate(0, 1, 0).Unix())
	log.Println(time.Unix(1444120312, 0))

}
