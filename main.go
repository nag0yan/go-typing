package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	answers := []string{"apple", "notify", "maybe", "ask again later"}
	fmt.Println("This is typing game.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		r := rand.Int() % len(answers)
		answer := answers[r]
		fmt.Println(answer)
		for scanner.Scan() {
			str := scanner.Text()
			if str == answer {
				fmt.Println("Correct!")
				break
			} else {
				fmt.Println("Wrong!")
			}
		}
	}
}
