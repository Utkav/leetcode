package main

import "fmt"

 type ListNode struct {
    Val int
    Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := head
    size := 0
    for node != nil{
        node = node.Next
        size += 1
    }
    
    node = head
    
    if size == n{
        return head.Next
    }
    if size == 1{
        return nil
    }
    for i := 0; i < size - n; i++{
        if i == size - n - 1{
            node.Next = node.Next.Next
        }
        node = node.Next
    }

    return head
}

func main(){
	node5 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	removeNthFromEnd(head, 2)

	for head != nil{
		fmt.Println(head.Val)
        head = head.Next
    }
}
