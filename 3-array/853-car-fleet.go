package __array

import (
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {
	var stack []float32

	var carList [][2]int

	for i := 0; i < len(position); i++ {
		carList = append(carList, [2]int{position[i], speed[i]})
	}

	slices.SortFunc(carList, func(a, b [2]int) int {
		return b[0] - a[0]
	})

	for _, car := range carList {
		carPosition := car[0]
		carSpeed := car[1]
		time := (float32(target) - float32(carPosition)) / float32(carSpeed)
		stack = append(stack, time)

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
