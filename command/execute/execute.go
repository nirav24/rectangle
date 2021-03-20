package execute

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	"github.com/nirav24/rectangle-assessment/rectangle"
	"github.com/nirav24/rectangle-assessment/util"
	"strconv"
)

// Command is the command for the execute command
func Command() *cobra.Command {

	c := &cobra.Command{
		Use:   "execute",
		Args:  cobra.MinimumNArgs(8),
		Short: "Get Properties/behaviour of 2 rectangles",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 8 {
				return errors.New("Invalid Numbers of argument given")
			}
			r1, err := getRectangle(args[0:4])
			if err != nil {
				return errors.Wrap(err, "While creating first rectangle")
			}

			r2, err := getRectangle(args[4:8])
			if err != nil {
				return errors.Wrap(err, "While creating second rectangle")
			}

			adj, err := util.CheckAdjacency(r1, r2)
			if err != nil {
				return errors.Wrap(err, "While figuring out adjancency")
			}

			var out = struct {
				ADJACENCY    util.Output
				INTERSECTION util.Output
				CONTAINMENT  util.Output
			}{
				ADJACENCY:    adj,
				INTERSECTION: util.CheckIntersection(r1, r2),
				CONTAINMENT:  util.CheckContainment(r1, r2),
			}

			fmt.Printf("%+v", out)
			return nil
		},
	}

	return c
}

func getInt(val string) (int, error) {
	return strconv.Atoi(val)
}

func getRectangle(points []string) (rectangle.Rectangle, error) {
	if len(points) != 4 {
		return rectangle.Rectangle{}, errors.New("Invalid Numbers of argument given")
	}

	x1, err := getInt(points[0])
	if err != nil {
		return rectangle.Rectangle{}, errors.Wrap(err, "While converting first/x1 Point")
	}

	y1, err := getInt(points[1])
	if err != nil {
		return rectangle.Rectangle{}, errors.Wrap(err, "While converting second/y1 Point")
	}

	x2, err := getInt(points[2])
	if err != nil {
		return rectangle.Rectangle{}, errors.Wrap(err, "While converting third/x2 Point")
	}

	y2, err := getInt(points[3])
	if err != nil {
		return rectangle.Rectangle{}, errors.Wrap(err, "While converting fourth/y2 Point")
	}

	p1 := rectangle.NewPoint(x1, y1)
	p2 := rectangle.NewPoint(x2, y2)

	return rectangle.NewRectangle(p1, p2), nil
}
