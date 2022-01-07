package main

func main() {
	DoOperation(2, increase)
	DoOperation(2, decrease)
}

func DoOperation(y int, f func(int, int) int) {
	println(f(y, 1))
}

func increase(a, b int) int {
	return a + b
}

func decrease(a, b int) int {
	return a - b
}
