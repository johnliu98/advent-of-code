package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const numBoxes = 256

func main() {
	initializationSequence := read.ReadLines("input.txt")[0]
	steps := strings.Split(initializationSequence, ",")

	var boxes [numBoxes][]item
	re := regexp.MustCompile(`\w+`)
	for _, step := range steps {
		lens := string(re.Find([]byte(step)))
		h := hash(lens)
		itemIndex := getIndex(boxes[h], lens)
		if n := step[len(step)-1]; unicode.IsNumber(rune(n)) {
			focalLength := conv.IntFromString(string(n))
			if itemIndex != -1 {
				boxes[h][itemIndex].focalLength = focalLength
			} else {
				boxes[h] = append(boxes[h], item{
					lens:        lens,
					focalLength: focalLength,
				})
			}
		} else if itemIndex != -1 {
			boxes[h] = append(boxes[h][:itemIndex], boxes[h][itemIndex+1:]...)
		}
	}

	var ans int
	for i, box := range boxes {
		for j, lens := range box {
			ans += (i + 1) * (j + 1) * lens.focalLength
		}
	}

	fmt.Println("Answer: ", ans)
}

func hash(s string) int {
	var h uint8
	for _, c := range s {
		h += uint8(c)
		h *= 17
	}
	return int(h)
}

type item struct {
	lens        string
	focalLength int
}

func getIndex(box []item, lens string) int {
	for bi, it := range box {
		if lens == it.lens {
			return bi
		}
	}
	return -1
}
