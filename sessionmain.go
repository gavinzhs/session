package main

import (
    "testgo/session"
    _ "testgo/session/driver"
    "log")

func main(){
    print("start")

    manager, err := session.NewManager("memory", "test-session", 3600)
    if err != nil{
        log.Fatalf("NewManager err : %+v", err)
    }

    session := manager.SessionTest()
    session.Set("uid", 123)
    log.Printf("uid:%d, sessionId:%s, 这个没有的是什么呀: %s", session.Get("uid").(int), session.SessionID(), session.Get("me"))
}
