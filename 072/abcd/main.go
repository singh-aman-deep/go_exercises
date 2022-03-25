package abcd

func Sum(param ...int) int {
	total := 0
	for _, v := range param {
		total = total + v
	}
	return total
}

func Mul(param ...int) int {
	total := 1
	for _, v := range param {
		total = total * v
	}
	return total
}
