package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const nDays = 256

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var fishes [9]int
	for _, s := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		fishes[n]++
	}

	//fmt.Println("Initial state:", fishes)
	for day := 0; day < nDays; day++ {
		creators := fishes[0]
		copy(fishes[:], fishes[1:])

		fishes[6] += creators
		fishes[8] = creators
		//fmt.Printf("After %d days: %v\n", day+1, fishes)
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += fishes[i]
	}
	fmt.Println(sum)
}
