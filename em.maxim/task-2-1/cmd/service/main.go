package main

import (
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
	signGreaterEq  = ">="
	signLessEq     = "<="
	noSolutionCode = -1
)

type TempRange struct {
	low  int
	high int
}

func initRange() TempRange {
	return TempRange{
		low:  defaultMinTemp,
		high: defaultMaxTemp,
	}
}

func updateRange(current TempRange, op string, value int) TempRange {
	switch op {
	case signGreaterEq:
		if value > current.low {
			current.low = value
		}
	case signLessEq:
		if value < current.high {
			current.high = value
		}
	}

	return current
}

func readInt(prompt string) (int, bool) {
	var val int
	if _, err := fmt.Scan(&val); err != nil {
		fmt.Println(prompt)

		return 0, false
	}

	return val, true
}

func main() {
	countDepts, ok := readInt("Invalid departments count")
	if !ok {
		return
	}

	for range countDepts {
		countStaff, ok := readInt("Invalid employees count")
		if !ok {
			return
		}

		rng := initRange()

		for range countStaff {
			var operation string
			var temperature int

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				fmt.Println("Invalid temperature constraint")

				return
			}

			rng = updateRange(rng, operation, temperature)

			if rng.low > rng.high {
				fmt.Println(noSolutionCode)

				return
			}

			fmt.Println(rng.low)
		}
	}
}
