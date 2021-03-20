package util_test

import (
	"testing"

	"github.com/nirav24/rectangle-assessment/rectangle"
	"github.com/nirav24/rectangle-assessment/util"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestUtil(t *testing.T) {
	testID := 0

	t.Logf("Test: %d:\tGiven the 2 rectangle, CheckContainment method", testID)
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
			cnt := util.CheckContainment(rc1, rc2)
			if cnt != util.CONTAINMENT {
				t.Fatalf("\t%s\tTest: %d:\tShould rc1 contains rc2", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc1 contains rc2", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 contains rc2 with unordered args.", testID)
		{
			cnt := util.CheckContainment(rc2, rc1)
			if cnt != util.CONTAINMENT {
				t.Fatalf("\t%s\tTest: %d:\tShould rc1 contains rc2", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc1 contains rc2", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 contains rc3.", testID)
		{
			cnt := util.CheckContainment(rc1, rc3)
			if cnt != util.NO_CONTAINMENT {
				t.Fatalf("\t%s\tTest: %d:\tShould rc1 not contains rc3", Failed, testID)
			}
			t.Logf("\t%s\tTest: %d:\tShould rc1 not contains rc3", Success, testID)
		}

	}

	testID = 1
	t.Logf("Test: %d:\tGiven the 2 rectangle, CheckAdjacency method", testID)
	{

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 should have no adjacency.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(1, 1)
			rt2 := rectangle.NewPoint(5, 5)
			rc2 := rectangle.NewRectangle(rt2, lb2)
			adj, err := util.CheckAdjacency(rc1, rc2)
			if err != nil {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have NO_ADJACENT: %v", Failed, testID, err)
			}
			if adj != util.NO_ADJACENT {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should NO_ADJACENT: %s", Failed, testID, adj)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should NO_ADJACENT", Success, testID)

		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Proper adjacency.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(10, 0)
			rt2 := rectangle.NewPoint(10, 10)
			rc2 := rectangle.NewRectangle(rt2, lb2)

			adj, err := util.CheckAdjacency(rc1, rc2)
			if err != nil {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have PROPER_ADJACENT: %v", Failed, testID, err)
			}
			if adj != util.PROPER_ADJACENT {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should PROPER_ADJACENT: %s", Failed, testID, adj)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should PROPER_ADJACENT", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Subline adjacency.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(10, 1)
			rt2 := rectangle.NewPoint(10, 8)
			rc2 := rectangle.NewRectangle(rt2, lb2)

			adj, err := util.CheckAdjacency(rc1, rc2)
			if err != nil {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have SUBLINE_ADJACENT: %v", Failed, testID, err)
			}
			if adj != util.SUBLINE_ADJACENT {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should SUBLINE_ADJACENT: %s", Failed, testID, adj)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should SUBLINE_ADJACENT", Success, testID)
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 have Partial adjacency.", testID)
		{
			lb1 := rectangle.NewPoint(0, 0)
			rt1 := rectangle.NewPoint(10, 10)
			rc1 := rectangle.NewRectangle(rt1, lb1)

			lb2 := rectangle.NewPoint(10, 1)
			rt2 := rectangle.NewPoint(10, 15)
			rc2 := rectangle.NewRectangle(rt2, lb2)

			adj, err := util.CheckAdjacency(rc1, rc2)
			if err != nil {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should have PARTIAL_ADJACENT: %v", Failed, testID, err)
			}
			if adj != util.PARTIAL_ADJACENT {
				t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should PARTIAL_ADJACENT: %s", Failed, testID, adj)
			}
			t.Logf("\t%s\tTest: %d:\trc1 & rc2 should PARTIAL_ADJACENT", Success, testID)
		}

	}


	testID = 2
	t.Logf("Test: %d:\tGiven the 2 rectangle, CheckIntersection method", testID)
	{

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 should have intersection.", testID)
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
				t.Run("Table test with various points", func(t *testing.T) {
					lb1 := rectangle.NewPoint(points.x1, points.y1)
					rt1 := rectangle.NewPoint(points.x2, points.y2)
					rc1 := rectangle.NewRectangle(rt1, lb1)

					lb2 := rectangle.NewPoint(points.a1, points.b1)
					rt2 := rectangle.NewPoint(points.a2, points.b2)
					rc2 := rectangle.NewRectangle(rt2, lb2)
					intersection := util.CheckIntersection(rc1, rc2)

					if (intersection != util.INTERSECTION) {
						t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should INTERSECTION: %s", Failed, testID, intersection)
					}
					t.Logf("\t%s\tTest: %d:\trc1 & rc2 should INTERSECTION", Success, testID)
				})
			}
		}

		t.Logf("\tTest: %d:\tCheck rc1 & rc2 should not have intersection.", testID)
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
				t.Run("Table test with various points", func(t *testing.T) {
					lb1 := rectangle.NewPoint(points.x1, points.y1)
					rt1 := rectangle.NewPoint(points.x2, points.y2)
					rc1 := rectangle.NewRectangle(rt1, lb1)

					lb2 := rectangle.NewPoint(points.a1, points.b1)
					rt2 := rectangle.NewPoint(points.a2, points.b2)
					rc2 := rectangle.NewRectangle(rt2, lb2)
					intersection := util.CheckIntersection(rc1, rc2)

					if (intersection != util.NO_INTERSECTION) {
						t.Fatalf("\t%s\tTest: %d:\trc1 & rc2 should NO_INTERSECTION: %s", Failed, testID, intersection)
					}
					t.Logf("\t%s\tTest: %d:\trc1 & rc2 should NO_INTERSECTION", Success, testID)
				})
			}
		}

	}
}
