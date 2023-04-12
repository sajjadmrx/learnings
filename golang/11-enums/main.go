package main

import (
	messages "enums/enums/messages"
	permissions "enums/enums/permissions"
	"fmt"
)

//func main1() {
//
//	const (
//		A = iota
//		B
//		C
//		D = B + C
//		E
//		F
//		G = iota
//		H
//		I = H
//		J
//		K
//	)
//	fmt.Println(A, B, C, D, E, F, G, H, I, J, K)
//}

func main() {

	fmt.Println(permissions.ADMIN)
	fmt.Println(messages.CREATED)
	fmt.Println(messages.USERNAME_IS_STRING)
}
