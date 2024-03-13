package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex)
}

// func NewUser(name, age int, sex int) User {
// 	return User{
// 		name: name,
// 		age:  age,
// 		sex:  sex,
// 	}
// }

func main() {
	// user1 := NewUser("Aza", 20, "Male")
	user := User{"Vasya", 23, "male"}

	// fmt.Println(user1.name)
	fmt.Printf("%+v\n", user.name)
	user.printUserInfo()
}
