package main

import (
	__array "algo/3-array"
	"errors"
	"fmt"
)

func main() {
	__array.LuckyNumbers([][]int{
		{1, 2, 3},
		{4, 5, 3},
		{6, 7, 9},
	})

}

func foo() error {
	err := errors.New("dfdf")
	return fmt.Errorf("bar failed: %w", err)
}
