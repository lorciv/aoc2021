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

	minCost := math.MaxInt
	for pos := min; pos <= max; pos++ {
		cost := 0
		for i := 0; i < len(positions); i++ {
			cost += fuel(abs(positions[i]-pos), 1, 0)
		}
		//fmt.Printf("Move to %d: %d fuel\n", pos, fuel)

		if cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func fuel(n, cost, tot int) int {
	if n == 0 {
		return tot
	}
	return fuel(n-1, cost+1, tot+cost)
}
