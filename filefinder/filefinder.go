package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main()  {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Println("Please enter the directory name")
		return
	}

	readdir, err := ioutil.ReadDir(args[0])
	if err != nil{
		log.Fatal(err)
	}

	var filesss []byte

	for _,v := range readdir{
		if v.Size() == 0 {
			filesss = append(filesss,v.Name()...)
		}
		err = ioutil.WriteFile(
			"hello.txt",
			filesss,
			666,
		)
		if err != nil{
			log.Fatal(err)
		}
	}
}
