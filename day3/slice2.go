package main

import "fmt"

func main()  {

	var a []string
	b := []string{"hello", "gowrisankar", "dhruv"}
	a = append(a,"gowriankar")
	b = append(b,"gowriankar")
	// fmt.Println(a,b)

	for _,v := range a {
		fmt.Println(v)
	}
	for _,v := range b {
		fmt.Println(v)
	}



}
