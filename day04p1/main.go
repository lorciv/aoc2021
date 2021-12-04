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

func main() {
	scan := bufio.NewScanner(os.Stdin)

	// read numbers
	var draws []int
	scan.Scan()
	for _, s := range strings.Split(scan.Text(), ",") {
		d, _ := strconv.Atoi(s)
		draws = append(draws, d)
	}
	//fmt.Println("draws", draws)

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

	// play
	for i := 0; i < len(draws); i++ {
		//fmt.Println("draw:", draws[i])
		over := false
		for _, b := range boards {
			b.mark(draws[i])
			//fmt.Println(b)
			if b.win() {
				//fmt.Println("winner!")
				fmt.Println(b.score(draws[i]))
				over = true
			}
		}
		if over {
			break
		}
	}

}
