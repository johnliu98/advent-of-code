package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
)

func ObjID(s string) int {
	x := strings.Split(s, ":")
	x = strings.Split(x[0], " ")
	i, err := strconv.Atoi(x[1])
	if err != nil {
		panic(err)
	}
	return i
}

func Values(s string) string {
	x := strings.Split(s, ":")
	return x[1]
}

func SplitValues(s, sep string) []string {
	x := strings.Split(s, ":")
	return strings.Split(x[1], sep)
}

func Numbers(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	byteNums := re.FindAll([]byte(s), -1)

	var nums []int
	for _, n := range byteNums {
		nums = append(nums, conv.IntFromString(string(n)))
	}

	return nums
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
