package util

import (
	"github.com/nirav24/rectangle-assessment/rectangle"
	"github.com/pkg/errors"
)

type Output string

const (
	// Intersection
	INTERSECTION    Output = "Intersection"
	NO_INTERSECTION Output = "No Intersection"

	// Containment
	CONTAINMENT    Output = "Containment"
	NO_CONTAINMENT Output = "No Containment"

	// Adjacent
	SUBLINE_ADJACENT Output = "Sub-line Adjacent"
	PARTIAL_ADJACENT Output = "Partial Adjacent"
	PROPER_ADJACENT  Output = "Proper Adjacent"
	NO_ADJACENT      Output = "No Adjacent"

	// ERROR
	NOT_APPLICABLE Output = "N/A"
)

// Checks Whether there is any containment formed by 2 rectangles
func CheckContainment(r1 rectangle.Rectangle, r2 rectangle.Rectangle) Output {
	if r1.Contains(r2) || r2.Contains(r1) {
		return CONTAINMENT
	}
	return NO_CONTAINMENT
}

// Checks adjacency type formed by 2 rectangles
func CheckAdjacency(r1 rectangle.Rectangle, r2 rectangle.Rectangle) (Output, error) {
	connectedLines := r1.GetConnectedLinePoints(r2)
	if len(connectedLines) <= 0 {
		return NO_ADJACENT, nil
	}

	if len(connectedLines) != 4 {
		return NOT_APPLICABLE, errors.New("Invalid Points")
	}

	fLine1 := connectedLines[0]
	fLine2 := connectedLines[1]

	sLine1 := connectedLines[2]
	sLine2 := connectedLines[3]

	// Both lines have same start & end Points
	if fLine1.IsEquals(sLine1) && fLine2.IsEquals(sLine2) {
		return PROPER_ADJACENT, nil
	}

	// Both lines Share X axes
	if fLine1.IsSameOnX(sLine1) {
		if (sLine1.IsAheadOnY(fLine1) && fLine2.IsAheadOnY(sLine2)) ||
			(fLine1.IsAheadOnY(sLine1) && sLine2.IsAheadOnY(fLine2)) {
			return SUBLINE_ADJACENT, nil
		}
		return PARTIAL_ADJACENT, nil
	}

	// Both lines Share Y axes
	if fLine1.IsSameOnY(sLine1) {
		if (sLine1.IsAheadOnX(fLine1) && fLine2.IsAheadOnX(sLine2)) ||
			(fLine1.IsAheadOnX(sLine1) && sLine2.IsAheadOnX(fLine2)) {
			return SUBLINE_ADJACENT, nil
		}
		return PARTIAL_ADJACENT, nil
	}

	return NO_ADJACENT, nil
}

// Checks Intersection type formed by 2 rectangles
func CheckIntersection(r1 rectangle.Rectangle, r2 rectangle.Rectangle) Output {
	if r1.Intersect(r2) || r2.Intersect(r1) {
		return INTERSECTION
	}
	return NO_INTERSECTION
}
