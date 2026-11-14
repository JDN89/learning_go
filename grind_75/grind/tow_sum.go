package grind

// You are given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, value := range nums {
		tocheck := target - value

		key, ok := m[tocheck]
		if ok {
			return []int{key, idx}
		} else {
			m[value] = idx
		}
	}
	return []int{0, 0}
}
