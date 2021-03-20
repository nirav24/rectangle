package rectangle_test

import (
	"testing"

	"github.com/nirav24/rectangle-assessment/rectangle"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestPoints(t *testing.T) {

	testID := 0
	t.Logf("Test: %d:\tGiven the 2 points where p2 point is ahead of p1 on both axis", testID)
	{
		p1 := rectangle.NewPoint(0, 0)
		p2 := rectangle.NewPoint(10, 10)

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X axis.", testID)
		{
			if !p2.IsAheadOnX(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on X axis", Success, testID)

			if p1.IsAheadOnX(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on Y axis.", testID)
		{
			if !p2.IsAheadOnY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on Y axis", Success, testID)

			if p1.IsAheadOnY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X and Y axis.", testID)
		{
			if !p2.IsAheadOnXY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on X&Y axis", Success, testID)

			if p1.IsAheadOnXY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Success, testID)
		}
	}

	testID = 1
	t.Logf("Test: %d:\tGiven the 2 points where p2 point is ahead of p1 on Y axis but not on X", testID)
	{
		p1 := rectangle.NewPoint(0, 0)
		p2 := rectangle.NewPoint(0, 10)

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X axis.", testID)
		{
			if p2.IsAheadOnX(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X axis", Success, testID)

			if p1.IsAheadOnX(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on Y axis.", testID)
		{
			if !p2.IsAheadOnY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 be ahead of p1 on Y axis", Success, testID)

			if p1.IsAheadOnY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X and Y axis.", testID)
		{
			if p2.IsAheadOnXY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X&Y axis", Success, testID)

			if p1.IsAheadOnXY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Success, testID)
		}
	}

	testID = 2
	t.Logf("Test: %d:\tGiven the 2 points where p2 point is ahead of p1 on X axis but not on Y", testID)
	{
		p1 := rectangle.NewPoint(0, 0)
		p2 := rectangle.NewPoint(10, 0)

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X axis.", testID)
		{
			if !p2.IsAheadOnX(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X axis", Success, testID)

			if p1.IsAheadOnX(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on Y axis.", testID)
		{
			if p2.IsAheadOnY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on Y axis", Success, testID)

			if p1.IsAheadOnY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on Y axis", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare P1 and P2 on X and Y axis.", testID)
		{
			if p2.IsAheadOnXY(p1) {
				t.Fatalf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p2 not be ahead of p1 on X&Y axis", Success, testID)

			if p1.IsAheadOnXY(p2) {
				t.Fatalf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould p1 not be ahead of p2 on X&Y axis", Success, testID)
		}
	}
}

func TestRectangle(t *testing.T) {
	testID := 0
	t.Logf("Test: %d:\tGiven the 2 points, check rectangle's Factory method", testID)
	{
		lb1 := rectangle.NewPoint(0, 0)
		rt1 := rectangle.NewPoint(10, 10)
		rc1 := rectangle.NewRectangle(lb1, rt1)

		t.Logf("\tTest: %d:\tCompare proper leftBottom point is set.", testID)
		{
			lb := rc1.GetLeftBottom()
			if !lb.IsEquals(lb1) {
				t.Fatalf("\t%s\tTest: %d:\tShould left Bottom Point of Rectable be same as lb1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould left Bottom Point of Rectable be same as lb1", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare proper rightTop point is set.", testID)
		{
			rt := rc1.GetRightTop()
			if !rt.IsEquals(rt1) {
				t.Fatalf("\t%s\tTest: %d:\tShould rightTop Point of Rectable be same as rt1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rightTop Point of Rectable be same as rt1", Success, testID)
		}

	}

	testID = 1
	t.Logf("Test: %d:\tGiven the 2 points, check rectangle's Factory method when points passed in opposite order", testID)
	{
		lb1 := rectangle.NewPoint(0, 0)
		rt1 := rectangle.NewPoint(10, 10)
		rc1 := rectangle.NewRectangle(rt1, lb1)

		t.Logf("\tTest: %d:\tCompare proper leftBottom point is set.", testID)
		{
			lb := rc1.GetLeftBottom()
			if !lb.IsEquals(lb1) {
				t.Fatalf("\t%s\tTest: %d:\tShould left Bottom Point of Rectable be same as lb1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould left Bottom Point of Rectable be same as lb1", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare proper rightTop point is set.", testID)
		{
			rt := rc1.GetRightTop()
			if !rt.IsEquals(rt1) {
				t.Fatalf("\t%s\tTest: %d:\tShould rightTop Point of Rectable be same as rt1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rightTop Point of Rectable be same as rt1", Success, testID)
		}

	}

	testID = 2
	t.Logf("Test: %d:\tGiven the 2 points, check rectangle's Other Points", testID)
	{
		lb1 := rectangle.NewPoint(0, 0)
		lt1 := rectangle.NewPoint(0, 10)
		rt1 := rectangle.NewPoint(10, 10)
		rb1 := rectangle.NewPoint(10, 0)
		rc1 := rectangle.NewRectangle(rt1, lb1)

		t.Logf("\tTest: %d:\tCompare proper leftTop point is set.", testID)
		{
			lt := rc1.GetLeftTop()
			if !lt.IsEquals(lt1) {
				t.Fatalf("\t%s\tTest: %d:\tShould left Top Point of Rectable be same as lt1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould left Top Point of Rectable be same as lt1", Success, testID)
		}

		t.Logf("\tTest: %d:\tCompare proper Right Bottom point is set.", testID)
		{
			rb := rc1.GetRightBottom()
			if !rb.IsEquals(rb1) {
				t.Fatalf("\t%s\tTest: %d:\tShould left Right Bottom of Rectable be same as rb1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould left Right Bottom of Rectable be same as rb1", Success, testID)
		}
	}

	testID = 3
	t.Logf("Test: %d:\tGiven the 2 rectangle, check a rectangle is contained within the other rectangle", testID)
	{
		lb1 := rectangle.NewPoint(0, 0)
		rt1 := rectangle.NewPoint(10, 10)
		rc1 := rectangle.NewRectangle(rt1, lb1)

		lb2 := rectangle.NewPoint(1, 1)
		rt2 := rectangle.NewPoint(5, 5)
		rc2 := rectangle.NewRectangle(rt2, lb2)

		lb3 := rectangle.NewPoint(11, 0)
		rt3 := rectangle.NewPoint(20, 5)
		rc3 := rectangle.NewRectangle(rt3, lb3)

		t.Logf("\tTest: %d:\tCheck rc1 contains rc2.", testID)
		{
			if !rc1.Contains(rc2) {
				t.Fatalf("\t%s\tTest: %d:\tShould rc1 contains rc2", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc1 contains rc2", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc2 contains rc1.", testID)
		{
			if rc2.Contains(rc1) {
				t.Fatalf("\t%s\tTest: %d:\tShould rc2 not contains rc1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc2 not contains rc1", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 contains rc3.", testID)
		{
			if rc1.Contains(rc3) {
				t.Fatalf("\t%s\tTest: %d:\tShould rc1 not contains rc3", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc1 not contains rc3", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc3 contains rc1.", testID)
		{
			if rc3.Contains(rc1) {
				t.Fatalf("\t%s\tTest: %d:\tShould rc3 not contains rc1", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc3 not contains rc1", Success, testID)
		}
	}

	testID = 4
	t.Logf("Test: %d:\tGiven the 2 rectangle, check rectangles have connected lines", testID)
	{

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 doesn't have connected lines.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(1, 1)
			rt2 := rectangle.NewPoint(5, 5)
			rc2 := rectangle.NewRectangle(rt2, lb2)
			connectedPoints := rc1.GetConnectedLinePoints(rc2)
			if len(connectedPoints) > 0 {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Right line as connected.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(10, 1)
			rt2 := rectangle.NewPoint(15, 5)
			rc2 := rectangle.NewRectangle(rt2, lb2)
			connectedPoints := rc1.GetConnectedLinePoints(rc2)
			if len(connectedPoints) < 0 {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Success, testID)

			if !connectedPoints[0].IsEquals(rc1.GetRightBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 RightBottom Should be first Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 RightBottom Should be first Point", Success, testID)

			if !connectedPoints[1].IsEquals(rc1.GetRightTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 RightTop Should be second Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 RightTop Should be second Point", Success, testID)

			if !connectedPoints[2].IsEquals(rc2.GetLeftBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 LeftBottom Should be third Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 LeftBottom Should be third Point", Success, testID)

			if !connectedPoints[3].IsEquals(rc2.GetLeftTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 LeftTop Should be forth Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 LeftTop Should be forth Point", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Top line as connected.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(5, 10)
			rt2 := rectangle.NewPoint(15, 15)
			rc2 := rectangle.NewRectangle(rt2, lb2)

			connectedPoints := rc1.GetConnectedLinePoints(rc2)
			if len(connectedPoints) < 0 {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Success, testID)

			if !connectedPoints[0].IsEquals(rc1.GetLeftTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 LeftTop Should be first Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 LeftTop Should be first Point", Success, testID)

			if !connectedPoints[1].IsEquals(rc1.GetRightTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 RightTop Should be second Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 RightTop Should be second Point", Success, testID)

			if !connectedPoints[2].IsEquals(rc2.GetLeftBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 LeftBottom Should be third Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 LeftBottom Should be third Point", Success, testID)

			if !connectedPoints[3].IsEquals(rc2.GetRightBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 RightBottom Should be forth Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 RightBottom Should be forth Point", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Left line as connected.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(-5, 5)
			rt2 := rectangle.NewPoint(0, 8)
			rc2 := rectangle.NewRectangle(rt2, lb2)
			connectedPoints := rc1.GetConnectedLinePoints(rc2)
			if len(connectedPoints) < 0 {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Success, testID)

			if !connectedPoints[0].IsEquals(rc1.GetLeftBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 LeftBottom Should be first Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 LeftBottom Should be first Point", Success, testID)

			if !connectedPoints[1].IsEquals(rc1.GetLeftTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 LeftTop Should be second Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 LeftTop Should be second Point", Success, testID)

			if !connectedPoints[2].IsEquals(rc2.GetRightBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 RightBottom Should be third Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 RightBottom Should be third Point", Success, testID)

			if !connectedPoints[3].IsEquals(rc2.GetRightTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 RightTop Should be forth Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 RightTop Should be forth Point", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Bottom line as connected.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(5, -5)
			rt2 := rectangle.NewPoint(8, 0)
			rc2 := rectangle.NewRectangle(rt2, lb2)
			connectedPoints := rc1.GetConnectedLinePoints(rc2)
			if len(connectedPoints) < 0 {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have connected lines", Success, testID)

			if !connectedPoints[0].IsEquals(rc1.GetLeftBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 LeftBottom Should be first Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 LeftBottom Should be first Point", Success, testID)

			if !connectedPoints[1].IsEquals(rc1.GetRightBottom()) {
				t.Fatalf("\t%s\tTest: %d:\trc1 RightBottom Should be second Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 RightBottom Should be second Point", Success, testID)

			if !connectedPoints[2].IsEquals(rc2.GetLeftTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 LeftTop Should be third Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 LeftTop Should be third Point", Success, testID)

			if !connectedPoints[3].IsEquals(rc2.GetRightTop()) {
				t.Fatalf("\t%s\tTest: %d:\trc2 RightTop Should be forth Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc2 RightTop Should be forth Point", Success, testID)
		}

	}

	testID = 5
	t.Logf("Test: %d:\tGiven the 1 rectangle and 1 point, check rectangle Contains Point", testID)
	{

		t.Logf("\tTest: %d:\tRectangle Should Contain Point.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			p := rectangle.NewPoint(1, 1)

			if !rc1.ContainsPoint(p) {
				t.Fatalf("\t%s\tTest: %d:\tRectangle Should Contain Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tRectangle Should Contain Point", Success, testID)
		}

		t.Logf("\tTest: %d:\tRectangle Should not Contain Point.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			p := rectangle.NewPoint(-1, -1)

			if rc1.ContainsPoint(p) {
				t.Fatalf("\t%s\tTest: %d:\tRectangle Should not Contain Point", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tRectangle Should not Contain Point", Success, testID)
		}
	}

	testID = 6
	t.Logf("Test: %d:\tGiven the 2 rectangle, check rectangles have intersection", testID)
	{

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have intersection.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(1, 1)
			rt2 := rectangle.NewPoint(15, 15)
			rc2 := rectangle.NewRectangle(rt2, lb2)

			if !rc1.Intersect(rc2) {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have intersection", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should have intersection", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have intersection.", testID)
		{
			var pointsArry = []struct {
				x1 int
				y1 int
				x2 int
				y2 int
				a1 int
				b1 int
				a2 int
				b2 int
			}{
				{0, 0, 10, 10, 1, 1, 15, 15},
				{0, 0, 10, 10, -1, -1, 5, 5},
				{0, 0, 10, 10, -10, -10, 1, 5},
			}
			for _, points := range pointsArry {
				t.Run("Table test with different points", func(t *testing.T) {
					lb1 := rectangle.NewPoint(points.x1, points.y1)
					rt1 := rectangle.NewPoint(points.x2, points.y2)
					rc1 := rectangle.NewRectangle(rt1, lb1)

					lb2 := rectangle.NewPoint(points.a1, points.b1)
					rt2 := rectangle.NewPoint(points.a2, points.b2)
					rc2 := rectangle.NewRectangle(rt2, lb2)

					if !rc1.Intersect(rc2) {
						t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have intersection", Failed, testID)
					}
					t.Logf("\t%s\tTest: %d:\trc1 & rc2 should have intersection", Success, testID)
				})
			}
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 doesn't have intersection.", testID)
		{
			var pointsArry = []struct {
				x1 int
				y1 int
				x2 int
				y2 int
				a1 int
				b1 int
				a2 int
				b2 int
			}{
				{0, 0, 10, 10, 11, 1, 15, 15},
				{0, 0, 10, 10, -1, -1, 15, 15},
			}
			for _, points := range pointsArry {
				t.Run("Table test with different points", func(t *testing.T) {
					lb1 := rectangle.NewPoint(points.x1, points.y1)
					rt1 := rectangle.NewPoint(points.x2, points.y2)
					rc1 := rectangle.NewRectangle(rt1, lb1)

					lb2 := rectangle.NewPoint(points.a1, points.b1)
					rt2 := rectangle.NewPoint(points.a2, points.b2)
					rc2 := rectangle.NewRectangle(rt2, lb2)

					if rc1.Intersect(rc2) {
						t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should not have intersection", Failed, testID)
					}
					t.Logf("\t%s\tTest: %d:\trc1 & rc2 should not have intersection", Success, testID)
				})
			}
		}
	}
}
