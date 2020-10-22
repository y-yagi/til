package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, asteroid)
			continue
		}

		for len(stack) > 0 && (stack[len(stack)-1] > 0 && asteroid < 0) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			collision := top + asteroid
			if collision > 0 {
				stack = append(stack, top)
				asteroid = 0
				break
			} else if collision == 0 {
				asteroid = 0
				break
			}
		}

		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
