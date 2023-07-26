package main

import "fmt"

func main() {
	println("Consuctive A")
	var Str string
	fmt.Println("Please enter the INPUT:")
	fmt.Scanln(&Str)
	for i := 0; i < len(Str); i++ {
		i++
		fmt.Println(i)
	}
}
