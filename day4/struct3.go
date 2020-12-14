package main

import "fmt"

type school struct {
	name []string
	add string
	pincode int
}
func main()  {

	schools := []school{
		{
			name: []string{"dm"},
			add: "test",
			pincode: 639,
		},
		{
			name: []string{"dm"},
			add: "test",
			pincode: 639,
		},
	}

	fmt.Println(schools)

	fmt.Println(schools[0].name	)
}
