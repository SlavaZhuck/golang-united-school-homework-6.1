package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
	load           int
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
		load:           0,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	var err error = nil
	if b.load < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		b.load++
	} else {
		err = fmt.Errorf(" Box capacity elapsed ")
	}
	return err
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var err error = nil
	if i < b.shapesCapacity {
		return b.shapes[i], err
	} else {
		return nil, fmt.Errorf(" Box capacity elapsed ")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var err error = nil
	var temp Shape = nil
	if b.load == 0 {
		err = fmt.Errorf(" box is empty ")
	} else if i >= 0 && i < b.load {
		temp = b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		b.load--
	} else if i < b.shapesCapacity {
		err = fmt.Errorf(" Shape at this index was not added ")
	} else {
		err = fmt.Errorf(" Index is out of range ")
	}
	return temp, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var err error = nil
	var temp Shape = nil
	if b.load == 0 {
		err = fmt.Errorf(" box is empty ")
	} else if i >= 0 && i < b.load {
		temp = b.shapes[i]
		b.shapes[i] = shape
	} else if i < b.shapesCapacity {
		err = fmt.Errorf(" Shape at this index was not added ")
	} else {
		err = fmt.Errorf(" Index is out of range ")
	}
	return temp, err

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var err error = nil
	var exist bool = false
	for i, v := range b.shapes {
		if "circle" == v.TypeOfShape() {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			b.load--
			exist = true
		}
	}

	if !exist {
		err = fmt.Errorf(" There is no circle in the box ")
	}
	return err
}
