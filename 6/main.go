package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   rune
	Value int
}
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	messages := []string{}
	for {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}

		messages = append(messages, s)
	}
	messageMap := make([]map[rune]int, len(messages[0]))
	for i, _ := range messages[0] {
		aMap := map[rune]int{}
		messageMap[i] = aMap
	}
	for _, msg := range messages {
		for i, c := range msg {
			messageMap[i][c]++
		}
	}
	for _, char := range messageMap {
		count := make(PairList, len(char))
		i := 0
		for k, v := range char {
			count[i] = Pair{k, v}
			i++
		}
		sort.Sort(count)
		fmt.Printf("%c", count[0].Key)
	}
}
