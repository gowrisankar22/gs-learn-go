package main

import "fmt"

type name []string
type age  []int

func main ()  {

	var Name name
	var Age age
	Name = append(Name,"DM","ME")
	fmt.Println(len(Name),len(Age))
	//Age = append(Age,1)
	Name = append(Name,"DM")
	fmt.Println(len(Name),len(Age))
	//Age = append(Age,1)
	Name = append(Name,"DM")
	fmt.Println(len(Name),len(Age))
	Name = append(Name,"DM")
	fmt.Println(len(Name),len(Age))
	fmt.Println(Name)
}
