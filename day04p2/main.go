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
	rows, cols = 5, 5
)

type board [rows][cols]int

func (b *board) String() string {
	buf := bytes.Buffer{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Fprintf(&buf, "%3d", b[i][j])
		}
		fmt.Fprintf(&buf, "\n")
	}
	return buf.String()
}

func (b *board) mark(n int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b[i][j] == n {
				b[i][j] = -1
			}
		}
	}
}

func (b *board) win() bool {

	// check for complete rows
	for i := 0; i < rows; i++ {
		complete := true
		for j := 0; j < cols; j++ {
			if b[i][j] != -1 {
				complete = false
				break
			}
		}
		if complete {
			return true
		}
	}

	// check for complete cols
	for j := 0; j < cols; j++ {
		complete := true
		for i := 0; i < rows; i++ {
			if b[i][j] != -1 {
				complete = false
				break
			}
		}
		if complete {
			return true
		}
	}

	return false
}

func (b *board) score(last int) int {
	s := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b[i][j] != -1 {
				s += b[i][j]
			}
		}
	}
	return s * last
}

type drawer struct {
	values []int
	cur    int
}

func (d *drawer) draw() int {
	d.cur++
	return d.values[d.cur]
}

func (d *drawer) last() int {
	return d.values[d.cur]
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	// read numbers
	drawer := drawer{
		cur: -1,
	}
	scan.Scan()
	for _, s := range strings.Split(scan.Text(), ",") {
		d, _ := strconv.Atoi(s)
		drawer.values = append(drawer.values, d)
	}

	// read boards
	var boards []*board
	for scan.Scan() {
		if scan.Text() == "" {
			continue
		}

		var b board
		for i := 0; i < rows; i++ {
			fields := strings.Fields(scan.Text())
			if len(fields) != cols {
				log.Fatal("wrong number of columns")
			}
			for j := 0; j < cols; j++ {
				n, _ := strconv.Atoi(fields[j])
				b[i][j] = n
			}
			scan.Scan()
		}
		boards = append(boards, &b)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	winners, winCount := make([]bool, len(boards)), 0

	// play until one board left
	for winCount < len(boards)-1 {
		draw := drawer.draw()
		fmt.Println("draw:", draw)

		for ib, b := range boards {
			if winners[ib] {
				continue
			}
			b.mark(draw)
			fmt.Println(b)
			if b.win() {
				fmt.Println("winner!")
				winners[ib] = true
				winCount++
			}
		}
	}

	// continue until last board wins
	var lastBoard *board
	for i, b := range boards {
		if !winners[i] {
			lastBoard = b
		}
	}
	for !lastBoard.win() {
		lastBoard.mark(drawer.draw())
	}

	fmt.Println(lastBoard.score(drawer.last()))
}
