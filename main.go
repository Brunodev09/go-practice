package main

import (
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

//	a := [...]int{1, 3, 2, 3, 3}
// print(almostIncreasingSequence(a[:]))
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isSequence(slice []int, index int) bool {
	aux := remove(slice[:], index)
	for i := 1; i < len(aux)-1; i++ {
		if aux[i] >= aux[i+1] || aux[i] <= aux[i-1] {
			return false
		}
	}
	return true
}

func almostIncreasingSequence(slice []int) bool {
	sequence := true
	for i := 1; i < len(slice)-1; i++ {
		if slice[i] >= slice[i+1] {
			sequence = isSequence(slice[:], i)
			if !sequence {
				return sequence
			}
		}
		if slice[i] < slice[i-1] {
			sequence = isSequence(slice[:], i-1)
			if !sequence {
				return sequence
			}
		}
	}
	return sequence
}

// Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue having an non-negative integer size. Since he likes to make things perfect, he wants to arrange them from smallest to largest so that each statue will be bigger than the previous one exactly by 1. He may need some additional statues to be able to accomplish that. Help him figure out the minimum number of additional statues needed.

// Example

// For statues = [6, 2, 3, 8], the output should be makeArrayConsecutive2(statues) = 3.

// Ratiorg needs statues of sizes 4, 5 and 7.

func main() {

}
