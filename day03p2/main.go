package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

var (
	width int
)

func main() {
	var reports []string

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		rep := scan.Text()
		reports = append(reports, rep)
		if width == 0 {
			width = len(rep)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	oxigenBits, err := find(reports, func(n1, n0 int) byte {
		if n1 >= n0 {
			return '1'
		}
		return '0'
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("oxigen", oxigenBits, bitsToInt(oxigenBits))

	co2Bits, err := find(reports, func(n1, n0 int) byte {
		if n0 > n1 {
			return '1'
		}
		return '0'
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("co2", co2Bits, bitsToInt(co2Bits))

	fmt.Println(bitsToInt(oxigenBits) * bitsToInt(co2Bits))
}

func find(reports []string, f func(int, int) byte) (string, error) {
	filtered := make([]string, len(reports))
	copy(filtered, reports)

	for col := 0; col < width; col++ {

		// find criteria
		count := 0
		for _, f := range filtered {
			if f[col] == '1' {
				count++
			}
		}

		crit := f(count, len(filtered)-count)

		//fmt.Println("col", col, "crit", crit)

		// filter records
		i := 0
		for i < len(filtered) {
			if filtered[i][col] != crit {
				//fmt.Println("remove", filtered[i])
				copy(filtered[i:], filtered[i+1:])
				filtered = filtered[:len(filtered)-1]
				continue
			}
			i++
		}

		// if one left, return
		if len(filtered) == 1 {
			return filtered[0], nil
		}
	}

	return "", errors.New("no record found")
}

func bitsToInt(bits string) int {
	res := 0
	for i := 0; i < len(bits); i++ {
		var n int
		if bits[len(bits)-1-i] == '0' {
			n = 0
		} else {
			n = 1
		}
		res += n * int(math.Pow(2, float64(i)))
	}
	return res
}
