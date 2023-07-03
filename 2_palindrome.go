package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isPalindromePhrase(str string) bool {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("masukkan string : ")
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	fmt.Println(isPalindromePhrase(str))

}
