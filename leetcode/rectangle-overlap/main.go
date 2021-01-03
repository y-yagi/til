package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// check if either rectangle is a line
	if rec1[0] == rec1[2] || rec1[1] == rec1[3] ||
		rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		// the line can't have positive overlap
		return false
	}

	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) //top

	return false
}
