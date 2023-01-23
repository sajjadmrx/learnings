package main

import "fmt"

var pl = fmt.Println

func main() {
	var userIds [2]int
	userIds[0] = 1
	userIds[1] = 2
	pl("Array Length", len(userIds))
	userRoles := [2]string{"ADMIN", "USER"}
	pl("User Roles:")
	for i, v := range userRoles {
		fmt.Printf("%d : %d", i, v)
	}
	var roles []string
	roles = append(roles, "ADMIN", "USER")
	pl("\n")
	pl(len(roles))
}
