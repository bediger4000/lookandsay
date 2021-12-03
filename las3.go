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

func lookAndSay(x string) string {

	asString := ""
	last := '0'
	count := 0
	for _, digit := range x {
		if digit != last {
			if count != 0 {
				asString += fmt.Sprintf("%d%c", count, last)
			}
			count = 0
		}
		last = digit
		count++

	}
	if count != 0 {
		asString += fmt.Sprintf("%d%c", count, last)
	}

	return asString
}
