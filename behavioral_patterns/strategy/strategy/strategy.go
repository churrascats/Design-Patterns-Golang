package strategy

import (
	"fmt"
	"slices"
)

type Strategy interface {
	ShowArrayItems(array []int)
}

type SortedArrayStrategy struct{}

func (s *SortedArrayStrategy) ShowArrayItems(array []int) {
	copiedArray := make([]int, len(array))
	copy(copiedArray, array)

	slices.Sort(copiedArray)

	fmt.Println(copiedArray)
}

type SortedReverseArrayStrategy struct{}

func (s *SortedReverseArrayStrategy) ShowArrayItems(array []int) {

	copiedArray := make([]int, len(array))
	copy(copiedArray, array)

	slices.Sort(copiedArray)
	slices.Reverse(copiedArray)

	fmt.Println(copiedArray)
}

type SortedDoubledArrayStrategy struct{}

func (s *SortedDoubledArrayStrategy) ShowArrayItems(array []int) {
	copiedArray := make([]int, len(array))
	copy(copiedArray, array)

	slices.Sort(copiedArray)
	for index, item := range copiedArray {
		copiedArray[index] = item * 2
	}

	fmt.Println(copiedArray)
}
