package main

import "testing"

func TestSimplifyPath(t *testing.T) {

}

func TestSeekMostDistantPoint(t *testing.T) {
	l := Line{Start: Point{0, 0}, End: Point{10, 0}}
	points := []Point{
		Point{13, 13},
		Point{1, 15},
		Point{1, 1},
		Point{3, 6},
	}

	idx, maxDist := seekMostDistantPoint(l, points)

	if idx != 1 {
		t.Error("failed to find most distant point away from a line")
	}

	if maxDist != 15 {
		t.Error("maximum distance is incorrect")
	}
}
