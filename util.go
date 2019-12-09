package main

import (
	"math"
)

func RotatePoints(points []float32, angle float64, x float32, y float32) []float32 {
	if len(points)%3 != 0 {
		return points
	}
	var s float32
	var c float32
	var dx float32
	var dy float32

	for i := 0; i < len(points); i += 3 {
		s = float32(math.Sin(angle * (math.Pi / 180)))
		c = float32(math.Cos(angle * (math.Pi / 180)))

		dx = (points[i] - x)
		dy = (points[i+1] - y)
		points[i] = x + dx*c - dy*s
		points[i+1] = y + dx*s + dy*c

	}
	return points
}

func TranslatePoints(points []float32, dx float32, dy float32) []float32 {
	if len(points)%3 != 0 {
		return points
	}
	for i := 0; i < len(points); i += 3 {
		points[i] = points[i] + dx
		points[i+1] = points[i+1] + dy
	}
	return points
}

func Clamp(val float32, min float32, max float32) float32 {
	if val < min {
		val = min
	}
	if val > max {
		val = max
	}
	return val
}
