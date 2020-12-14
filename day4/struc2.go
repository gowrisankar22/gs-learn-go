package main

import "fmt"

type address []string

type name struct {
	name string
	age int
	address
}


func main()  {


names := name{
	name:    "sasd",
	age:     0,
	address: address{"tes","test"},
}

for _, v := range names.address {

	fmt.Println(v)

}
fmt.Println(names)

}