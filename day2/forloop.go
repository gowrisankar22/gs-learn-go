package main

import (
	"fmt"
	"os"
)

func main() {
	//i := 5
	//for i=0;i < 5; i++ {
	//	fmt.Println(i)
	//}
	//for i=0;i < 5; i++ {
	//	fmt.Println(i)
	//}
	//j := 0
	//for j < 10{
	//	sum += j
	//	fmt.Println(j)
	//}
	//i := 0
	//for i < 3 {
	//	i += 2
	//	// i+2
	//}
	//fmt.Println(i)
	//var a = [5]int{1,3,3,3,3}
	//for v , b:= range a {
	//	fmt.Println(v,b)
	//
	//}

	// i+=2 // i = i+2
	//i*=2 // i = i*2

	args := os.Args
	for _, v := range args {
		//fmt.Println(i,v)
		if v == "hello" {
			fmt.Println("asdasdasd")
			continue
			fmt.Println("not in here")
		}
	}
}
