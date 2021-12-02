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
		case "down":
			depth += n
		case "up":
			depth -= n
		default:
			log.Fatalf("unknown command %q", cmd[0])
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pos * depth)
}
