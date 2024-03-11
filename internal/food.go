package internal

type Food struct {
	Point *Point
}

func NewFood(point *Point) *Food {
	return &Food{
		Point: point,
	}
}
