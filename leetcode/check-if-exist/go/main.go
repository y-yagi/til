package main

func checkIfExist(arr []int) bool {
	dict := map[int]bool{}

	for i := 0; i < len(arr); i++ {
		if _, found := dict[arr[i]*2]; found {
			return true
		}

		if _, found := dict[arr[i]/2]; found && arr[i]%2 == 0 {
			return true
		}

		dict[arr[i]] = true
	}

	return false
}
