package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our pizza:")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error : %v", err)
	} else {
		fmt.Printf("Thanks for the rating %s", input)
	}

}
