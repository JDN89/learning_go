package grind_test

import (
	"slices"
	"testing"

	"grind_75/grind"
)

func TestTwoSum(t *testing.T) {
	got := grind.TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !slices.Equal(got, want) {
		t.Errorf(" TwoSum result got %d != want %d", got, want)
	}
}

func TestIsValid(t *testing.T) {
	input := "()[]{}"
	got := grind.IsValid(input)

	want := true
	// NOTE print boolean https://pkg.go.dev/fmt
	if got != want {
		t.Errorf("wanted %t , but got %t ", got, want)
	}
}

func listFromSlice(sl []int) *grind.ListNode {
	emptyHeadNode := &grind.ListNode{}
	tail := emptyHeadNode

	for _, val := range sl {
		// append node after tail, then move tail to the new last node.

		// dummy -> nil

		// dummy -> [1] -> nil
		// dummy -> [1] -> [2] -> nil

		// ListNode without Next assigned. zero value of a pointer is nil!!! here implicit, maybe better to set visually to nil?
		// tail.Next = &grind.ListNode{Val: val}
		tail.Next = &grind.ListNode{Val: val, Next: nil}
		tail = tail.Next
	}
	return emptyHeadNode.Next
}

func TestMergeTwoLists(t *testing.T) {
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	got := grind.MergeTwoLists(listFromSlice(list1), listFromSlice(list2))

	want := listFromSlice([]int{1, 1, 2, 3, 4, 4})

	gotPointer := got
	wantPointer := want

	for i := 0; gotPointer != nil && wantPointer != nil; i++ {
		if gotPointer.Val != wantPointer.Val {
			t.Errorf("node %d: want %d, got %d", i, wantPointer.Val, gotPointer.Val)
		}

		gotPointer = gotPointer.Next
		wantPointer = wantPointer.Next
	}

	if gotPointer != nil || wantPointer != nil {
		t.Errorf("lists have different lengths")
	}
}
