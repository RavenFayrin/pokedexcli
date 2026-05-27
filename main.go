package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			words := cleanInput(text)
			fmt.Println("Your command was:", words[0])
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}