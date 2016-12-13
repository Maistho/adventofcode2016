package main

import (
	"fmt"
)

func getKey(x int, y int, keypad [][]rune) rune {
	if x < 0 || y < 0 {
		return '-'
	}
	if y >= len(keypad) || x >= len(keypad[y]) {
		return '-'
	}
	return keypad[y][x]
}

func main() {
	var keypad = [][]rune{
		[]rune{'-', '-', '1'},
		[]rune{'-', '2', '3', '4'},
		[]rune{'5', '6', '7', '8', '9'},
		[]rune{'-', 'A', 'B', 'C'},
		[]rune{'-', '-', 'D'},
	}
	var x = 0
	var y = 2
	for {
		var d string
		_, err := fmt.Scan(&d)
		if err != nil {
			break
		}
		for _, c := range d {
			switch c {
			case 'U':
				if getKey(x, y-1, keypad) != '-' {
					y -= 1
				}
			case 'D':
				if getKey(x, y+1, keypad) != '-' {
					y += 1
				}
			case 'L':
				if getKey(x-1, y, keypad) != '-' {
					x -= 1
				}
			case 'R':
				if getKey(x+1, y, keypad) != '-' {
					x += 1
				}
			}
		}
		fmt.Printf("%c", getKey(x, y, keypad))
	}

}
