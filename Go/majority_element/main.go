package main

func main() {
	//arr := []int{3, 2, 3}
	//arr := []int{1, 5, 1, 8, 1, 9, 1, 1}
	arr := []int{1, 1, 1, 5, 8, 9}
	result := majorityElement(arr)
	println(result)
}

func majorityElement(nums []int) int {
	majority := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			majority = num
			count = 1
		} else if num == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}
