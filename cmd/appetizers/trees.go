package main

import "errors"

type BinaryNode struct {
	Key   Element
	Depth int
	Left  *BinaryNode
	Right *BinaryNode
}

type NaryNode struct {
	Key   Element
	Depth int
	Nodes []*NaryNode
}

func (n *BinaryNode) InsertInt32(e Element) error {
	key, ok := e.(int32)
	if !ok {
		return errors.New("invalid element")
	}

	if key < n.Key.(int32) {
		if n.Left.Key == nil {
			n.Left = &BinaryNode{Key: key, Depth: n.Depth + 1}
		} else {
			n.Left.InsertInt32(key)
		}
	} else {
		if n.Left.Right == nil {
			n.Right = &BinaryNode{Key: key, Depth: n.Depth + 1}
		} else {
			n.Right.InsertInt32(key)
		}
	}
	return nil
}

func (n *BinaryNode) SearchInt32(e Element) (bool, error) {
	key, ok := e.(int32)
	if !ok {
		return false, errors.New("invalid element")
	}

	if n == nil {
		return false, nil
	}

	if key < n.Key.(int32) {
		return n.Left.SearchInt32(key)
	} else {
		return n.Right.SearchInt32(key)
	}
}
