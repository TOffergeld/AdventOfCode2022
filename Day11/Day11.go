package Day11

import (
	"fmt"
	"math"
)

type Monkey struct {
	items     []int
	op        func(int) int
	test      func(int) bool
	tar       map[bool]int
	inspected int
}

func postProcess(val, div int) int {
	r := val / div
	return int(math.Round(float64(r) - .5)) // "round down to nearest integer"
}

func Day11() {
	monkeys := []Monkey{
		{
			[]int{64},
			func(old int) int { return old * 7 },
			func(val int) bool { return (val % 13) == 0 },
			map[bool]int{true: 1, false: 3},
			0,
		},
		{
			[]int{60, 84, 84, 65},
			func(old int) int { return old + 7 },
			func(val int) bool { return (val % 19) == 0 },
			map[bool]int{true: 2, false: 7},
			0,
		},
		{
			[]int{52, 67, 74, 88, 51, 61},
			func(old int) int { return old * 3 },
			func(val int) bool { return (val % 5) == 0 },
			map[bool]int{true: 5, false: 7},
			0,
		},
		{
			[]int{67, 72},
			func(old int) int { return old + 3 },
			func(val int) bool { return (val % 2) == 0 },
			map[bool]int{true: 1, false: 2},
			0,
		},
		{
			[]int{80, 79, 58, 77, 68, 74, 98, 64},
			func(old int) int { return old * old },
			func(val int) bool { return (val % 17) == 0 },
			map[bool]int{true: 6, false: 0},
			0,
		},
		{
			[]int{62, 53, 61, 89, 86},
			func(old int) int { return old + 8 },
			func(val int) bool { return (val % 11) == 0 },
			map[bool]int{true: 4, false: 6},
			0,
		},
		{
			[]int{86, 89, 82},
			func(old int) int { return old + 2 },
			func(val int) bool { return (val % 7) == 0 },
			map[bool]int{true: 3, false: 0},
			0,
		},
		{
			[]int{92, 81, 70, 96, 69, 84, 83},
			func(old int) int { return old + 4 },
			func(val int) bool { return (val % 3) == 0 },
			map[bool]int{true: 4, false: 5},
			0,
		},
	}
	var subjects []Monkey
	subjects = monkeys
	prodMod := 13 * 19 * 5 * 2 * 17 * 11 * 7 * 3 //hardcoded divisor to keep modulo
	for iteration := 0; iteration < 10000; iteration++ {
		for m_idx, m := range subjects {
			for _, item := range m.items {
				newValue := m.op(item)
				newValue = postProcess(newValue, 1)

				sub_idx := m.tar[m.test(newValue)]
				newValue = newValue % prodMod
				subjects[sub_idx].items = append(subjects[sub_idx].items, newValue)
				//fmt.Println("Throwing item with value", newValue, "to Monkey", sub_idx)
				subjects[m_idx].inspected++
			}
			subjects[m_idx].items = []int{}
		}
	}
	for i, m := range subjects {
		fmt.Println("Activity of monkey", i, "is", m.inspected)
	}
}