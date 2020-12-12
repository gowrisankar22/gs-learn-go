package main

import "fmt"

var (
	name   string  = "dm"
	age    int     = 1
	height float32 = 4.1
	hello1         = "test"
)

func main() {
	fmt.Println("hello", hello1)
	_, _, _ = name, age, height

}
