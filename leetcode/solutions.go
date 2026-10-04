package leetcode

// TwoSum returns the indices of two numbers that add up to target.
// It returns nil when no such pair exists.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

// IsValidParentheses reports whether every opening bracket is closed
// by a matching bracket in the correct order.
func IsValidParentheses(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, bracket := range s {
		switch bracket {
		case '(', '[', '{':
			stack = append(stack, bracket)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (bracket == ')' && last != '(') ||
				(bracket == ']' && last != '[') ||
				(bracket == '}' && last != '{') {
				return false
			}
		default:
			return false
		}
	}
	return len(stack) == 0
}
