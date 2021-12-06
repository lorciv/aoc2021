package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const nDays = 80

func main() {
	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var timers []int
	for _, s := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		timers = append(timers, n)
	}

	//fmt.Println("Initial state:", timers)
	for day := 0; day < nDays; day++ {
		lim := len(timers)
		for i := 0; i < lim; i++ {
			timers[i]--
			if timers[i] == -1 {
				timers[i] = 6
				timers = append(timers, 8)
			}
		}
		//fmt.Printf("After %d days: %v\n", day+1, timers)
	}

	fmt.Println(len(timers))
}
