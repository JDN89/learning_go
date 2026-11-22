package grind

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	result := -1

	for right >= left {
		// 0 + (0 - 6) / 2 = 3
		// 3 + (3 - 6) / 2 = 4 -> tel links erbij. Links is altijd start positie

		mid := left + (right-left)/2
		fmt.Println(mid)

		if nums[mid] == target {
			result = mid
			break
		}

		// left side of mid
		if nums[mid] > target {
			fmt.Println("mid > target")
			right = mid - 1
			continue
		}

		// right side of mid
		if nums[mid] < target {
			fmt.Println("mid < target")
			left = mid + 1
			continue
		}
	}

	return result
}
