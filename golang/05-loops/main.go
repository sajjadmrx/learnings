package main

import "fmt"

var pl = fmt.Println

func main() {
	for i := 1; i < 6; i++ {
		pl("i:", i)
	}

	for true {
		pl("print for Ever! 😁")
	}

}
