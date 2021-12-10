package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func (c coord) sum(d coord) coord {
	return coord{
		x: c.x + d.x,
		y: c.y + d.y,
	}
}

var directions = []coord{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func main() {
	heightMap := make(map[coord]int)

	rows, cols := 0, 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		if cols == 0 {
			cols = len(split)
		}
		for y, c := range split {
			h, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("could not parse height: %v", err)
			}
			heightMap[coord{rows, y}] = h
		}
		rows++
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	count, risk := 0, 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cur, low := coord{i, j}, true
			for _, d := range directions {
				adj, ok := heightMap[cur.sum(d)]
				if !ok {
					adj = 10 // height is always < 10
				}
				if adj <= heightMap[cur] {
					low = false
					break
				}
			}
			if low {
				count++
				risk += 1 + heightMap[cur]
			}
		}
	}

	fmt.Printf("dim/count/risk = %dx%d/%d/%d\n", rows, cols, count, risk)
}
