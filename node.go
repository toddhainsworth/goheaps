package goheaps

// The Node structure that consists of a Weight and Payload
// TODO: Payload should be polymorphic (possible in Go?) I've tried adding it as a byte
// but a byte won't be enough space to store all possible values
type Node struct {
  Weight int
  Payload int
}

