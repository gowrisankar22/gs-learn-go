package main

import "fmt"

func main ()  {

	a := [2][3]int{
		{1,2},
		{1,2,4},
	}

	//fmt.Println(a)

	for _, v := range a{
		fmt.Println(v)
		for _, j:= range v {
			fmt.Println(j)
		}
	}
	s := "test hello saasd adsad"

	for _,v := range s {
		//fmt.Println(v)
		fmt.Printf("%c\n",v)
	}

}
