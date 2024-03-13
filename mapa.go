package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya": 15,
		"Petya": 23,
		"Aza":   20,
	}

	age, exists := users["cas"]
	if exists {
		fmt.Println("Petya", age)
	} else {
		fmt.Println("Usera net in spis")
	}

}
