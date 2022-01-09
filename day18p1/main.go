package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lorciv/aoc2021/snailmath"
)

func main() {
	var num *snailmath.Number

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		n, err := snailmath.Parse(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		if num == nil {
			num = n
			continue
		}
		num = snailmath.Add(num, n)
		snailmath.Reduce(num)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(num, "mag", num.Mag())
}
