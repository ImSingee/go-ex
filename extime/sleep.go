package extime

import (
	"time"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

func SleepSeconds[Number number](seconds Number) {
	time.Sleep(time.Duration(float64(seconds) * float64(time.Second)))
}
