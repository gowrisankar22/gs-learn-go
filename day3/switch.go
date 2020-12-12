package main

import (
	"fmt"
	"os"
)

func main ()  {

	args := os.Args[1:]

	switch args[0] {
	case "monday":
		fmt.Println("Monday")
	case "tuesday":
		fmt.Println("Tuesday")
	default:
		fmt.Println("Sorry")
	}
}
