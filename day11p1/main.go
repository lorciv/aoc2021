package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	rows, cols = 10, 10
)

type grid [rows][cols]int

func (g grid) step() (next grid, flashes int) {
	// increase by 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g[i][j]++
		}
	}
	// flash
	type coord struct {
		x, y int
	}
	flashed := make(map[coord]bool)
	changed := true
	for changed {
		changed = false
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if g[i][j] > 9 && !flashed[coord{i, j}] {
					//fmt.Printf("(%d, %d) flashes\n", i, j)
					flashed[coord{i, j}] = true
					for _, di := range []int{-1, 0, 1} {
						for _, dj := range []int{-1, 0, 1} {
							ni, nj := i+di, j+dj
							if ni == i && nj == j {
								continue
							}
							if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
								continue
							}
							//fmt.Printf("incr (%d, %d)\n", ni, nj)
							g[ni][nj]++
						}
					}
					//fmt.Println(g)
					changed = true
					break
				}
			}
			if changed {
				break
			}
		}
	}
	// set energy to 0 for those who flashed
	for f := range flashed {
		g[f.x][f.y] = 0
	}
	return g, len(flashed)
}

func (g grid) String() string {
	buf := bytes.Buffer{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Fprint(&buf, g[i][j])
		}
		fmt.Fprint(&buf, "\n")
	}
	return buf.String()
}

func main() {
	var start grid

	scan := bufio.NewScanner(os.Stdin)
	row := 0
	for scan.Scan() {
		for col, s := range strings.Split(scan.Text(), "") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse digit: %v", err)
			}
			start[row][col] = n
		}
		row++
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	flashes := 0

	fmt.Println("Before any steps:")
	fmt.Println(start)
	for i := 0; i < 100; i++ {
		var f int
		start, f = start.step()
		flashes += f
		fmt.Printf("After step %d:\n", i+1)
		fmt.Println(start)
	}

	fmt.Println("Flashes:", flashes)
}
