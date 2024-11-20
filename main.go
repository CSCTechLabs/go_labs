package main

import (
	variables "github.com/kushallabs/go_csc_2024/demos/variables"
	variablesBasics "github.com/kushallabs/go_csc_2024/demos/variables/basics"
)

var (
	printMessage string
	a            int
	myInt8       uint8
)

func main() {

	// var flag bool = true
	// var myFloat float32 = 10.5

	// printMessage = "hellow csc"
	// fmt.Println(printMessage)

	// printMessage = "hellow csc1"
	// fmt.Println(printMessage)

	// a = 10
	// fmt.Println(a)

	// myInt8 = 0
	// fmt.Println(myInt8)

	// // print memory of all int types by unsafe.Sizeof(a)
	// fmt.Println("Memory of int8: ", unsafe.Sizeof(int8(10)))
	// fmt.Println("Memory of int16: ", unsafe.Sizeof(int16(10)))
	// fmt.Println("Memory of int32: ", unsafe.Sizeof(int32(10)))
	// fmt.Println("Memory of int64: ", unsafe.Sizeof(int64(10)))
	// fmt.Println("Memory of int: ", unsafe.Sizeof(a))

	// //TODO int vs int64

	// // zero values
	// fmt.Println(flag)
	// fmt.Println(myFloat)

	// call package variables
	variablesBasics.M = 10
	variablesBasics.Basics()

	variables.Conversions()

}
