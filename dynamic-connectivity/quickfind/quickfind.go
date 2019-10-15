package quickfind

import (
	"errors"
)

type QuickFinder struct {
	Components []int
}

func New(elems int) *QuickFinder {
	components := make([]int, elems)
	return &QuickFinder{
		Components: components,
	}
}

func (q *QuickFinder) Union(pointA, pointB int) error {
	if len(q.Components) < pointA+1 || len(q.Components) < pointB+1 {
		return errors.New("could not make union, point is greater than length of initialized finder")
	}

	var newValue int

	if q.Components[pointB] > 0 {
		newValue = q.Components[pointB]
	} else {
		newValue = pointB + 1
	}

	for index, componentValue := range q.Components {
		if componentValue > 0 && componentValue == q.Components[pointA] {
			q.Components[index] = newValue
		}
	}

	q.Components[pointA] = newValue
	q.Components[pointB] = newValue

	return nil
}

func (q *QuickFinder) Connected(pointA, pointB int) bool {
	if len(q.Components) < pointA+1 || len(q.Components) < pointB+1 {
		return false
	}

	valueA := q.Components[pointA]
	valueB := q.Components[pointB]

	// components list is initialised with 0 in every element. So we cannot set the value of an index as 0.
	// 0 index value is 1, which means n index value is n + 1
	return (valueA > 0 && valueB > 0) && (valueA == valueB)
}
