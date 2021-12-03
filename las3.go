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

	x := "1"
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t%s\n", i+1, x)
		x = lookAndSay(x)
	}
}

type Digit struct {
	value rune
	count int
}

func lookAndSay(x string) string {

	// count base-10 digits in string representation
	var digits []*Digit
	last := '0'
	count := 0
	for _, digit := range x {
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
	asString := ""
	for _, digit := range digits {
		asString += fmt.Sprintf("%d%c", digit.count, digit.value)
	}

	return asString
}
