package grind

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked li

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Example 2:
//
// Input: list1 = [], list2 = []
// Output: []
//
// Example 3:
//
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// for list1.Next != nil {
	// 	list1.Val = list1.Next.Val
	// 	list1.Next = list1.Next.Next
	// }

	dummy := &ListNode{}
	tail := dummy
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next

	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	// Dummy is empty node HeadPOinter so we return Head.Next
	return dummy.Next

	// NOTE: MESSED UP CODE JAN
	// merged := &ListNode{}
	// mergedTail := merged
	//
	// list1Pointer := list1
	// list2Pointer := list2

	// for list1Pointer.Next != nil {
	// 	if list2Pointer == nil {
	//
	// 		mergedTail.Val = list1Pointer.Val
	// 		// we point now to list2 next node
	// 		mergedTail.Next = list1Pointer.Next
	// 		mergedTail = mergedTail.Next
	// 		list1Pointer = list1Pointer.Next
	// 		continue
	// 	}
	// 	if list1Pointer == nil {
	//
	// 		mergedTail.Val = list2Pointer.Val
	// 		// we point now to list2 next node
	// 		mergedTail.Next = list2Pointer.Next
	// 		mergedTail = mergedTail.Next
	// 		list2Pointer = list2Pointer.Next
	// 		continue
	// 	}
	// 	if list1Pointer.Val > list2Pointer.Val {
	// 		mergedTail.Val = list2Pointer.Val
	// 		// we point now to list2 next node
	// 		mergedTail.Next = list2Pointer.Next
	// 		mergedTail = mergedTail.Next
	// 		list2Pointer = list2Pointer.Next
	// 		continue
	// 	}
	// 	if list2Pointer.Val > list1Pointer.Val {
	// 		mergedTail.Val = list1Pointer.Val
	// 		// we point now to list2 next node
	// 		mergedTail.Next = list1Pointer.Next
	// 		mergedTail = mergedTail.Next
	// 		list1Pointer = list1Pointer.Next
	// 		continue
	// 	}
	// 	if list2Pointer.Val == list1Pointer.Val {
	// 		mergedTail.Val = list1Pointer.Val
	// 		// we point now to list2 next node
	// 		mergedTail.Next = list1Pointer.Next
	// 		mergedTail = mergedTail.Next
	// 		list1Pointer = list1Pointer.Next
	// 		continue
	// 	}
	// }

	// merged
}
