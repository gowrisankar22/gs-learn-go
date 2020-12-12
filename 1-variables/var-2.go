package main

import "fmt"

func main ()  {

	var name = "dm"
	var age = 1.2
	name, age = "mithran", 30
	fmt.Println(name,age)

	const(
		nameop = "this is an error"
		jockey = "hellow"
	)
	fmt.Println(nameop,jockey)
}
