package main

// SimplifyPath

// SimplifyPath accepts a list of points and epsilon as tolerance, simplifies a path by dropping
// some portion of the points.
func SimplifyPath(points []Point, ep float64) []Point {
	if len(points) <= 2 {
		return []Point{points[0]}
	}

	l := Line{Start: points[0], End: points[len(points)-1]}

	idx, maxDist := seekMostDistantPoint(l, points)
	if maxDist >= ep {
		left := SimplifyPath(points[:idx+1], ep)
		right := SimplifyPath(points[idx:], ep)
		return append(left[:len(left)-1], right...)
	}

	// If the most distant point fails to pass the tolerance test, then just return the two points
	return []Point{points[0], points[len(points)-1]}
}

func seekMostDistantPoint(l Line, points []Point) (idx int, maxDist float64) {
	for i := 0; i < len(points); i++ {
		d := l.DistanceToPoint(points[i])
		if d > maxDist {
			maxDist = d
			idx = i
		}
	}

	return idx, maxDist
}
