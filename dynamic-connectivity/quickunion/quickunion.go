package quickunion

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
	q.Components[pointA] = q.Components[pointB]

	return nil
}

func (q *QuickUnioner) Connected(pointA, pointB int) bool {
	var pointAHead int
	while pointA
	return true
}
