package main

import (
	"fmt"
	"strings"
)

/*
Question
Given an array, words, of n word strings (words[0], words[1],..., words[n-1]), choose a word from it and, in each step, remove a single letter from the chosen word if and only if doing so yields another word that is already in the library. Each successive character removal should be performed on the result of the previous removal, and you cannot remove a character if the resulting string is not an element in words(see Explanation below for detail). The length of a string chain is the maximum number of strings in a chain of successive character removals.

Complete the longestChain function in your editor. It has 1 parameter: an array of n strings, words, where the value of each element words; (where 0 <= i < n) is a word. It must return single integer denoting the length of the longest possible string chain in words.

Input Format
The locked stub code in your editor reads the following input from stdin and passes it to your function: The fist line contains an integer. n, the size of the words array. Each line i of the n subsequent lines (where 0 <= i < n) contains an integer describing the respective strings in words.

Constraints
1 <= n <= 50000

1 <= |words_i| <= 50, where 0 <= i < n

Each string in words is composed of lowercase ASCII letters.

Output Format
Your function must return a single integer denoting the length of the longest chain of character removals possible.

Sample Input 1
6
a
b
ba
bca
bda
bdca
Sample Output 1
4

Explanation
Sample Case 1: words = {"a", "b", "ba", "bca", "bda", "bdca"} Because "a" and "b" are single-character words, we cannot remove any characters from them as that would result in the empty string (which is not an element in words), so the length for both of these string chains is 1.

The word "ba" can create two different string chains of length 2 ("ba" -> "a" and "ba" -> "b"). This means our current longest string chain is 2.

The word "bca" can create two different string chains of length 3 ("bca" -> "ba" -> "a" and "bca" -> "ba" -> "b"). This means our current longest string chain is now 3.

The word "bda" can create two different string chains of length 3 ("bda" -> "ba" -> "a" and "bda" -> "ba" -> "b"). This means our current longest string chain is now 3.

The word "bdca" can create four different string chains of length 4 ("bdca" -> "bda" -> "ba" -> "a" , "bdca" -> "bda" -> "ba" -> "b", "bdca" -> "bca" -> "ba" -> "a", "bdca" -> "bca" -> "ba" -> "b"). This means our current longest string chain is now 4.
*/

/*
func longestChain(words []string) int {
	for k, v := range words {

	}
	return 1
}

func procStr(s string) []string {
	return nil
}
*/
var stacks map[string]int
var pcnt int
var st []int

func RemoveChar(word string, index int) string {
	s := strings.Split(word, "")
	fmt.Printf("ini %v", s)
	x := strings.Join(append(s[:index], s[index+1:]...), "")
	fmt.Println("itu " + x)
	return x
}

func isExist(s string, libs []string) bool {
	for _, v := range libs {
		if s == v {
			return true
		}
	}
	return false
}

func RecursiveChain(s string, c int, libs []string) {

	c = c + 1

	for idx := 0; idx < len(s); idx++ {
		v := RemoveChar(s, idx)
		if isExist(v, libs) {
			RecursiveChain(v, c, libs)
		}
	}

	if pcnt < c {
		pcnt = c
	}
	c = 1

	return
}

func main() {
	x := []string{"a", "b", "ba", "bd", "dc", "bdc", "bca", "bda", "bdca", "bdcal"}
	//t := "bdc"

	for _, v := range x {
		pcnt = 0
		RecursiveChain(v, 0, x)
		st = append(st, pcnt)
	}
	fmt.Println(st)
}
