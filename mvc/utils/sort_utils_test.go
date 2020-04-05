package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortUtilsWorstCase(t *testing.T) {
	elem := []int{7, 6, 5, 4, 3}
	BubbleSort(elem)

	assert.NotNil(t, elem)
	assert.EqualValues(t, 5, len(elem))
	assert.EqualValues(t, []int{3, 4, 5, 6, 7}, elem)
}

func TestSortUtilBestCase(t *testing.T) {
	elem := []int{5, 6, 7, 8, 9}
	BubbleSort(elem)

	assert.NotNil(t, elem)
	assert.EqualValues(t, 5, len(elem))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, elem)
}

func TestSortUtilNil(t *testing.T) {
	BubbleSort(nil)
}

func getElements(n int) []int {
	result := make([]int, n)
	for j, i := n-1, 0; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort10(b *testing.B) {
	elems := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(elems)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elems := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(elems)
	}
}

func BenchmarkSort50000(b *testing.B) {
	elems := getElements(50000)
	for i := 0; i < b.N; i++ {
		Sort(elems)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elems := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elems)
	}
}
