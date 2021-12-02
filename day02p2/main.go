package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	pos   = 0
	depth = 0
	aim   = 0
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		cmd := strings.Fields(scan.Text())
		n, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err)
		}

		switch cmd[0] {
		case "forward":
			pos += n
			depth += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pos * depth)
}
