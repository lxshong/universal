package int

// gets the sum of multiple int values
func Sum(params ...int) int {
	var sum int
	for _, param := range params {
		sum += param
	}
	return sum
}
