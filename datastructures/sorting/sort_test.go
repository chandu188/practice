package sorting

import "testing"
import "github.com/stretchr/testify/assert"

func TestSelectionSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	SelectionSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestInsertionSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	InsertionSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestShellSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	ShellSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestMergeSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	MergeSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestBottomUpMergeSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	bottomUpMergeSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestQuickSort(t *testing.T) {
	in := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	QuickSort(in)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func Test3WaySort(t *testing.T) {
	in := []int{5, 2, 5, 3, 5, 7, 9, 13, 5}
	_3WaySort(in)

	res := []int{2, 3, 5, 5, 5, 5, 7, 9, 13}
	for i, v := range res {
		assert.Equal(t, v, in[i])
	}
}
