/*
135. Candy
Hard
Topics
premium lock icon
Companies
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.



Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.


Constraints:

n == ratings.length
1 <= n <= 2 * 104
0 <= ratings[i] <= 2 * 104
*/

package main

import "fmt"

func candy(ratings []int) int {
	baseArray := getBaseArray(ratings)

	return calculateCandies(ratings, baseArray)
}

func getBaseArray(ratings []int) []int {
	newArray := make([]int, len(ratings))
	for i := range newArray {
		newArray[i] = 1
	}
	return newArray
}

func calculateCandies(ratings []int, array []int) int {
	for i := 0; i < len(ratings); i++ {
		if i-1 > 0 && ratings[i] > ratings[i-1] {
			array[i] += 1
		}
		if i+1 < len(ratings) && ratings[i] > ratings[i+1] {
			array[i] += 1
		}
	}

	return sumArray(array)
}

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println(candy(ratings))
}
