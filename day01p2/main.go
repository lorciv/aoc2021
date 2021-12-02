package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var window [3]int

func sum() int {
	s := 0
	for i := 0; i < len(window); i++ {
		s += window[i]
	}
	return s
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(window); i++ {
		if !scan.Scan() {
			log.Fatal("could not scan initial values")
		}
		n, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		window[i] = n
	}

	last, count := sum(), 0

	for scan.Scan() {
		n, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}

		// left shift
		copy(window[:], window[1:])
		window[len(window)-1] = n

		s := sum()
		if s > last {
			count++
		}
		last = s
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
