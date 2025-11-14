package main

import (
	"fmt"
	"slices"
)

func findSmallInteger(nums []int, value int) int {
	existingNumbers := getExistingNumbers(nums)
	numbersToFind := getNumbersToFind(nums, existingNumbers)

	modifiedNums := calculateNumberToFind(nums, numbersToFind, existingNumbers, value)

	minimums := getNumbersToFind(modifiedNums, modifiedNums)

	minimum := slices.Min(minimums)

	fmt.Println("minimum:", minimum)
	return -1
}

func getExistingNumbers(nums []int) []int {
	minimum := 0
	existingNumbers := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == minimum {
				existingNumbers = append(existingNumbers, minimum)
				break
			}
		}

		minimum++
	}

	return existingNumbers
}

func getNumbersToFind(numbers []int, values []int) []int {
	numbersToFind := []int{}

	for i := 0; i < len(numbers); i++ {
		if !contains(values, i) {
			numbersToFind = append(numbersToFind, i)
		}
	}

	return numbersToFind
}

func contains(nums []int, value int) bool {
	for _, num := range nums {
		if num == value {
			return true
		}
	}
	return false
}

func calculateNumberToFind(nums []int, toFind []int, finded []int, referncedValue int) []int {
	if referncedValue < 0 {
		referncedValue = -1 * referncedValue
	}
	var auxValue int
	modifiedNums := nums

	for index, num := range nums {
		if !contains(toFind, num) && !contains(finded, num) {
			auxValue = num
			for i := 0; i < len(nums); i++ {
				if num < 0 {
					auxValue += referncedValue
				} else {
					auxValue -= referncedValue
				}

				if contains(toFind, auxValue) && !contains(nums, auxValue) {
					modifiedNums[index] = auxValue
					break
				}
			}

		}
	}

	return modifiedNums
}

func main() {
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 5

	findSmallInteger(nums, value)
}
