package main

import "fmt"

func main() {
	var var1 string
	var var2, var3, var4 string
	var1 = "e"
	var2 = "f"
	var3 = "r"
	var4 = "t"
	var var5, var6 string = "a", "b"
	var name = "gowrisankar"
	var Name, Age, Country = "gowrisankar", 30, "India"
	_,_,_,_,_,_,_,_,_  = var1,var2,var3,var4,var5,var6,Age,Country,Name
	var7, var8 := "test",3
	fmt.Println(Country)
	fmt.Printf("%T,%T",var7,var8 )
	age := 10
	var var9 string = "k"
	var var10 int
	var var11 float32
	fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Errorf(var9)
	fmt.Println("gooo")
	fmt.Scan(var10)
	fmt.Println(var10)
	fmt.Scanf("%f", &var11)
	//fmt.Println(var10,var11)
}