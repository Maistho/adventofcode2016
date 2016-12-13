package main

import "fmt"

func main() {
	ssl := 0
	for {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		valid := false
		inBrackets := false
		abas := make(map[string]int)
		for i := 0; i < len(s); i++ {
			if s[i] == '[' {
				inBrackets = true
				continue
			}
			if s[i] == ']' {
				inBrackets = false
				continue
			}
			if i < 2 {
				continue
			}

			if s[i] == s[i-2] && s[i] != s[i-1] {
				if inBrackets {
					str := fmt.Sprintf("%c%c", s[i-1], s[i])
					if abas[str] == 2 {
						valid = true
						break
					}
					abas[str] = 1
				} else {
					str := fmt.Sprintf("%c%c", s[i], s[i-1])
					if abas[str] == 1 {
						valid = true
						break
					}
					abas[str] = 2
				}
			}
		}
		if valid {
			ssl++
		}
	}
	fmt.Print(ssl)
}
