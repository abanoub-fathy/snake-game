package internal

type PointList []*Point

func NewPointList(points ...*Point) *PointList {
	return (*PointList)(&points)
}

func (pl *PointList) GetLength() int {
	return len(*pl)
}

func (pl *PointList) GetFirstPoint() *Point {
	return (*pl)[0]
}

func (pl *PointList) GetLastPoint() *Point {
	return (*pl)[pl.GetLength()-1]
}

func (pl *PointList) GetPointListWithoutLastPoint() *PointList {
	return NewPointList((*pl)[:pl.GetLength()-1]...)
}

func (pl *PointList) IncludesPoint(p Point) bool {
	for _, pointInList := range *pl {
		if p.IsLocatedOn(pointInList) {
			return true
		}
	}

	return false
}

func (pl *PointList) Push(point *Point) {
	*pl = append(*pl, point)
}

func (pl *PointList) Shift() {
	*pl = (*pl)[1:]
}
