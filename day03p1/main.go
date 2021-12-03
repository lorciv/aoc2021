package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var (
	countZeros []int
	countOnes  []int
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		bits := strings.Split(scan.Text(), "")

		if countZeros == nil {
			countZeros = make([]int, len(bits))
			countOnes = make([]int, len(bits))
		}

		for i, b := range bits {
			switch b {
			case "0":
				countZeros[i]++
			case "1":
				countOnes[i]++
			}
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(countZeros)
	// fmt.Println(countOnes)

	gamma := make([]int, len(countZeros))
	epsilon := make([]int, len(countZeros))
	for i := 0; i < len(countZeros); i++ {
		if countOnes[i] > countZeros[i] {
			gamma[i] = 1
		} else {
			epsilon[i] = 1
		}
	}

	// fmt.Println(gamma)
	// fmt.Println(epsilon)

	g, e := 0, 0
	for i := 0; i < len(countZeros); i++ {
		g += gamma[len(gamma)-1-i] * int(math.Pow(2, float64(i)))
		e += epsilon[len(epsilon)-1-i] * int(math.Pow(2, float64(i)))
	}

	fmt.Println(g * e)
}
