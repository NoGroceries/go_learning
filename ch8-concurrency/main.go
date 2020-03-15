package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var flag bool = false
	v :=l1.Val + l2.Val
	head :=&ListNode {
	}
	if v <10 {
		head.Val = v
	}else {
		head.Val = v%10
		flag = true
	}
	h := head
	l1 = l1.Next
	l2 = l2.Next
	for  {
		var v int
		if l1==nil && l2==nil {
			if flag {
				h.Next=&ListNode {
					Val:1,
					Next:nil,
				}
			}
			break
		}
		if l1==nil {
			v = l2.Val
		}
		if l2 ==nil {
			v=l1.Val
		}
		if l1!=nil && l2!=nil {
			v =l1.Val + l2.Val
		}

		if flag {
			v++
			flag = false
		}

		if v <10 {
			h.Next=&ListNode{
				Val:v,
				Next:nil,
			}

		}else {
			h.Next=&ListNode{
				Val:v%10,
				Next:nil,
			}

			flag = true
		}

		h = h.Next
		if l1!=nil {
			l1 = l1.Next
		}
		if l2!=nil {
			l2 = l2.Next
		}


	}
	return head

}

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 8,
			Next: nil,
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: nil,
	}

	addTwoNumbers(l1,l2)
}
