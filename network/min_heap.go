package network

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	arr     []Link
	indices map[*Switch]int
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	h.arr = make([]Link, 0)
	h.indices = make(map[*Switch]int)
	return h
}

func (h *MinHeap) Extract() (Link, error) {
	if len(h.arr) == 0 {
		return Link{}, errors.New("tried to extract from an empty heap")
	}

	min := h.arr[0]
	h.swap(0, len(h.arr)-1)

	// Remove the min value from the heap
	h.arr = h.arr[:len(h.arr)-1]
	delete(h.indices, min.Dest)

	h.heapify(0)

	return min, nil
}

func (h *MinHeap) heapify(currentIndex int) {
	// If there are 0 or 1 elements in the heap then there is nothing to heapify
	if len(h.arr) <= 1 {
		return
	}

	currentLink := h.arr[currentIndex]

	// Keep track of the min value and what index to swap with
	minLink := currentLink
	minIndex := currentIndex

	// See if we should swap with left child
	leftIndex := h.leftChild(currentIndex)
	if leftIndex < len(h.arr) {
		leftLink := h.arr[leftIndex]
		if leftLink.Distance < minLink.Distance {
			minLink = leftLink
			minIndex = leftIndex
		}
	}

	// See if we should swap with right child
	rightIndex := h.rightChild(currentIndex)
	if rightIndex < len(h.arr) {
		rightLink := h.arr[rightIndex]
		if rightLink.Distance < minLink.Distance {
			minLink = rightLink
			minIndex = rightIndex
		}
	}

	// If out of order, then swap and recursively heapify
	if minIndex != currentIndex {
		h.swap(currentIndex, minIndex)
		h.heapify(minIndex)
	}
}

func (h *MinHeap) Insert(node *Switch, dist int) {
	link := Link{Dest: node, Distance: dist}

	h.arr = append(h.arr, link)
	h.indices[link.Dest] = len(h.arr) - 1
	h.percolateUp(len(h.arr) - 1)
}

func (h *MinHeap) percolateUp(currentIndex int) {
	// If we are at the root of the tree, then stop
	if currentIndex == 0 {
		return
	}

	currentLink := h.arr[currentIndex]
	parentIndex := h.parent(currentIndex)
	parentLink := h.arr[parentIndex]
	if currentLink.Distance < parentLink.Distance {
		h.swap(currentIndex, parentIndex)
		h.percolateUp(parentIndex)
	}
}

func (h *MinHeap) Update(node *Switch, dist int) error {
	newLink := Link{Dest: node, Distance: dist}

	index, ok := h.indices[newLink.Dest]
	if !ok {
		return errors.New("tried to update an item that doesn't exist")
	}

	currentLink := h.arr[index]
	h.arr[index] = newLink
	if newLink.Distance > currentLink.Distance {
		h.heapify(index)
	} else if newLink.Distance < currentLink.Distance {
		h.percolateUp(index)
	}
	return nil
}

func (h *MinHeap) Lookup(s *Switch) (int, error) {
	index, ok := h.indices[s]
	if !ok {
		return 0, fmt.Errorf("tried to lookup %+v, but that doesn't exist", *s)
	}
	return h.arr[index].Distance, nil
}

func (h *MinHeap) leftChild(i int) int {
	return i*2 + 1
}

func (h *MinHeap) rightChild(i int) int {
	return i*2 + 2
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) swap(firstIndex, secondIndex int) {
	firstValue := h.arr[firstIndex]
	secondValue := h.arr[secondIndex]

	// Swap the values into the other index, update the indices map
	h.arr[secondIndex] = firstValue
	h.indices[firstValue.Dest] = secondIndex

	h.arr[firstIndex] = secondValue
	h.indices[secondValue.Dest] = firstIndex
}
