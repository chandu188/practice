package sorting

import "testing"
import "github.com/stretchr/testify/assert"
func TestSelectionSort(t *testing.T) {
	in := []int {4,5,2,1,3,9,7,8,6}
	SelectionSort(in)
	for i:=1;i<10;i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestInsertionSort(t *testing.T) {
	in := []int {4,5,2,1,3,9,7,8,6}
	InsertionSort(in)
	for i:=1;i<10;i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestShellSort(t *testing.T) {
	in := []int {4,5,2,1,3,9,7,8,6}
	ShellSort(in)
	for i:=1;i<10;i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestMergeSort(t *testing.T) {
	in := []int {4,5,2,1,3,9,7,8,6}
	MergeSort(in)
	for i:=1;i<10;i++ {
		assert.Equal(t, i, in[i-1])
	}
}

func TestBottomUpMergeSort(t *testing.T) {
	in := []int {4,5,2,1,3,9,7,8,6}
	bottomUpMergeSort(in)
	for i:=1;i<10;i++ {
		assert.Equal(t, i, in[i-1])
	}
}