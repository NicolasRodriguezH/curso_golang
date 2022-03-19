package main

import (
	"fmt"
	"strings"
)

func palindromo(text string) {
	var alReves string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		alReves += string(text[i])
	}

	if text == alReves {
		fmt.Println("Es palindrom")
	} else {
		fmt.Println("nockout")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	palindromo("Ojorojo")
}
