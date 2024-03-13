package main

import "fmt"

func main() {
	defer printMessage()
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
	}
	fmt.Println(messages)
	messages[3] = "mess 3"

	// defer printMessage()
	fmt.Println("main")

}

func printMessage() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

}

// 	for _, message := range messages {
// 		fmt.Println(message)
// 	}
// 	counter := 0
// 	for {
// 		if counter == 100 {
// 			break
// 		}
// 		counter++
// 		fmt.Println(counter)
// 	}
// }
