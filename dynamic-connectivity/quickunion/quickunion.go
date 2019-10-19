package quickunion

import (
	"errors"
)

// need to record length of tree somewhere. It has to be calculated somehow
type QuickUnioner struct {
	Components  []int
	TreeLengths map[int]int
}

func New(elems int) *QuickUnioner {
	components := make([]int, elems)

	for index, _ := range components {
		components[index] = index
	}

	return &QuickUnioner{
		Components:  components,
		TreeLengths: map[int]int{},
	}
}

func (q *QuickUnioner) Union(pointA, pointB int) error {
	if len(q.Components) < pointA+1 || len(q.Components) < pointB+1 {
		return errors.New("could not make union, index is greater than length of initialized unioner")
	}

	historicalTreeLengthA := q.TreeLengths[q.root(pointA)]
	historicalTreeLengthB := q.TreeLengths[q.root(pointB)]
	currentTreeLengthA := q.treeLength(pointA)
	currentTreeLengthB := q.treeLength(pointB)

	var treeLengthA int
	var treeLengthB int

	if historicalTreeLengthA > currentTreeLengthA {
		treeLengthA = historicalTreeLengthA
	} else {
		treeLengthA = currentTreeLengthA
		q.TreeLengths[q.root(pointA)] = currentTreeLengthA
	}

	if historicalTreeLengthB > currentTreeLengthB {
		treeLengthB = historicalTreeLengthB
	} else {
		treeLengthB = currentTreeLengthB
		q.TreeLengths[q.root(pointB)] = currentTreeLengthB
	}

	if treeLengthA < treeLengthB {
		q.Components[pointA] = q.Components[q.root(pointB)]
	} else {
		q.Components[pointB] = q.Components[q.root(pointA)]
	}

	return nil
}

func (q *QuickUnioner) Connected(pointA, pointB int) bool {
	return q.root(pointA) == q.root(pointB)
}

func (q *QuickUnioner) root(index int) int {
	for q.Components[index] != index {
		index = q.Components[index]
	}

	return index
}

func (q *QuickUnioner) treeLength(index int) int {
	var treeLength int
	for q.Components[index] != index {
		index = q.Components[index]
		treeLength = treeLength + 1
	}

	return treeLength
}
