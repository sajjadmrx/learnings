//package xx
//
//import (
//	"fmt"
//)
//
////		Address With &
////	 get Value With *
//var pl = fmt.Println
//
//func updateCount(value *int) {
//	*value = *value + 1
//}
//func main1() {
//	a := 2
//	b := &a
//	*b = 4
//	fmt.Println(a)
//}
//func main2() {
//	count := 1
//	for i := 0; i < 5; i++ {
//		updateCount(&count)
//	}
//	pl(count)
//}
//
//func validatePass(value *string) bool {
//	if len(*value) < 5 {
//		return false
//	}
//	return true
//}
//func main3() {
//	password := "1234"
//	isValid := validatePass(&password)
//	pl("isValid?: ", isValid)
//}
//func saveUsers(value *[]string) {
//	pl("saved", len(*value))
//}
//func main4() {
//	users := []string{}
//
//	users = append(users, "elon")
//	users = append(users, "bob")
//	saveUsers(&users)
//}
//
//type Todo struct {
//	UserId    int    `json:"userId"`
//	TodoId    int    `json:"id"`
//	Title     string `json:"title"`
//	Completed bool   `json:"completed"`
//}
//
//func saveTodos(todos *map[string]Todo) {
//	pl("saved,", *todos)
//}
//
//func main() {
//	todos := map[string]Todo{}
//	todosTitle := [2]string{"learning Go", "learning web3"}
//	for i := 0; i < len(todosTitle); i++ {
//		todos[string(i)] = Todo{TodoId: i, Title: todosTitle[i], UserId: 5, Completed: true}
//	}
//	saveTodos(&todos)
//}
