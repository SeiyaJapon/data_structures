/*
You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.



Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]

Output: [0,9]

Explanation:

The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

Output: []

Explanation:

There is no concatenated substring.

Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

Output: [6,9,12]

Explanation:

The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].



Constraints:

1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
s and words[i] consist of lowercase English letters.
*/

package main

import (
	"fmt"
	"sort"
)

func findSubstring(s string, words []string) []int {
	strings := getAllStrings(words)

	return findStrings(s, strings)
}

func getAllStrings(words []string) []string {
	var results []string

	for index := range words {
		results = append(getWordsFromArray(words, index, results))
	}

	words = doReverse(words)
	for index := range words {
		results = append(getWordsFromArray(words, index, results))
	}

	return results
}

func getWordsFromArray(words []string, index int, results []string) []string {
	wordToFind := ""
	count := index

	for i := 0; i < len(words); i++ {
		wordToFind += words[count]
		count++

		if count == len(words) {
			count = 0
		}
	}

	results = append(results, wordToFind)
	return results
}

func doReverse(s []string) []string {
	var reversed []string

	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	return reversed
}

func findStrings(s string, strings []string) []int {
	var results []int

	for _, theString := range strings {
		if len(theString) > len(s) {
			continue
		}

		index := -1

		for i := 0; i < len(s); i++ {
			if i+len(theString) < len(s) && s[i:i+len(theString)] == theString {
				index = i
				results = append(results, index)
				break
			}
		}
	}

	return results
}

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}

	result := findSubstring(s, words)
	sort.Ints(result)

	fmt.Println(result)
}
