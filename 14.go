package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = make(chan string)
	s := reflect.TypeOf(x).String()
	fmt.Println(s)
}
