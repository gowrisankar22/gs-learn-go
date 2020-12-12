package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	args := os.Args[1:]
	if len(args) < 3 {
		log.Println("not enough parameters. enter name and age")
	}
	name := args[0]
	age := args[1]
	marks := args[2]
	agestring, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	marksstring, err := strconv.Atoi(marks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name is %T and age is %T marks is %T", name,agestring,marksstring)
}

