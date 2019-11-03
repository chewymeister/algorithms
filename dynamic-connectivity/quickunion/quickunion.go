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

	treeLengthA := q.TreeLengths[q.root(pointA)]
	treeLengthB := q.TreeLengths[q.root(pointB)]

	if treeLengthA < treeLengthB {
		q.Components[q.root(pointA)] = q.Components[q.root(pointB)]
		q.TreeLengths[q.root(pointA)] = treeLengthA + 1
	} else {
		q.Components[q.root(pointB)] = q.Components[q.root(pointA)]
		q.TreeLengths[q.root(pointB)] = treeLengthB + 1
	}

	return nil
}

func (q *QuickUnioner) Connected(pointA, pointB int) bool {
	return q.root(pointA) == q.root(pointB)
}

func (q *QuickUnioner) root(index int) int {
	for q.Components[index] != index {
		index = q.Components[index]
		q.Components[index] = q.Components[q.Components[index]]
	}

	return index
}
