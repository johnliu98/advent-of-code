package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	g := game{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		g.hands = append(g.hands, hand(split[0]))
		g.bids = append(g.bids, conv.IntFromString(split[1]))
	}

	sort.Sort(g)

	var ans int
	for i, bid := range g.bids {
		rank := i + 1
		ans += rank * bid
	}

	fmt.Println("Answer: ", ans)
}

type game struct {
	hands []hand
	bids  []int
}

func (g game) Len() int {
	return len(g.hands)
}

func (g game) Less(i, j int) bool {
	return g.hands[i].value() < g.hands[j].value()
}

func (g game) Swap(i, j int) {
	g.hands[i], g.hands[j] = g.hands[j], g.hands[i]
	g.bids[i], g.bids[j] = g.bids[j], g.bids[i]
}

type hand string

const (
	numCardInHand = 5
	cardBits      = 4
)

func (h hand) value() int {
	return h.typeValue() + h.cardValue()
}

func (h hand) cardValue() int {
	var value int
	for i, r := range h {
		bitShift := (numCardInHand - i - 1) * cardBits
		cardValue := card(r).value()
		value += cardValue << bitShift
	}

	return value
}

const (
	pair         int = 1 << (numCardInHand*cardBits + 0)
	threeOfAKind int = 1 << (numCardInHand*cardBits + 2)
	fourOfAKind  int = 1 << (numCardInHand*cardBits + 3)
	fiveOfAKind  int = 1 << (numCardInHand*cardBits + 4)
)

func (h hand) typeValue() int {
	cardCounts := make(map[rune]int)
	for _, r := range h {
		cardCounts[r]++
	}

	var value int
	for _, counts := range cardCounts {
		switch counts {
		case 2:
			value += pair
		case 3:
			value += threeOfAKind
		case 4:
			value += fourOfAKind
		case 5:
			value += fiveOfAKind
		}
	}

	return value
}

type card string

const (
	two   card = "2"
	three card = "3"
	four  card = "4"
	five  card = "5"
	six   card = "6"
	seven card = "7"
	eight card = "8"
	nine  card = "9"
	ten   card = "T"
	jack  card = "J"
	queen card = "Q"
	king  card = "K"
	ace   card = "A"
)

func (c card) value() int {
	switch c {
	case two:
		return 2
	case three:
		return 3
	case four:
		return 4
	case five:
		return 5
	case six:
		return 6
	case seven:
		return 7
	case eight:
		return 8
	case nine:
		return 9
	case ten:
		return 10
	case jack:
		return 11
	case queen:
		return 12
	case king:
		return 13
	case ace:
		return 14
	default:
		panic(fmt.Sprintf("invalid card: %s", c))
	}
}
