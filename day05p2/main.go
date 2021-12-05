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

func (c coord) add(step coord) coord {
	return coord{
		x: c.x + step.x,
		y: c.y + step.y,
	}
}

func main() {
	overlaps := make(map[coord]int)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		var from, to coord
		fmt.Sscanf(scan.Text(), "%d,%d -> %d,%d", &from.x, &from.y, &to.x, &to.y)

		step := coord{
			x: to.x - from.x,
			y: to.y - from.y,
		}
		if step.x != 0 {
			step.x /= abs(to.x - from.x)
		}
		if step.y != 0 {
			step.y /= abs(to.y - from.y)
		}

		cur := from
		for cur != to {
			overlaps[cur]++
			cur = cur.add(step)
		}
		// last cell included
		overlaps[cur]++
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
