package main

import "fmt"

type money int

func main ()  {
	var rupee money = 10
	var i int = int(rupee)
	fmt.Println(i)

	fmt.Println(rupee)

}
