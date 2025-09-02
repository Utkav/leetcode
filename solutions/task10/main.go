package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
 }

func hasCycle(head *ListNode) bool {
	repNumber := make(map[*ListNode]bool)
	current := head

	for current != nil{
		_, ok := repNumber[current]

		if ok{
			return true
		} else{
			repNumber[current] = true
		}
		current = current.Next
	}
	return false
}

func main(){
	node5 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 1, Next: node5}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 3, Next: node2}
	node5.Next = node3
	fmt.Println(hasCycle(head))
}