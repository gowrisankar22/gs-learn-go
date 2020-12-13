package main

import "fmt"

func main ()  {

	var age []int

	age = append(age, 5)
	fmt.Println(len(age), cap(age)) // {1} or {1,3} len and cap will be same. more than 2 odd cap = len+1 even cap =len
	//age = append(age, 5,5)
	//fmt.Println(len(age), cap(age))

}
