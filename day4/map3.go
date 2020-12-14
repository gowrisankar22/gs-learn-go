package main

import "fmt"

func main()  {
	var student map[int][]string = make(map[int][]string)

	student[0] = []string{"a","b"}

	fmt.Printf("%#v",student)
}
