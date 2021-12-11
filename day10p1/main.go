package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	score := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		stack := make([]string, 0)
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
					switch char {
					case ")":
						score += 3
					case "]":
						score += 57
					case "}":
						score += 1197
					case ">":
						score += 25137
					}
					break
				}
				continue
			}
			// other
			log.Fatalf("unknown char %q", char)
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
}
