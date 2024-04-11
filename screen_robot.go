package main

import (
	"image"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
)

type ScreenRobot struct {
	numOfDisplay  int
	displayBounds image.Rectangle
}

func NewScreenController(numOfDisplay int, rect image.Rectangle) *ScreenRobot {
	bounds := screenshot.GetDisplayBounds(numOfDisplay)
	bounds = bounds.Intersect(rect)
	return &ScreenRobot{
		numOfDisplay:  numOfDisplay,
		displayBounds: bounds,
	}
}

func (s *ScreenRobot) TakeScreenshot() (*image.RGBA, error) {
	img, err := screenshot.CaptureRect(s.displayBounds)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (s *ScreenRobot) Scroll(x, y int, args ...int) {
	robotgo.Scroll(x, y, args...)
}

func (s *ScreenRobot) ScrollDir(x int, direction ...interface{}) {
	robotgo.ScrollDir(x, direction...)
}
