package leetcode

func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	for day := 0; day < len(temperatures); day++ {
		for futureDay := day + 1; futureDay < len(temperatures); futureDay++ {
			if temperatures[futureDay] > temperatures[day] {
				answer[day] = futureDay - day
				break
			}
		}
	}
	return answer
}

func DailyTemperaturesStack(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := []int{}
	for day := 0; day < len(temperatures); day++ {
		curTemp := temperatures[day]
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < curTemp {
			prevDay := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prevDay] = day - prevDay
		}
		stack = append(stack, day)
	}
	return answer
}
