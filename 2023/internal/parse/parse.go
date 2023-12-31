package parse

import (
	"regexp"
	"strconv"
)

func Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Ints(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	bytes := re.FindAll([]byte(s), -1)
	ints := make([]int, len(bytes))
	for i, b := range bytes {
		ints[i] = Int(string(b))
	}
	return ints
}

func Blocks(ss []string) [][]string {
	var blocks [][]string

	var start int
	for i := 0; i < len(ss); i++ {
		if ss[i] != "" {
			continue
		}

		if len(ss[start:i]) != 0 {
			blocks = append(blocks, ss[start:i])
		}

		start = i + 1
	}
	if len(ss[start:]) != 0 {
		blocks = append(blocks, ss[start:])
	}

	return blocks
}
