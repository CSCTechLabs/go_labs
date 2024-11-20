package variables

import "fmt"

var M int

func Basics() {
	i := 10
	BasicsLocal()
	fmt.Println(i)
}

func BasicsLocal() {
	fmt.Println("calling basics from variables/basics folders")
}
