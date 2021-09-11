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

func TestMinHeap_ExtractEmptyHeap(t *testing.T) {
	h := NewMinHeap()

	_, err := h.Extract()
	assert.Error(t, err)
}

func TestMinHeap_ExtractHappyPath(t *testing.T) {
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
		assert.NoError(t, err)
		verify(t, h, 0)
		assert.Equal(t, expected, min.Distance)
	}
}

func TestMinHeap_LookupDoesNotExist(t *testing.T) {
	h := NewMinHeap()

	s := NewSwitch("switch")
	_, err := h.Lookup(s)
	assert.Error(t, err)
}

func TestMinHeap_UpdateSameValue(t *testing.T) {
	h := NewMinHeap()

	distances := []int{9, 23, 45, 8, 7, 6}
	switches := []*Switch{}
	for i, dist := range distances {
		s := NewSwitch(strconv.Itoa(i))
		switches = append(switches, s)
		h.Insert(s, dist)
		verify(t, h, 0)
	}

	switchToUpdate := switches[2]
	err := h.Update(switchToUpdate, distances[2])
	assert.NoError(t, err)
	verify(t, h, 0)

	dist, err := h.Lookup(switchToUpdate)
	assert.NoError(t, err)
	assert.Equal(t, distances[2], dist)
}

func TestMinHeap_UpdateSmallerValue(t *testing.T) {
	h := NewMinHeap()

	distances := []int{9, 23, 45, 8, 7, 6}
	switches := []*Switch{}
	for i, dist := range distances {
		s := NewSwitch(strconv.Itoa(i))
		switches = append(switches, s)
		h.Insert(s, dist)
		verify(t, h, 0)
	}

	switchToUpdate := switches[1]
	newDistance := 2
	err := h.Update(switchToUpdate, newDistance)
	assert.NoError(t, err)
	verify(t, h, 0)

	dist, err := h.Lookup(switchToUpdate)
	assert.NoError(t, err)
	assert.Equal(t, newDistance, dist)
}

func TestMinHeap_UpdateLargerValue(t *testing.T) {
	h := NewMinHeap()

	distances := []int{9, 23, 45, 8, 7, 6}
	switches := []*Switch{}
	for i, dist := range distances {
		s := NewSwitch(strconv.Itoa(i))
		switches = append(switches, s)
		h.Insert(s, dist)
		verify(t, h, 0)
	}

	switchToUpdate := switches[4]
	newDistance := 100
	err := h.Update(switchToUpdate, newDistance)
	assert.NoError(t, err)
	verify(t, h, 0)

	dist, err := h.Lookup(switchToUpdate)
	assert.NoError(t, err)
	assert.Equal(t, newDistance, dist)
}

func TestMinHeap_UpdateDoesNotExist(t *testing.T) {
	h := NewMinHeap()

	s := NewSwitch("switch")
	err := h.Update(s, 12)
	assert.Error(t, err)
}

func TestMinHeap_Lookup(t *testing.T) {
	h := NewMinHeap()

	distances := []int{9, 23, 45, 1}
	switches := []*Switch{}
	for i, dist := range distances {
		s := NewSwitch(strconv.Itoa(i))
		switches = append(switches, s)
		h.Insert(s, dist)
		verify(t, h, 0)
	}

	for i, s := range switches {
		dist, err := h.Lookup(s)
		assert.NoError(t, err)
		assert.Equal(t, distances[i], dist)

		verify(t, h, 0)
	}
}

// TODO: Split into specific tests that target different edge cases
func TestMinHeap_Random(t *testing.T) {
	h := NewMinHeap()

	for i := 0; i < 100; i++ {
		s := NewSwitch(strconv.Itoa(i))
		h.Insert(s, rand.Intn(100000))
		verify(t, h, 0)
	}

	for i := 0; i < 100; i++ {
		_, err := h.Extract()
		assert.NoError(t, err)

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
			return
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
