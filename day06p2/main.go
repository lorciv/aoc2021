package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var days = flag.Int("days", 256, "number of days to simulate")

func main() {
	flag.Parse()

	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// [timer value] -> # of fishes
	var fishes [9]int
	for _, s := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fishes[n]++
	}

	//fmt.Println("Initial state:", fishes)
	for day := 0; day < *days; day++ {
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
