package main

import (
	"fmt"
	"log"
	"testing"
)

func TestInterface(t *testing.T) {
	log.Println("test interface")

	type test struct {
		Name string
		age  int
	}

	add := func(num ...interface{}) {
		log.Println("add")
	}

	slice := []*test{&test{"asai", 28}, &test{"asai", 29}, &test{"asai", 30}}
	//
	//
	add(slice)
	add([]*test{&test{"asai", 28}})

}

type square struct {
	r int
}
type circle struct {
	r int
}

func TestHttp1(t *testing.T) {

}

func (s square) area() int {
	return s.r * s.r
}
func (c circle) area() int {
	return c.r * 3
}

func TestInterface2(t *testing.T) {

	log.Println(fmt.Sprintf("\\%s", "+123"))

	s := square{1}
	c := circle{1}
	a := [2]interface{}{s, c}
	fmt.Println(s, c, a)

	sum := 0
	for _, t := range a {
		switch v := t.(type) {
		case square:
			sum += v.area()
		case circle:
			sum += v.area()
		}
	}
	fmt.Println(sum)
}
