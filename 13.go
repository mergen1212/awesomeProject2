package main

import "fmt"

func main() {
	first, last := "1", "2"
	fmt.Println(first, last)
	last, first = first, last
	fmt.Println(first, last)
}
