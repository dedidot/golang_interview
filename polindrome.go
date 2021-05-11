package main
import (
	"fmt"
)
func IsPalindrome(word string) bool {
	var temp = ""
	var newWord = []rune(word)
	for i := len(newWord) - 1; i >= 0; i-- {
		temp += string(newWord[i])	
	}
	return temp == word
}
func IsPalindromeV2(word string) bool {
	var newWord = []rune(word)
	result := true
	for i := 0; i < len(newWord); i++ {
		firstString := string(newWord[i])
		lastString := string(newWord[len(newWord) - i - 1])
		if firstString != lastString {
			result = false
			break
		}
	}
	return result
}
func main() {
	isPal := IsPalindrome("kodok")
	fmt.Println(isPal)
	isPalV2 := IsPalindromeV2("kodok")
	fmt.Println(isPalV2)
}
