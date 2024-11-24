package main

type Color int

const (
	Red Color = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	violet
)

func main() {
	println(Red, Orange, Yellow, Green, Blue, Indigo, violet)
}
