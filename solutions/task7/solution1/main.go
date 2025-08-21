package main

import (
	"fmt"
)

type ListNode struct {
 	Val int
	 Next *ListNode
 }

func reverseList(head *ListNode) *ListNode {
    
    var listValues []int
    for node := head; node != nil; node = node.Next{
        value := node.Val
        listValues = append(listValues, value)
    }
    
    index := len(listValues) - 1
    for node := head; node != nil; node = node.Next{
        node.Val = listValues[index]
        index--
    }
    return head
}

func main(){
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	reverseList(head)

	for head != nil{
		fmt.Println(head.Val)
        head = head.Next
    }
}