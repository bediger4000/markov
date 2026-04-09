package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

const (
	Unknown = iota
	Vowel
	Consonant
)

func main() {

	dotOut := false
	fileArg := 1
	if len(os.Args) > 1 && os.Args[1] == "-g" {
		dotOut = true
		fileArg = 2
	}

	fin := os.Stdin
	if len(os.Args[fileArg:]) > 0 {
		var err error
		if fin, err = os.Open(os.Args[fileArg]); err != nil {
			log.Fatal(err)
		}
		defer fin.Close()
	}

	scanner := bufio.NewScanner(fin)

	lineCounter := 0

	letterCount := 0
	vowelCount := 0
	consonantCount := 0
	ccCount := 0
	cvCount := 0
	vvCount := 0
	vcCount := 0

	lastLetter := Unknown

	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()

		for _, r := range line {
			if !unicode.IsLetter(r) {
				continue
			}
			letterCount++
			r = unicode.ToLower(r)
			currentLetter := Unknown
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
				// vowel
				vowelCount++
				currentLetter = Vowel
				switch lastLetter {
				case Vowel:
					vvCount++
				default:
					vcCount++
				}
			default:
				// consonant
				consonantCount++
				currentLetter = Consonant
				switch lastLetter {
				case Vowel:
					cvCount++
				default:
					ccCount++
				}
			}
			lastLetter = currentLetter
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("problem line %d: %v", lineCounter, err)
	}

	fLetterCount := float64(letterCount)

	if dotOut {
		fmt.Printf("/*\n")
	}

	fmt.Printf("total letters %d\n", letterCount)
	fmt.Printf("consonants %d (%.04f)\n", consonantCount, float64(consonantCount)/fLetterCount)
	fmt.Printf("vowels     %d (%.04f)\n\n", vowelCount, float64(vowelCount)/fLetterCount)

	cCount := float64(ccCount + cvCount)
	vCount := float64(vcCount + vvCount)

	fmt.Printf("CC %d (%.04f)\n", ccCount, float64(ccCount)/cCount)
	fmt.Printf("CV %d (%.04f)\n", cvCount, float64(cvCount)/cCount)
	fmt.Printf("VV %d (%.04f)\n", vvCount, float64(vvCount)/vCount)
	fmt.Printf("VC %d (%.04f)\n", vcCount, float64(vcCount)/vCount)
	fmt.Printf("%d total pairs\n", ccCount+cvCount+vvCount+vcCount)

	if dotOut {
		fmt.Printf("*/\n")
		fmt.Printf("digraph g {\nrankdir = \"LR\";\n")
		fmt.Printf("C -> C [label=\"%.04f\"];\n", float64(ccCount)/cCount)
		fmt.Printf("C -> V [label=\"%.04f\"];\n", float64(cvCount)/cCount)
		fmt.Printf("V -> V [label=\"%.04f\"];\n", float64(vvCount)/vCount)
		fmt.Printf("V -> C [label=\"%.04f\"];\n", float64(vcCount)/vCount)
		fmt.Println("};")
	}
}
