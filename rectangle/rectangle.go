package rectangle

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p1 Point) IsAheadOnX(p2 Point) bool {
	return p1.x > p2.x
}

func (p1 Point) IsSameOnX(p2 Point) bool {
	return p1.x == p2.x
}

func (p1 Point) IsAheadOnY(p2 Point) bool {
	return p1.y > p2.y
}

func (p1 Point) IsSameOnY(p2 Point) bool {
	return p1.y == p2.y
}

func (p1 Point) IsAheadOnXY(p2 Point) bool {
	return p1.IsAheadOnX(p2) && p1.IsAheadOnY(p2)
}

func (p1 Point) IsEquals(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

type Rectangle struct {
	leftBottom Point
	rightTop   Point
}

func NewRectangle(p1 Point, p2 Point) Rectangle {
	if p1.x < p2.x {
		return Rectangle{
			leftBottom: p1,
			rightTop:   p2,
		}
	}

	return Rectangle{
		leftBottom: p2,
		rightTop:   p1,
	}
}

func (r Rectangle) GetLeftTop() Point {
	return Point{
		x: r.leftBottom.x,
		y: r.rightTop.y,
	}
}

func (r Rectangle) GetRightBottom() Point {
	return Point{
		x: r.rightTop.x,
		y: r.leftBottom.y,
	}
}

func (r Rectangle) GetLeftBottom() Point {
	return r.leftBottom
}

func (r Rectangle) GetRightTop() Point {
	return r.rightTop
}

func (r1 Rectangle) Contains(r2 Rectangle) bool {
	return r2.GetLeftBottom().IsAheadOnXY(r1.GetLeftBottom()) && r1.GetRightTop().IsAheadOnXY(r2.GetRightTop())
}

func (r1 Rectangle) GetConnectedLinePoints(r2 Rectangle) []Point {
	// Right Side connected Line
	//  -----------------
	//  |				|
	//  |				|-----
	//  |		r1   	| r2 |
	//  |				|-----
	//  |				|
	//  -----------------
	if r1.GetRightTop().x == r2.GetLeftBottom().x {
		return []Point{r1.GetRightBottom(), r1.GetRightTop(), r2.GetLeftBottom(), r2.GetLeftTop()}
	}

	// Up Side connected Line
	// 		 ------
	// 		 | r2 |
	//  -----------------
	//  |				|
	//  |				|
	//  |		r1   	|
	//  |				|
	//  |				|
	//  -----------------
	if r1.GetRightTop().y == r2.GetLeftBottom().y {
		return []Point{r1.GetLeftTop(), r1.GetRightTop(), r2.GetLeftBottom(), r2.GetRightBottom()}
	}

	// Left Side connected Line
	//  	-----------------
	//  	|				|
	// -----|				|
	// | r2	|		r1   	|
	// -----|				|
	//  	|				|
	//  	-----------------
	if r1.GetLeftBottom().x == r2.GetRightTop().x {
		return []Point{r1.GetLeftBottom(), r1.GetLeftTop(), r2.GetRightBottom(), r2.GetRightTop()}
	}

	// Bottom Side connected Line
	//  -----------------
	//  |				|
	//  |				|
	//  |		r1   	|
	//  |				|
	//  |				|
	//  -----------------
	// 		 | r2 |
	// 		 ------
	if r1.GetLeftBottom().y == r2.GetRightTop().y {
		return []Point{r1.GetLeftBottom(), r1.GetRightBottom(), r2.GetLeftTop(), r2.GetRightTop()}
	}

	// No connection
	return []Point{}
}

func (r1 Rectangle) Intersect(r2 Rectangle) bool {
	if(r1.ContainsPoint(r2.GetLeftBottom())) {
		return !r1.ContainsPoint(r2.GetRightTop())
	}
	
	if(r1.ContainsPoint(r2.GetLeftTop())) {
		return !r1.ContainsPoint(r2.GetRightBottom())
	} 

	if(r1.ContainsPoint(r2.GetRightBottom())) {
		return !r1.ContainsPoint(r2.GetLeftTop())
	}

	if(r1.ContainsPoint(r2.GetRightTop())) {
		return !r1.ContainsPoint(r2.GetLeftBottom())
	}

	return false;
}

func (r1 Rectangle) ContainsPoint(p Point) bool {
	return p.IsAheadOnXY(r1.GetLeftBottom()) && r1.GetRightTop().IsAheadOnXY(p) && r1.GetRightBottom().IsAheadOnX(p) && r1.GetLeftTop().IsAheadOnY(p)
}
