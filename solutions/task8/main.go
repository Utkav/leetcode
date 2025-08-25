package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    resultList := &ListNode{}
    current := resultList
    
    for list1 != nil && list2 != nil{
        if list1.Val <= list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    
    if list1 == nil{
        current.Next = list2
    } else{
        current.Next = list1
    }
    
    return resultList.Next
}


func main() {
	node23 := &ListNode{Val: 4}
	node22 := &ListNode{Val: 2, Next: node23}
	head2 := &ListNode{Val: 1, Next: node22}

	node13 := &ListNode{Val: 4}
	node12 := &ListNode{Val: 3, Next: node13}
	head1 := &ListNode{Val: 1, Next: node12}

	head := mergeTwoLists(head1, head2)

	for head != nil{
		fmt.Println(head.Val)
        head = head.Next
    }
}