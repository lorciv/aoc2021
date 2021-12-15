package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type rule struct {
	pair   [2]string
	insert string
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	if !scan.Scan() {
		log.Fatal("missing template line")
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	template := strings.Split(scan.Text(), "")

	// fmt.Println("template", template)

	var rules []rule
	for scan.Scan() {
		text := scan.Text()
		if text == "" {
			continue
		}

		split := strings.Split(text, " -> ")
		if len(split) != 2 {
			log.Fatalf("invalid rule format %q", text)
		}

		rules = append(rules, rule{
			pair:   [2]string{split[0][:1], split[0][1:]},
			insert: split[1],
		})
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	for step := 0; step < 10; step++ {
		for i := 0; i < len(template)-1; i++ {
			found := false
			for _, r := range rules {
				if template[i] == r.pair[0] && template[i+1] == r.pair[1] {
					template = append(template, "")
					copy(template[i+2:], template[i+1:])
					template[i+1] = r.insert

					// fmt.Println("apply", r, "get", template)

					found = true
					break
				}
			}
			if found {
				i++
			}
		}
	}

	// fmt.Println(len(template))

	count := make(map[string]int)
	for _, el := range template {
		count[el]++
	}
	// fmt.Println(count)

	min, max := math.MaxInt, math.MinInt
	for _, v := range count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println(max - min)
}
