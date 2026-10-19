package solution

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := make([]int, 0)

	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

	}

	return result
}
