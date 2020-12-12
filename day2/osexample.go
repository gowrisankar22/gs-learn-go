package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main ()  {
	//
	//name := os.Args[1]
	//age := os.Args[2]
	//agestring := string(age)
	//
	//fmt.Printf("%T %T", name, agestring)

	args := os.Args[1:]

	if len(args) <2 {
		//fmt.Println("enter age and name")
		log.Println("enter age and name")
		return
	} else {

		name := args[0]
		agestring := args[1]
		ageint, err := strconv.Atoi(agestring)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(name)
		fmt.Println(ageint)
	}


}
