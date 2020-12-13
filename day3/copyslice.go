package main

import "fmt"

func main()  {

	a := []int{1,2,3,4}

	b := make([]int, 3)

	copy(b,a) // copy will copy only based on the length on make cases.

	fmt.Println(a,b)


}
