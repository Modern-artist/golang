package main

import (
	"fmt"
)

func main() {
	var Str string
	var twoC string
	flag := true
	fmt.Println("Please enter the INPUT:")
	fmt.Scanln(&Str)
	fmt.Println("One's Complement")
	for i := 0; i < len(Str); i++ {
		if Str[i] == '0' {
			fmt.Printf("%c ", '1')
		}
		if Str[i] == '1' {
			fmt.Printf("%c ", '0')
		}
	}
	fmt.Println("")
	fmt.Println("Two's Complement:")
	for i := len(Str) - 1; i >= 0; i-- {
		if flag {
			if Str[i] == '0' {
				twoC = "0 " + twoC
			}
			if Str[i] == '1' {
				twoC = "1 " + twoC
				flag = false
			}
		} else {
			if Str[i] == '0' {
				twoC = "1 " + twoC
			}
			if Str[i] == '1' {
				twoC = "0 " + twoC
			}
		}

	}
	fmt.Println(twoC)

}
