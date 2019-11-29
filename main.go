package main

import (
	"fmt"
	"strconv"
)

// Write a function that returns the sum of two numbers.

// Example

// For param1 = 1 and param2 = 2, the output should be add(param1, param2) = 3.
func add(n1 int, n2 int) string {
	return strconv.Itoa(n1 + n2)
}

// Given the string, check if it is a palindrome.

// Example

// For inputString = "aabaa", the output should be checkPalindrome(inputString) = true; For inputString = "abac", the output should be checkPalindrome(inputString) = false; For inputString = "a", the output should be checkPalindrome(inputString) = true.
func checkPalindrome(s string) bool {
	unicodeArray := []rune(s)
	reverseArray := []rune(s)

	for i, j := len(unicodeArray)-1, 0; i >= 0; i, j = i-1, j+1 {
		reverseArray[j] = unicodeArray[i]
	}
	print(string(reverseArray), string(unicodeArray))
	return string(reverseArray) == string(unicodeArray)
}

// Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.

// Example

// For year = 1905, the output should be centuryFromYear(year) = 20; For year = 1700, the output should be centuryFromYear(year) = 17.

// Given the string, check if it is a palindrome.

// Example

// For inputString = "aabaa", the output should be checkPalindrome(inputString) = true; For inputString = "abac", the output should be checkPalindrome(inputString) = false; For inputString = "a", the output should be checkPalindrome(inputString) = true.
func centuryFromYear(year int) int {
	return (year / 100) + 1
}

// Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

// Example

// For inputArray = [3, 6, -2, -5, 7, 3], the output should be adjacentElementsProduct(inputArray) = 21.

// 7 and 3 produce the largest product.
// a := [...]int{3, 6, -2, -5, 7, 3}
// adjacentElementsProduct(a[:])

func adjacentElementsProduct(slice []int) int {
	var product int = -100000000000
	for i := 0; i < len(slice)-1; i++ {
		if slice[i]*slice[i+1] > product {
			product = slice[i] * slice[i+1]
		}
	}
	return product
}

// Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

// Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.

// Example

// For sequence = [1, 3, 2, 1], the output should be almostIncreasingSequence(sequence) = false.

// There is no one element in this array that can be removed in order to get a strictly increasing sequence.

// For sequence = [1, 3, 2], the output should be almostIncreasingSequence(sequence) = true.

// You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].

func almostIncreasingSequence(slice []int) bool {
	auxArr1 := make([]int, 0)
	auxArr2 := make([]int, 0)
	aux := make([]int, 0)

	if len(slice) == 1 {
		return true
	}
	for i := 0; i < len(slice)-1; i++ {
		if i == 0 {

		}
		if i == (len(slice) - 2) {

		}
		if slice[i] < slice[i+1] {
			auxArr1 = slice[:i]
			auxArr2 = slice[i:]
			aux = append(aux, auxArr1...)
			aux = append(aux, auxArr2...)
			fmt.Println(aux, auxArr1, auxArr2)
			for j := 0; j < len(aux)-1; j++ {
				if aux[j] < aux[j+1] {
					break
				}
				return true
			}
		}
	}
	return true
}

func main() {
	a := [...]int{1, 3, 2, 0}
	print(almostIncreasingSequence(a[:]))
}
