package main

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}
