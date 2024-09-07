package main

import "fmt"

var name string = "Boris"
var kek *string = &name
var kek2 **string = &kek

func test(name *string) *string {
	fmt.Println(name)
	fmt.Println(*name)
	*name = "NotBoris"
	return name
}

func test2(name **string) **string {
	fmt.Println("check", *name)
	**name = "Hello"

	return name
}

func main() {
	ptr := test(&name)
	test2(&ptr)
	fmt.Println(name)
}
