package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var scores []int

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		stack := make([]string, 0)
		valid := true
		for _, char := range strings.Split(scan.Text(), "") {
			// open
			if strings.Contains("([{<", char) {
				switch char {
				case "(":
					stack = append(stack, ")")
				case "[":
					stack = append(stack, "]")
				case "{":
					stack = append(stack, "}")
				case "<":
					stack = append(stack, ">")
				}
				continue
			}
			// close
			if strings.Contains(")]}>", char) {
				want := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if char != want {
					valid = false
					break
				}
				continue
			}
			// other
			log.Fatalf("unknown char %q", char)
		}
		if !valid {
			continue
		}

		score := 0
		for i := 0; i < len(stack); i++ {
			score *= 5
			switch stack[len(stack)-1-i] {
			case ")":
				score += 1
			case "]":
				score += 2
			case "}":
				score += 3
			case ">":
				score += 4
			}
		}
		scores = append(scores, score)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
