package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

func main() {
	overlaps := make(map[coord]int)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		var from, to coord
		fmt.Sscanf(scan.Text(), "%d,%d -> %d,%d", &from.x, &from.y, &to.x, &to.y)

		if from.x == to.x {
			// horizontal
			if from.y > to.y {
				from, to = to, from
			}
			for y := from.y; y <= to.y; y++ {
				overlaps[coord{from.x, y}]++
			}
		} else if from.y == to.y {
			// vertical
			if from.x > to.x {
				from, to = to, from
			}
			for x := from.x; x <= to.x; x++ {
				overlaps[coord{x, from.y}]++
			}
		}
	}
	if scan.Err() != nil {
		log.Fatal(scan.Err())
	}

	count := 0
	for _, v := range overlaps {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)
}
