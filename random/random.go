package random

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

var Word = randomWord()

func words() string {
	data, err := os.ReadFile("./random/words.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(data)
}

func randomWord() string {
	n := rand.IntN(313)
	word := ""
	str := strings.Split(words(), "\n")

	for i, elem := range str {
		if i == n {
			word = elem
		}
	}

	return word
}
