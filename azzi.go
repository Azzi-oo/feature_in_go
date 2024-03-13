package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := [3]string{"1", "2", "3"}
	// messages := []string{"1", "2", "3"}
	messages := make([]string, 5, 15)
	messages = append(messages, "6")
	// messages[0] = "1"
	// printMessages(messages)
	fmt.Println(messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))

}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	messages[1] = "5"
	fmt.Println(messages)
	return nil
}
