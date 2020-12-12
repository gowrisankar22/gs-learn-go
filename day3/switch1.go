package main

import "fmt"

func main()  {

	var day string
	fmt.Println("Enter the day: ")
	fmt.Scanln(&day)

	switch day {

	case "monday":
		fmt.Println("yeah Monday")
		fallthrough
	case "tuesday":
		fmt.Println("yeah Tuesday")
		fallthrough
	default:
		fmt.Println("sorry")



	}
}