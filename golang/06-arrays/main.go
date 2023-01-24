package main

import "fmt"

var pl = fmt.Println

func main1() {
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

func main2() {
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 4

	pl(b, a)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	c := a[3:]
	d := a[3:6]
	pl(b)
	pl(c)
	pl(d)
}
