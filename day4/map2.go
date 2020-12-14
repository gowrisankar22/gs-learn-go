package main

import "fmt"

type names []string

func main()  {

	class := map[int][]string{
		2 : {"fm","test"},
		4 : {"fm","test","tesr2","dm"},
	}
	fmt.Println(class)
	fmt.Println(class[2])
	fmt.Println("******")
	for _, v := range class {
		fmt.Println(v)

	}
}
