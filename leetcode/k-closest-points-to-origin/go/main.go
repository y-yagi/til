package main

import "sort"

type EuclideanPoints [][]int

func (pts EuclideanPoints) Len() int { return len(pts) }
func (pts EuclideanPoints) Less(i, j int) bool {
	return pts[i][0]*pts[i][0]+pts[i][1]*pts[i][1] < pts[j][0]*pts[j][0]+pts[j][1]*pts[j][1]
}
func (pts EuclideanPoints) Swap(i, j int) { pts[i], pts[j] = pts[j], pts[i] }

func kClosest(points [][]int, K int) [][]int {
	sort.Sort(EuclideanPoints(points))
	return points[:K]
}
