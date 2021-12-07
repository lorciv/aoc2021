package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var positions []int
	min, max := math.MaxInt, math.MinInt
	for _, s := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not parse position: %v", err)
		}
		positions = append(positions, n)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	//fmt.Println(positions, max, min)

	minFuel := math.MaxInt
	for pos := min; pos <= max; pos++ {
		fuel := 0
		for i := 0; i < len(positions); i++ {
			fuel += abs(positions[i] - pos)
		}
		//fmt.Printf("Move to %d: %d fuel\n", pos, fuel)

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
