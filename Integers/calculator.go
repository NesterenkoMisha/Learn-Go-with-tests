package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

// Subtracts y from x and returns the result.
func Sub(x, y int) int {
	return x - y
}

// Multiplies x by y and returns the result.
func Mul(x, y int) int {
	return x * y
}

// Divides x by y and returns the result as float.
func Div(x, y int) float64 {
	if y == 0 {
		panic("Sadly, it's impossible to divide by 0")
	}
	return float64(x / y)
}
