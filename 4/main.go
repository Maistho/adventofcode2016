package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	if len(s[i]) < len(s[j]) {
		return true
	} else if len(s[i]) > len(s[j]) {
		return false
	} else {
		return s[i][0] > s[j][0]
	}
}

func main() {
	var sum = 0
	for {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		room := regexp.MustCompile(`-|\[|\]`).Split(s, -1)
		roomCheckSum := room[len(room)-2]
		roomId, _ := strconv.Atoi(room[len(room)-3])
		roomNameArr := strings.Split(strings.Join(room[:len(room)-3], ""), "")
		sort.Strings(roomNameArr)
		for i := 1; i < len(roomNameArr); i++ {
			if roomNameArr[i-1] != roomNameArr[i] {
				roomNameArr = append(roomNameArr[:i], append([]string{"-"}, roomNameArr[i:]...)...)
				i++
			}
		}
		roomNameArr = strings.Split(strings.Join(roomNameArr, ""), "-")
		sort.Sort(sort.Reverse(ByLength(roomNameArr)))
		var roomName = ""
		for i, c := range roomNameArr {
			if i > 4 {
				break
			}
			roomName += fmt.Sprintf("%c", c[0])

		}
		if roomName == roomCheckSum {
			enc := room[:len(room)-3]

			ri := roomId % 26
			for _, word := range enc {
				for _, c := range word {
					char := int(c) + ri
					if char < int('a') && char > int('Z') || char > int('z') {
						char -= 26
					}
					fmt.Printf("%c", char)
				}
				fmt.Printf(" ")
			}
			fmt.Printf(" - %d\n", roomId)

			sum += roomId
		}
	}
}
