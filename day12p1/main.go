package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Key idea: if the current vertex has not been visited IN THE CURRENT PATH

func main() {
	neighbors := make(map[string][]string)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		s := strings.Split(scan.Text(), "-")
		neighbors[s[0]] = append(neighbors[s[0]], s[1])
		neighbors[s[1]] = append(neighbors[s[1]], s[0])
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	var paths [][]string

	queue := [][]string{
		{"start"},
	}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path)-1]

		if node == "end" {
			paths = append(paths, path)
			continue
		}

		for _, neigh := range neighbors[node] {
			// Skip visited nodes if lowercase
			visited := false
			for _, v := range path {
				if strings.ToLower(v) == v && neigh == v {
					visited = true
					break
				}
			}
			if visited {
				continue
			}

			newPath := make([]string, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, neigh)

			queue = append(queue, newPath)
		}
	}

	fmt.Println(len(paths))
}
