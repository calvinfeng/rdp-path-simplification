package rdp

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	points := []Point{
		Point{0, 0},
		Point{1, 2},
		Point{2, 7},
		Point{3, 1},
		Point{4, 8},
		Point{5, 2},
		Point{6, 8},
		Point{7, 3},
		Point{8, 3},
		Point{9, 0},
	}

	t.Run("Threshold=0", func(t *testing.T) {
		if len(SimplifyPath(points, 0)) != 10 {
			t.Error("simplified path should have all points")
		}
	})

	t.Run("Threshold=2", func(t *testing.T) {
		if len(SimplifyPath(points, 2)) != 7 {
			t.Error("simplified path should only have 7 points")
		}
	})

	t.Run("Threshold=5", func(t *testing.T) {
		if len(SimplifyPath(points, 100)) != 2 {
			t.Error("simplified path should only have two points")
		}
	})
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
