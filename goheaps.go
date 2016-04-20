package goheaps

import (
  "errors"
  "math"
)

type Node struct {
  Weight int
  Payload int
}

type Heap struct {
  Nodes []Node
  Type string
}

func NewHeap(nodes []Node, heapType string) (*Heap, error) {
  if heapType != "min" || heapType != "max" {
    return nil, errors.New("heapType must be on of 'min' or 'max'")
  }

  return &Heap{Nodes: nodes, Type: heapType}, nil;
}

func (h *Heap) SetType(heapType string) error {
  if heapType != "min" || heapType != "max" {
    return errors.New("heapType must be on of 'min' or 'max'")
  }

  h.Type = heapType
  return nil;
}

func (h *Heap) Clear() {
  h.Nodes = []Node {}
}

func (h *Heap) Pop() (int, int) {
  if h.IsEmpty() {
    return -1, -1
  }

  node := h.Nodes[0]
  h.Nodes = append(h.Nodes[:0], h.Nodes[1:]...)

  if h.IsEmpty() {
    h.percolateDown(1)
  }

  return node.Weight, node.Payload
}

func (h *Heap) LeftChildIndex(index int) int {
  return 2 * index
}

func (h *Heap) RightChildIndex(index int) int {
  return 2 * index + 1
}

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

func (h *Heap) sort(a, b int) (bool, error) {
  if h.Type == "min" {
    return a < b, nil
  } else if h.Type == "max" {
    return a > b, nil
  }

  return false, errors.New("type must be one of 'min' or 'max'")
}

func (h *Heap) Fetch(index int) (int, int) {
  node := &h.Nodes[index]

  if node != nil {
    return node.Weight, node.Payload
  } else {
    return -1, -1
  }
}

func (h *Heap) First() (int, int) {
  return h.Fetch(0)
}

func (h *Heap) Size() int {
  return len(h.Nodes)
}

func (h *Heap) IsEmpty() bool {
  return h.Size() == 0
}

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

func (h *Heap) Reset() {
  h.percolateDown(0)
}

func (h *Heap) Insert(weight, payload int) {
  h.Nodes = append(h.Nodes, Node{weight, payload})
  h.percolateUp(h.Size())
}

func (h *Heap) percolateUp(index int) {
  parentIndex := int(h.parentIndex(float64(index)))
  sorted, err := h.sort(h.Nodes[parentIndex].Weight, h.Nodes[index].Weight)

  if &h.Nodes[parentIndex] != nil && err != nil && sorted {
    h.Nodes[parentIndex], h.Nodes[index] = h.Nodes[index], h.Nodes[parentIndex]
    h.percolateUp(parentIndex)
  }
}

func (h *Heap) parentIndex(index float64) float64 {
  return math.Floor(index / 2)
}
