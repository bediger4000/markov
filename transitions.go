package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

type letterState int

const (
	Unknown   letterState = iota
	Vowel     letterState = iota
	Consonant letterState = iota
)

func main() {
	ccProb := flag.Float64("cc", 0.500, "C-to-C probability")
	cvProb := flag.Float64("cv", 0.500, "C-to-V probability")
	vcProb := flag.Float64("vc", 0.500, "V-to-C probability")
	vvProb := flag.Float64("vv", 0.500, "V-to-V probability")

	iterations := flag.Int("N", 100, "number of transitions")
	showRatio := flag.Int("I", 10, "how often to show consonant-to-vowel ratio on stderr")

	startStateC := flag.Bool("C", false, "start with consonant")
	startStateV := flag.Bool("V", false, "start with vowel")

	flag.Parse()

	if math.Abs(*ccProb+*cvProb-1.00) > .001 {
		fmt.Fprintf(os.Stderr, "Sum of C out transitions (%.04f) must be 1.00\n", *ccProb+*cvProb)
		return
	}
	if math.Abs(*vcProb+*vvProb-1.00) > .001 {
		fmt.Fprintf(os.Stderr, "Sum of V out transitions (%.04f) must be 1.00\n", *vcProb+*vvProb)
		return
	}

	if (*startStateC && *startStateV) || !(*startStateC && *startStateV) {
		fmt.Fprintf(os.Stderr, "only one of -C or -V can/must be set\n")
	}

	cCount := 0
	vCount := 0

	machineState := Vowel
	if *startStateC {
		machineState = Consonant
	}

	for i := 1; i <= *iterations; i++ {
		if (i % *showRatio) == 0 {
			fmt.Fprintf(os.Stderr, "%d V/C: %.04f\n", i, float64(vCount)/float64(cCount))
		}
		if (i % 80) == 0 {
			fmt.Println()
		}
		p := rand.Float64()
		switch machineState {
		case Vowel:
			fmt.Print("a")
			vCount++
			if p <= *vcProb {
				machineState = Consonant
			} // otherwise state stays Vowel
		case Consonant:
			fmt.Print("b")
			cCount++
			machineState = Vowel
			if p <= *ccProb {
				machineState = Consonant
			}
		}
	}
}
