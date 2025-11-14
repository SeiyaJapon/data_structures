/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1: [The numbers I indicate between double brackets are nodes of a linked list]
((2)) -> ((4)) -> ((3))
((5)) -> ((6)) -> ((4))
----------------------------
((7)) -> ((0)) -> ((8))

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
package main

import (
	"fmt"

	"github.com/SeiyaJapon/linked_list_add_two_numbers/linked_list"
)

func main() {
	list := linked_list.NewLinkedList()
	list.Append(2)
	list.Append(4)
	list.Append(3)

	anotherList := linked_list.NewLinkedList()
	anotherList.Append(5)
	anotherList.Append(6)
	anotherList.Append(4)

	num1 := myPow(list)
	fmt.Println(num1)

	num2 := myPow(anotherList)
	fmt.Println(num2)

	sum := num1 + num2
	fmt.Println(sum)

	newList := sumToList(sum)

	newList.Iterate(func(v interface{}) {
		fmt.Println(v)
	})
}

func myPow(list *linked_list.LinkedList) int {
	num := 0
	count := 0
	for node := list.Head; node != nil; node = node.Next {
		pow := 1
		for i := 0; i < count; i++ {
			pow *= 10
		}
		num += node.Val.(int) * pow
		count++
	}
	return num
}

func sumToList(sum int) *linked_list.LinkedList {
	l := linked_list.NewLinkedList()
	if sum == 0 {
		l.Append(0)
		return l
	}
	for sum > 0 {
		l.Append(sum % 10)
		sum /= 10
	}
	return l
}
