/*

35. Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/

package main

import (
	"fmt"

	"github.com/SeiyaJapon/search_insert_position/search_insert"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	aux := search_insert.SearchInsert(nums, target, 0, len(nums)-1)

	fmt.Println("Result: ", aux)

	fmt.Println("----- Casos de prueba adicionales -----")

	// Construye un slice grande y ordenado con gaps variables
	n := 200000
	nums = make([]int, n)
	val := -1_000_000
	for i := 0; i < n; i++ {
		gap := (i % 13) + 1 // gaps entre 1 y 13
		if i%10000 == 0 && i != 0 {
			gap += 100 // saltos ocasionales grandes
		}
		val += gap
		nums[i] = val
	}

	// Targets que cubren casos extremos y puntos intermedios
	targets := []int{
		nums[0] - 500,    // mucho menor que el mínimo
		nums[0],          // primer elemento (existente)
		nums[1] - 1,      // entre primer y segundo
		nums[n/4],        // elemento en el primer cuarto (existente)
		nums[n/4] + 7,    // cercano pero posiblemente no existente
		nums[n/2] - 2,    // cercano al medio
		nums[n-1],        // último elemento (existente)
		nums[n-1] + 1000, // mucho mayor que el máximo
		123456789,        // valor arbitrario grande
		4,
		345,
		124,
		879,
	}

	for _, t := range targets {
		idx := search_insert.SearchInsert(nums, t, 0, len(nums)-1)
		fmt.Printf("target=%d -> index=%d\n", t, idx)
	}
}
