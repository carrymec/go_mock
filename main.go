package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   string
	Name string `json:"name"`
}

func (u User) Say() {
	fmt.Println("ok")
}

func main() {
	u := User{Name: "chen"}

	uf := reflect.ValueOf(u)
	method := uf.MethodByName("Say")
	method.Call(nil)

	x(0)
	//fmt.Println(r)
}

func x(i int) {
	defer fmt.Println("a")

	defer fmt.Println("b")

	defer func() {
		fmt.Println(1 / i)
	}()
	defer fmt.Println("c")
}

func f() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)

	}
}
