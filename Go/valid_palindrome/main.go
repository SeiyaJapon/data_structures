/*
Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Esto funciona pero sería más eficiente con dos punteros desde los extremos hacia el centro
	// El motivo es que al crear una nueva cadena, se está usando espacio adicional y tiempo para copiar los caracteres.
	
	s := "A man, a plan, a canal: Panama"
	//s := "race a car"
	//s := ""
	//s := "No 'x' in Nixon"
	//s := "lalalala"
	//s := "asd,dsa"

	myString := onlyLetters(s)

	fmt.Println(myString)

	result := isPalindrome(myString)

	fmt.Println("Is palindrome: ", result)
}

func onlyLetters(s string) string {
	re := regexp.MustCompile(`[^a-z0-9]+`)
	return re.ReplaceAllString(
		strings.ToLower(s),
		"",
	)
}

func isPalindrome(myString string) bool {
	for i := 0; i < len(myString); i++ {
		if myString[i] != myString[len(myString)-1-i] {
			return false
		}
	}

	return true
}
