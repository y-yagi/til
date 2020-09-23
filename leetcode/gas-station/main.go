package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	var totalTank, currTank, startingStation int
	for i := 0; i < n; i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]
		if currTank < 0 {
			startingStation = i + 1
			currTank = 0
		}
	}

	if totalTank >= 0 {
		return startingStation
	}
	return -1
}
