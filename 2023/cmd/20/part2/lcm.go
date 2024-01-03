package main

func leastCommonMultiple(ints ...int) int {
	if len(ints) == 0 {
		return 0
	}

	if len(ints) == 1 {
		return ints[0]
	}

	result := ints[0] * ints[1] / greatestCommonDivisor(ints[0], ints[1])

	for i := 2; i < len(ints); i++ {
		result = leastCommonMultiple(result, ints[i])
	}

	return result
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
