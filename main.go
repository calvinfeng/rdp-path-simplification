package main

func main() {
	if err := SavePlot(RandomPoints(10)); err != nil {
		panic(err)
	}
}
