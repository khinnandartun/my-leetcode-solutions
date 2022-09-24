package main

import "fmt"

/**
 * DEFINITION FOR SINGLY-LINKED LIST. */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) addNode(Value int) {
	ln.Next = &ListNode{Val: Value}
	ln = ln.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tempResult ListNode
	carry := 0
	var temp1, temp2 ListNode
	temp1 = *l1
	temp2 = *l2
	for temp1.Next != nil || temp2.Next != nil || carry != 0 {
		tempResult.Val = (carry + temp1.Val + temp2.Val) % 10
		carry = (temp1.Val + temp2.Val) / 10
		if temp1.Next == nil {
			temp1.Next = &ListNode{}
		}
		if temp2.Next == nil {
			temp2.Next = &ListNode{}
		}
		temp1 = *temp1.Next
		temp2 = *temp2.Next
		fmt.Println("result ", tempResult.Val, "carry ", carry)
		tempResult.Next = &ListNode{}
		tempResult = *tempResult.Next

	}
	return tempResult.Next

}
func main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8}}}
	result := addTwoNumbers(&l1, &l2)
	fmt.Println(result)

}
