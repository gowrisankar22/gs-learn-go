package main

import "fmt"

func main()  {

	a := []int{1,2,3,4,5,6,7,8,9}
	b := a[1:6]
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))
	b = append(b, 89)
	fmt.Println(a,b)

}