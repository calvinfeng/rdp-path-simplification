package rdp

import "testing"

func TestLine(t *testing.T) {
	t.Run("TestCoefficients", func(t *testing.T) {
		lines := []Line{
			Line{Start: Point{0, 0}, End: Point{2, 2}},
			Line{Start: Point{-1, 0}, End: Point{3, 5}},
			Line{Start: Point{-3, -5}, End: Point{3, 7}},
		}

		for _, l := range lines {
			a, b, c := l.Coefficients()

			// Expected slope
			expected := (l.End.Y - l.Start.Y) / (l.End.X - l.Start.X)

			slope := -1 * a / b
			if slope != expected {
				t.Error("Wrong slope", slope)
			}

			// Expected y intercept
			expected = l.End.Y - (expected * l.End.X)
			yIntercept := -1 * c / b
			if yIntercept != expected {
				t.Error("Wrong y-intercept", yIntercept)
			}
		}
	})

	t.Run("TestDistanceToPoint", func(t *testing.T) {
		l := Line{Start: Point{0, 0}, End: Point{10, 0}}

		points := []Point{
			Point{5, 5},
			Point{3, 6},
			Point{1, -7},
			Point{-4, -10},
		}

		expectations := []float64{5, 6, 7, 10}

		for i, p := range points {
			if expectations[i] != l.DistanceToPoint(p) {
				t.Errorf("%f is not equal to %f", expectations[i], l.DistanceToPoint(p))
			}
		}
	})
}
