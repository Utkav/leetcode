package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return true
    }
    
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    current := slow
    var prev *ListNode
    prev = nil
    
    for current != nil{
        nextNode := current.Next
        current.Next = prev
        prev = current
        current = nextNode
    }
    
    for prev != nil{
        if head.Val != prev.Val{
            return false
        } 
        head = head.Next
        prev = prev.Next
    }
    return true
}

func main(){
	node5 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 2, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}
	fmt.Println(isPalindrome(head))
}