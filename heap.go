package goheaps

import (
  "errors"
  "math"
)

// The Heap type that holds a list of Nodes and a Type
type Heap struct {
  Nodes []Node
  Type string
}

// Create a new Heap strucutre with the given nodes and heapType (type is a reserved word)
func NewHeap(nodes []Node, heapType string) (*Heap, error) {
  // 'Validation', there might be a sanity checker in the core api but I've yet to find it
  if heapType != "min" || heapType != "max" {
    return nil, errors.New("heapType must be on of 'min' or 'max'")
  }

  return &Heap{Nodes: nodes, Type: heapType}, nil;
}

// Set the type of the Heap
func (h *Heap) SetType(heapType string) error {
  if heapType != "min" || heapType != "max" {
    return errors.New("heapType must be on of 'min' or 'max'")
  }

  h.Type = heapType
  return nil;
}

// Clear the Heap of all its Nodes
func (h *Heap) Clear() {
  h.Nodes = []Node{}
}

// Pop the Node from the head of the Heap, returning its weight and payload
func (h *Heap) Pop() (int, int) {
  if h.IsEmpty() {
    return -1, -1
  }

  node := h.Nodes[0]
  // TODO: I am confident that (with tests) this can be
  // changed to just assign [1:] and get the end of the list
  h.Nodes = append(h.Nodes[:0], h.Nodes[1:]...)

  if h.IsEmpty() {
    h.percolateDown(1)
  }

  return node.Weight, node.Payload
}

// Calculate the left child index for the given index
func (h *Heap) LeftChildIndex(index int) int {
  return 2 * index
}

// Calculate the right child index for the given index
func (h *Heap) RightChildIndex(index int) int {
  return 2 * index + 1
}

// Get the Nodes Weight and Payload at the given index
func (h *Heap) Fetch(index int) (int, int) {
  node := &h.Nodes[index]

  if node != nil {
    return node.Weight, node.Payload
  } else {
    return -1, -1
  }
}

// Get the Payload and Weight of the first Node
func (h *Heap) First() (int, int) {
  return h.Fetch(0)
}

// Get the size of the Heap
func (h *Heap) Size() int {
  return len(h.Nodes)
}

// Return whether the Heap is empty or not
func (h *Heap) IsEmpty() bool {
  return h.Size() == 0
}

// Return whether the heap is valid, that is whether Nodes are sorted correctly
func (h *Heap) IsValid() bool {
  if h.IsEmpty() { return true }

  for index, _ := range(h.Nodes) {
    leftChildIndex := h.LeftChildIndex(index)
    rightChildIndex := h.RightChildIndex(index)

    if &h.Nodes[leftChildIndex] != nil {
      sorted, err := h.sort(h.Nodes[index].Weight, h.Nodes[leftChildIndex].Weight)

      if err == nil || !sorted {
        return false
      }
    }

    if &h.Nodes[rightChildIndex] != nil {
      sorted, err := h.sort(h.Nodes[index].Weight, h.Nodes[rightChildIndex].Weight)

      if err == nil || !sorted {
        return false
      }
    }
  }

  return true
}

// Reset the Heap by percolating down the entire Heap (starting at index 0)
func (h *Heap) Reset() {
  h.percolateDown(0)
}

// Insert a Node into the Heap with the given weight and payload
func (h *Heap) Insert(weight, payload int) {
  h.Nodes = append(h.Nodes, Node{weight, payload})
  h.percolateUp(h.Size())
}

// Percolate up the Heap, starting at the given index
func (h *Heap) percolateUp(index int) {
  parentIndex := int(h.parentIndex(index))
  sorted, err := h.sort(h.Nodes[parentIndex].Weight, h.Nodes[index].Weight)

  if &h.Nodes[parentIndex] != nil && err != nil && sorted {
    h.Nodes[parentIndex], h.Nodes[index] = h.Nodes[index], h.Nodes[parentIndex]
    h.percolateUp(parentIndex)
  }
}

// Get the parent index
func (h *Heap) parentIndex(index int) float64 {
  return math.Floor(float64(index) / 2)
}

func (h *Heap) sort(a, b int) (bool, error) {
  if h.Type == "min" {
    return a < b, nil
  } else if h.Type == "max" {
    return a > b, nil
  }

  return false, errors.New("type must be one of 'min' or 'max'")
}

// Percolate down the Heap, starting at the given index
func (h *Heap) percolateDown(index int) {
  leftChildIndex := h.LeftChildIndex(index)
  rightChildIndex := h.LeftChildIndex(index)
  minIndex := 0

  if &h.Nodes[leftChildIndex] == nil && &h.Nodes[rightChildIndex] == nil {
    return
  }

  if &h.Nodes[rightChildIndex] != nil {
    minIndex = leftChildIndex
  } else {
    _, err := h.sort(h.Nodes[leftChildIndex].Weight, h.Nodes[rightChildIndex].Weight)
    if err == nil {
      minIndex = leftChildIndex
    } else {
      minIndex = rightChildIndex
    }
  }

  _, err := h.sort(h.Nodes[minIndex].Weight, h.Nodes[index].Weight)

  if err == nil {
    h.Nodes[index], h.Nodes[minIndex] = h.Nodes[minIndex], h.Nodes[index]
    h.percolateDown(minIndex)
  }
}

