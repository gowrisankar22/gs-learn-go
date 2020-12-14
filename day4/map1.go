package main

import "fmt"

func main ()  {

	var end map[string]string = map[string]string{
		"up" : "upper",
		"down" : "below",
	}

	fmt.Println(end)

	stu := map[int]int{

		5 : 3,
		7 : 6,
	}
	fmt.Println(stu)
}
