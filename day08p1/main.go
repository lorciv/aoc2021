package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	count := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		split := strings.Split(scan.Text(), " | ")
		if len(split) != 2 {
			log.Fatal("invalid input line")
		}

		for _, f := range strings.Fields(split[1]) {
			if len(f) == 2 || len(f) == 3 || len(f) == 4 || len(f) == 7 {
				count++
			}
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
