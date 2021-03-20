// rectangel Package provides fns/types to interact with Rectangle shape
package rectangle

// A Point in 2D constructed using x and y point
type Point struct {
	x int
	y int
}

// Creates New Point using given x and y position in 2D axis
func NewPoint(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

// Returns true if p1 Point is ahead than p2 Point on X axis
func (p1 Point) IsAheadOnX(p2 Point) bool {
	return p1.x > p2.x
}

// Returns true if p1 Point is at same place as p2 Point on X axis
func (p1 Point) IsSameOnX(p2 Point) bool {
	return p1.x == p2.x
}

// Returns true if p1 Point is ahead than p2 Point on Y axis
func (p1 Point) IsAheadOnY(p2 Point) bool {
	return p1.y > p2.y
}

// Returns true if p1 Point is at same place as p2 Point on Y axis
func (p1 Point) IsSameOnY(p2 Point) bool {
	return p1.y == p2.y
}

// Returns true if p1 Point is ahead than p2 Point on both X and Y axises
func (p1 Point) IsAheadOnXY(p2 Point) bool {
	return p1.IsAheadOnX(p2) && p1.IsAheadOnY(p2)
}

// Returns true if p1 Point is at same place as p2 Point on 2D plane
func (p1 Point) IsEquals(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

// Represents Rectangle Shape.
type Rectangle struct {
	leftBottom Point
	rightTop   Point
}

// Construct new Rectangle using given 2 points.
// It compares 2 points to determine the position of both points
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

// Returns Left Top point of the Rectangle
func (r Rectangle) GetLeftTop() Point {
	return Point{
		x: r.leftBottom.x,
		y: r.rightTop.y,
	}
}

// Returns Right Bottom point of the Rectangle
func (r Rectangle) GetRightBottom() Point {
	return Point{
		x: r.rightTop.x,
		y: r.leftBottom.y,
	}
}

// Returns Left Bottom point of the Rectangle
func (r Rectangle) GetLeftBottom() Point {
	return r.leftBottom
}

// Returns Right Top point of the Rectangle
func (r Rectangle) GetRightTop() Point {
	return r.rightTop
}

// Returns true if r1 Rectangle contains r2 rectangle.
func (r1 Rectangle) Contains(r2 Rectangle) bool {
	return r2.GetLeftBottom().IsAheadOnXY(r1.GetLeftBottom()) && r1.GetRightTop().IsAheadOnXY(r2.GetRightTop())
}

// Returns list of points (always 4 in numbers) which will form a connected line for both rectangles
// If no line is connected, then it will return empty array
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

// Returns true if r1 intersects with r2
func (r1 Rectangle) Intersect(r2 Rectangle) bool {
	if r1.ContainsPoint(r2.GetLeftBottom()) {
		return !r1.ContainsPoint(r2.GetRightTop())
	}

	if r1.ContainsPoint(r2.GetLeftTop()) {
		return !r1.ContainsPoint(r2.GetRightBottom())
	}

	if r1.ContainsPoint(r2.GetRightBottom()) {
		return !r1.ContainsPoint(r2.GetLeftTop())
	}

	if r1.ContainsPoint(r2.GetRightTop()) {
		return !r1.ContainsPoint(r2.GetLeftBottom())
	}

	return false
}

// Returns true if a given point will be inside the rectangle
func (r1 Rectangle) ContainsPoint(p Point) bool {
	return p.IsAheadOnXY(r1.GetLeftBottom()) && r1.GetRightTop().IsAheadOnXY(p) && r1.GetRightBottom().IsAheadOnX(p) && r1.GetLeftTop().IsAheadOnY(p)
}
