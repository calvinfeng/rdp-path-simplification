# Ramer-Douglas-Peucker Algorithm
Given a curve, composed of line segments, find a similar curve with fewer points.

## Results
![result_1](./examples/path.png)

## Usage
Go to `examples/` and run `dep ensure`. Then run `go install` if you don't mind the binary name is 
`examples`. If you do, use `go build`.

    go build -o rdp && ./rdp

Use `SimplifyPath` to run RDP path simplfication algorithm.
```golang
points = make([]rdp.Point{}, 0, 100)
for range 100 {
    // Insert your points
}

threshold := 0.5
results := rdp.SimplifyPath(points, threshold)
```
