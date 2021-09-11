package network

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap_Insert(t *testing.T) {
	h := NewMinHeap()

	inputs := []int{5, 12, 18, 7, 2, 15}
	for i, dist := range inputs {
		s := NewSwitch(strconv.Itoa(i))
		h.Insert(s, dist)
		verify(t, h, 0)
	}
}

func TestMinHeap_Extract(t *testing.T) {
	h := NewMinHeap()

	inputs := []int{18, 12, 16, 2, 9, 8, 26}
	for i, dist := range inputs {
		s := NewSwitch(strconv.Itoa(i))
		h.Insert(s, dist)
		verify(t, h, 0)
	}

	expectedOutputs := []int{2, 8, 9, 12, 16, 18, 26}
	for _, expected := range expectedOutputs {
		min, err := h.Extract()
		if err != nil {
			t.Errorf("error extracting min: %v", err)
		}
		verify(t, h, 0)
		assert.Equal(t, expected, min.Distance)
	}
}

func TestMinHeap_Random(t *testing.T) {
	h := NewMinHeap()

	for i := 0; i < 100; i++ {
		s := NewSwitch(strconv.Itoa(i))
		h.Insert(s, rand.Intn(100000))
		verify(t, h, 0)
	}

	for i := 0; i < 100; i++ {
		_, err := h.Extract()
		if err != nil {
			t.Errorf("error extracting min: %v", err)
		}

		verify(t, h, 0)
	}
}

// Verify that the heap maintains the heap constraint where the parent
// must be less than the children
func verify(t *testing.T, h *MinHeap, currentIndex int) {
	if len(h.arr) == 0 {
		return
	}
	currentValue := h.arr[currentIndex].Distance

	leftIndex := currentIndex*2 + 1
	if leftIndex < len(h.arr) {
		leftValue := h.arr[leftIndex].Distance
		if leftValue < currentValue {
			t.Errorf("Left child is less than parent. Parent: %v, Child: %v", currentValue, leftValue)

		}

		verify(t, h, leftIndex)
	}

	rightIndex := currentIndex*2 + 2
	if rightIndex < len(h.arr) {
		rightValue := h.arr[rightIndex].Distance
		if rightValue < currentValue {
			t.Errorf("Right child is less than parent. Parent: %v, Child: %v", currentValue, rightValue)
			return
		}

		verify(t, h, rightIndex)
	}

}
