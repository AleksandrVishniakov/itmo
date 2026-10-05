package main

import (
	"fmt"
	"lab2/hamming"
	"log"
)

func main() {
	var input string
	fmt.Print("Enter a Hamming's code (7;4): ")
	fmt.Scanln(&input)

	hamming := hamming.NewHammingEncoder(4)
	ans, err := hamming.Decode(input)
	if err != nil {
		log.Fatalf("decoding error: %s", err.Error())
	}

	if !ans.Valid {
		fmt.Printf("Error found at %d bit\nCorrected string: %s\n", ans.Syndrome, ans.ValidString)
	}

	p := len(ans.Value) - 1
	n := 0
	for i := range ans.Value {
		n += int(ans.Value[i]-'0') * (1 << p)
		p -= 1
	}

	fmt.Printf("Result: %s (bin) = %d (dec)\n", ans.Value, n)
}
