package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	x := 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t%d\n", i+1, x)
		x = lookAndSay(x)
	}
}

func lookAndSay(x int) int {
	asString := strconv.Itoa(x)

	// count base-10 digits in string representation
	digits := make(map[int]int)
	for _, r := range asString {
		digit := int(r - '0')
		digits[digit]++
	}

	// Construct next sequence element as a string
	asString = ""
	for i := 0; i < 10; i++ {
		if digits[i] == 0 {
			continue
		}
		asString += fmt.Sprintf("%d%d", digits[i], i)
	}

	r, err := strconv.Atoi(asString)
	if err != nil {
		log.Fatal(err)
	}

	return r
}
