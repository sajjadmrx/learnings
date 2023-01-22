package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var name string = "sajjad"
	var v1, v2 float32 = 1.2, 1.3
	var v3 = "hello"
	v4 := "Hello World"
	v5 := 20
	var isAdmin bool = true
	var age int = 20
	fmt.Println(name, v1, v2, v3, v4, v5, isAdmin, age, reflect.TypeOf(isAdmin))
	price := 2000.1
	fmt.Println("Product Price: ", int(price))
	count := "12"
	countInt, _ := strconv.Atoi(count)
	fmt.Println(countInt)
}
