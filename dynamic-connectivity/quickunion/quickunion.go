package quickunion

import (
	"errors"
)

type QuickUnioner struct {
	Components []int
}

func New(elems int) *QuickUnioner {
	components := make([]int, elems)

	for index, _ := range components {
		components[index] = index
	}

	return &QuickUnioner{
		Components: components,
	}
}

func (q *QuickUnioner) Union(pointA, pointB int) error {
	if len(q.Components) < pointA+1 || len(q.Components) < pointB+1 {
		return errors.New("could not make union, index is greater than length of initialized unioner")
	}

	q.Components[pointA] = q.Components[pointB]

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
