package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
)

func main() {
	consProb := flag.Float64("cons", 0.500, "consonant probability")
	totalLetters := flag.Int("letters", 2000, "number of letters out")
	flag.Parse()

	for i := 0; i < *totalLetters; i++ {
		p := rand.Float64()
		c := 'a' // 'a' is a vowel
		if p <= *consProb {
			c = 'b' // 'b' is a consonant
		}
		if i%80 == 0 {
			fmt.Printf("%c\n", c)
		} else {
			fmt.Printf("%c", c)
		}
	}
}
