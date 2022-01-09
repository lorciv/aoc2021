package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/lorciv/aoc2021/snailmath"
)

func main() {
	var nums []string

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		nums = append(nums, scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			ni, err := snailmath.Parse(nums[i])
			if err != nil {
				log.Fatal(err)
			}
			nj, err := snailmath.Parse(nums[j])
			if err != nil {
				log.Fatal(err)
			}
			sum := snailmath.Add(ni, nj)
			snailmath.Reduce(sum)
			m := sum.Mag()
			if m > max {
				max = m
			}
		}
	}
	fmt.Println(max)
}
