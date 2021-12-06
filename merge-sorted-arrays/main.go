package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	for {
		if l2.Next == nil {
			break
		}
	}
	return l1
}

func main() {
	val11 := &ListNode{4, nil}
	val12 := &ListNode{2, val11}
	val13 := &ListNode{1, val12}
	val21 := &ListNode{4, nil}
	val22 := &ListNode{3, val21}
	val23 := &ListNode{1, val22}
	returnVal := mergeTwoLists(val13, val23)
	for {
		fmt.Println(returnVal.Val)
		returnVal = returnVal.Next
		if returnVal == nil {
			break
		}
	}
}
