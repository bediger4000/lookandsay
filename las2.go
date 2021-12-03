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

type Digit struct {
	value int
	count int
}

func lookAndSay(x int) int {
	asString := strconv.Itoa(x)

	// count base-10 digits in string representation
	var digits []*Digit
	last := 0
	count := 0
	var digit int
	for _, r := range asString {
		digit = int(r - '0')
		if digit != last {
			if count != 0 {
				digits = append(digits, &Digit{value: last, count: count})
			}
			count = 0
		}
		last = digit
		count++

	}
	if count != 0 {
		digits = append(digits, &Digit{value: last, count: count})
	}

	// Construct next sequence element as a string
	asString = ""
	for _, digit := range digits {
		asString += fmt.Sprintf("%d%d", digit.count, digit.value)
	}

	r, err := strconv.Atoi(asString)
	if err != nil {
		log.Fatal(err)
	}

	return r
}
