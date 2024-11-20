package types

import "fmt"

func MyType() {
	type (
		myInt     int   // creating a new data type
		myIntCopy = int // alias
	)

	var iMyInt myInt = 10
	var jMyIntCopy myIntCopy = 10

	// try to add myInt to int

	kInt := 10
	kInt = int(iMyInt)
	fmt.Println("value of kInt before copying from alias", kInt)

	kInt = jMyIntCopy

	fmt.Println(iMyInt, jMyIntCopy, kInt, "int my types")
}
