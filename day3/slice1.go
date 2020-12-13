package main

import "fmt"

func main()  {

	//a := []int{1,2,3}
	var a[]int
	fmt.Printf("%p\n", a)
	fmt.Println(a)
	fmt.Println(len(a),cap(a))
	a = append(a, 5,8)
	fmt.Printf("%p\n", a)
	fmt.Printf("%#v", a)
	fmt.Println(len(a),cap(a))

	a=append(a, 5)
	fmt.Printf("%#v", a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)
	a=append(a, 5)
	fmt.Printf("%#v", a)
	fmt.Println(len(a),cap(a))
	fmt.Printf("%p\n", a)

}
