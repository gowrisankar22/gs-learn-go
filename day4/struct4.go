package main

import "fmt"

type sal struct {
	pf int
	basepay int

}

// Fulltimeemployee simple type struct
type Fulltimeemployee struct {
	name string
	emptId int
	workinghours int
	sal

}

type Partimeemployee struct {
	name string
	emptId int
	duration int
	sal // anonymous field

}

func main()  {

	fulltime := Fulltimeemployee{
		name:         "wew",
		emptId:       555,
		workinghours: 4,
		sal:          sal{pf: 5, basepay: 6	},
	}

	partime := Partimeemployee{
		name:     "erwer",
		emptId:   234,
		duration: 5,
		sal:          sal{pf: 6, basepay: 7	},
	}
	fmt.Println(fulltime,partime)
}
