package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

var (
	rows, cols = 0, 0
	heightMap  = make(map[coord]int)
	visited    = make(map[coord]bool)
)

func explore(start coord) int {
	queue := []coord{start}
	count := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		val, ok := heightMap[cur]
		if !ok {
			val = 9
		}
		if val == 9 {
			continue
		}

		if visited[cur] {
			continue
		}
		visited[cur] = true

		count++

		for _, dir := range directions {
			queue = append(queue, cur.sum(dir))
		}
	}

	return count
}

func main() {
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

	var sizes []int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cur := coord{i, j}
			size := explore(cur)
			if size > 0 {
				sizes = append(sizes, size)
			}
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	if len(sizes) < 3 {
		log.Fatal("not enough basins")
	}
	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}
