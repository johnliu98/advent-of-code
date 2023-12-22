package conv

import (
	"regexp"
	"strconv"
)

func IntFromString(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IntsFromString(s string) []int {
	re := regexp.MustCompile(`\d+`)
	return IntsFromByteSlices(re.FindAll([]byte(s), -1))
}

func IntsFromByteSlices(bs [][]byte) []int {
	ints := make([]int, len(bs))
	for i, b := range bs {
		ints[i] = IntFromString(string(b))
	}
	return ints
}
