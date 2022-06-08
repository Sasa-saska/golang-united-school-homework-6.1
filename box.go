package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return errors.New("превышенно кол-во, нельзя добавить новую фигуру")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return b.shapes[i], nil
	} else {
		return nil, errors.New("фигуры с таким индексом нет")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		b.shapes[i] = b.shapes[len(b.shapes)-1]
		b.shapes[len(b.shapes)] = nil
		b.shapes = b.shapes[:len(b.shapes)-1]
		return b.shapes[i], nil
	} else {
		return nil, errors.New("фигуры с таким индексом нет")
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		replace := b.shapes[i]
		b.shapes[i] = shape
		return replace, nil
	} else {
		return nil, errors.New("фигуры с таким индексом нет")
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcPerimeter()
	}
	return result

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for i := range b.shapes {
		result += b.shapes[i].CalcArea()
	}
	return result

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var CirCount int
	for i := range b.shapes {
		_, ok := b.shapes[i].(*Circle)
		if ok {
			CirCount++
			b.shapes[i] = b.shapes[len(b.shapes)-1]
			b.shapes[len(b.shapes)] = nil
			b.shapes = b.shapes[:len(b.shapes)-1]
		}
	}
	if CirCount == 0 {
		return errors.New("нет кругов в боксе")
	}
	return nil
}
