package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{6, 2, 3, 3, 1, 9, 19}
	n := BubbleSort(nums, false)

	nums = []int{6, 2, 3, 3, 1, 9, 19}
	r := BubbleSort(nums, true)

	fmt.Println("new: ", n)
	fmt.Println("reversed: ", r)

	if len(n) != len(r) {
		t.Error("not equal")
	}

	if n[0] != r[len(nums)-1] {
		t.Error("sort error")
	}
}
