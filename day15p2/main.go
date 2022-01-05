package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	risk        map[coord]int
	best        map[coord]int
	rows, cols  int
	start, dest coord
)

func initialize(in io.Reader) {
	risk = make(map[coord]int)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		split := strings.Split(scan.Text(), "")
		if cols == 0 {
			cols = len(split)
		}
		for i, c := range split {
			val, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			risk[coord{rows, i}] = val
		}
		rows++
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for si := 0; si < 5; si++ {
				for sj := 0; sj < 5; sj++ {
					if si == 0 && sj == 0 {
						continue
					}
					val := risk[coord{i, j}] + (si + sj)
					for val > 9 {
						val -= 9
					}
					risk[coord{i + rows*si, j + cols*sj}] = val
				}
			}
		}
	}

	rows *= 5
	cols *= 5

	start, dest = coord{0, 0}, coord{rows - 1, cols - 1}

	best = make(map[coord]int)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			best[coord{i, j}] = math.MaxInt
		}
	}
	best[start] = 0
}

func display(m map[coord]int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(m[coord{i, j}])
		}
		fmt.Println()
	}
}

type coord struct {
	row, col int
}

func (c coord) valid() bool {
	return c.row >= 0 && c.row < rows && c.col >= 0 && c.col < cols
}

func (c coord) sum(d coord) coord {
	return coord{c.row + d.row, c.col + d.col}
}

func (c coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.row, c.col)
}

var directions = []coord{
	{0, 1},  // right
	{1, 0},  // down
	{-1, 0}, // up
	{0, -1}, // left
}

func main() {
	initialize(os.Stdin)

	queue := []coord{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := cur.sum(dir)
			if !next.valid() {
				continue
			}

			if best[cur]+risk[next] < best[next] {
				best[next] = best[cur] + risk[next]
				queue = append(queue, next)
			}
		}
	}

	fmt.Println(best[dest])
}
