package main

import "fmt"

func main() {

	//s := "hello " + "world"
	fmt.Println("hello ", "world")
	fmt.Println(2 + 2)
	fmt.Println(2 - 2)
	fmt.Println(2 / 2)
	fmt.Println(3 % 2)
	a, b := 6, 5
	fmt.Println(a <= b)

	var abc string
	fmt.Println(abc)
	var bcd int
	fmt.Println(bcd)
	var ok bool
	fmt.Println(ok)
	fmt.Printf("int=%d\nstring=%q\nboolean=%t", bcd,abc,ok)
}

