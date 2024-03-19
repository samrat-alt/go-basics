package main

import "fmt"

func main() {

	// var name string = "samrat pandey"

	// fullname := "samrat pandey"

	// var age int
	// if name == "samrat pandey" {
	// 	age = 21

	// } else {
	// 	age = 0
	// }

	// fmt.Printf("\tmy name is %s.\nmy age is %d", name, age)

	names := []string{
		"samrat", "sagar",
	}

	for i, value := range names {
		fmt.Printf("%d. my name is %v\n", i+1, value)
	}

}
