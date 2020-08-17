package main

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	i := 0
	candy := 1

	for candies > 0 {
		candies -= candy
		ans[i] += candy

		i++
		if i >= num_people {
			i = 0
		}

		candy++
		if candies-candy < 0 {
			candy = candies
		}
	}

	return ans
}
