package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}
	var i = 0
	var passwordLength = 0
	var password = strings.Split("********", "")
	for {
		hash := GetMD5Hash(fmt.Sprintf("%s%d", s, i))
		i++
		if hash[:5] == "00000" {
			//fmt.Println(hash, hash[5:6], hash[6:7])
			pos, err := strconv.Atoi(hash[5:6])
			if err != nil {
				continue
			}
			if pos >= len(password) {
				continue
			}
			if password[pos] == "*" {
				password[pos] = hash[6:7]
				fmt.Println(strings.Join(password, ""))
				passwordLength++
				if passwordLength > 7 {
					break
				}
			}
		}
	}
	fmt.Println()
}
