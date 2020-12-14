package main

func main()  {

	var student map[int]string

	student = map[int]string{
		1: "dm",
		2: "dmm",
	}

	class := map[int][]string{
		1: {"test","test"},
	}

	var name = make(map[int][]string)
	//name := make(map[int][]string)

	name = map[int][]string{
		1 : {"f"},
	}

	name[0] = []string{"3434","3434"}
	name[1] = []string{"3434","3434"}


	_,_ = student,class
}