package main

//import (
//	"fmt"
//	"testing"
//)
//
//func TestDesEncrypt(t *testing.T) {
//	key := []byte(DES_KEY)
//	origtext := []byte("hello world oh yeah desc encrypt")
//	erytext, err := DesEncrypt(origtext, key)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Printf("erytext: %v\n", erytext)
//
//	destext, err := DesDecrypt(erytext, key)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Println("destext string:", string(destext))
//	fmt.Println("len eq:", len(origtext), len(string(destext)))
//	fmt.Println("string eq:", string(origtext) == string(destext))
//}
