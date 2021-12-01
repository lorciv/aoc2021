package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	last, count := math.MaxInt, 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		n, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		if n > last {
			count++
		}
		last = n
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
