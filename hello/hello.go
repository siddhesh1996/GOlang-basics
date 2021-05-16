package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("siddhesh")
	fmt.Println(message)

}
