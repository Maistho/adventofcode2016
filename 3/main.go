package main

import (
	"fmt"
)

func isValid(a, b, c int) bool {
	return a+b > c &&
		a+c > b &&
		b+c > a
}

func main() {
	var valid = 0
	for {
		var a1, a2, a3, b1, b2, b3, c1, c2, c3 int
		_, err := fmt.Scan(&a1, &b1, &c1, &a2, &b2, &c2, &a3, &b3, &c3)
		if err != nil {
			break
		}
		if isValid(a1, a2, a3) {
			valid++
		}
		if isValid(b1, b2, b3) {
			valid++
		}
		if isValid(c1, c2, c3) {
			valid++
		}

	}
	fmt.Print(valid)
}
