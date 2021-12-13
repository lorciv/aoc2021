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

func (c coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func main() {
	page := make(map[coord]bool)
	width, height := 0, 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		var c coord
		fmt.Sscanf(line, "%d,%d", &c.x, &c.y)
		page[c] = true

		if c.x+1 > width {
			width = c.x + 1
		}
		if c.y+1 > height {
			height = c.y + 1
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(width, height)

	for scan.Scan() {
		instr := strings.Split(scan.Text(), "=")
		fold, _ := strconv.Atoi(instr[1])
		if instr[0] == "fold along y" {
			for k := range page {
				if k.y > fold {
					page[coord{k.x, height - k.y - 1}] = true
				}
			}
			height = fold
		}
		if instr[0] == "fold along x" {
			for k := range page {
				if k.x > fold {
					page[coord{width - k.x - 1, k.y}] = true
				}
			}
			width = fold
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if page[coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
