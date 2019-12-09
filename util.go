/**************************************************************************************************
|   Assignment:  Final Project Part 3:  util.go
|      Authors:  Thomas Alexander & Cameron Larson
				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
|       Grader:  Tito Ferra & Josh Xiong
|
|       Course:  CSc 372
|   Instructor:  L. McCann
|     Due Date:  12-9-19 3:30pm
|
|  Description:  This file holds utility functions used by the asteroid, laser, and ship structs,
|				 and their associated Functions. The main use is for translating and rotating arrays
|				 of points around an axis.
|
|     Language:  GoLang
| Ex. Packages:  OpenGl, GLFW
|				 github.com/go-gl/gl/v4.1-core/gl
|				 github.com/go-gl/glfw/v3.2/glfw
|
| Deficiencies:  I know of no unsatisfied requirements and no logic errors.
**************************************************************************************************/

package main

import (
	"math"
)

/*---------------------------------------------------------------------
|  Function	 	   RotatePoints
|
|  Purpose:  	   This Function takes an array of vertices and creates a
|				   a new array of vertices rotated around an x,y coordinate
|
|  Pre-condition:  None
|
|  Post-condition: None
|
|  Parameters:
|   	points 	-- Array of floats representing vertices. Floats must be
|				   in multiples of three, stored in the following format:
|				   [x1,y1,z1,x2,y2,z2,...]
|		angle	-- Angle to rotate the object, in Degrees
|		x		-- X coordinate to rotate around
|		y		-- Y coordinate to rotate around
|
|  Returns:  	   []float32 holding the new array of points in the same format:
|				   [x1,y1,z1,x2,y2,z2,...]
*-------------------------------------------------------------------*/
func RotatePoints(points []float32, angle float64, x float32, y float32) []float32 {
	if len(points)%3 != 0 {
		return points
	}
	var s float32
	var c float32
	var dx float32
	var dy float32

	for i := 0; i < len(points); i += 3 {
		//Get sin and cos  of angle
		s = float32(math.Sin(angle * (math.Pi / 180)))
		c = float32(math.Cos(angle * (math.Pi / 180)))

		dx = (points[i] - x)   //Change in x
		dy = (points[i+1] - y) //Change in ys
		//New point location
		points[i] = x + dx*c - dy*s
		points[i+1] = y + dx*s + dy*c

	}
	return points
}

/*---------------------------------------------------------------------
|  Function	 	   TranslatePoints
|
|  Purpose:  	   This Function takes an array of vertices and creates a
|				   a new array of vertices translated by an x,y coordinate
|
|  Pre-condition:  None
|
|  Post-condition: None
|
|  Parameters:
|   	points 	-- Array of floats representing vertices. Floats must be
|				   in multiples of three, stored in the following format:
|				   [x1,y1,z1,x2,y2,z2,...]
|		dx		-- X coordinate to translate
|		dy		-- Y coordinate to tranlate
|
|  Returns:  	   []float32 holding the new array of points in the same format:
|				   [x1,y1,z1,x2,y2,z2,...]
*-------------------------------------------------------------------*/
func TranslatePoints(points []float32, dx float32, dy float32) []float32 {
	if len(points)%3 != 0 {
		return points
	}
	for i := 0; i < len(points); i += 3 {
		points[i] = points[i] + dx     //change x coordinate
		points[i+1] = points[i+1] + dy //change y coordinate
	}
	return points
}

/*---------------------------------------------------------------------
|  Function	 	   Clamp
|
|  Purpose:  	   This Function takes a value and clamps the value
|				   between a min and a max
|
|  Pre-condition:  None
|
|  Post-condition: None
|
|  Parameters:
|   	val		-- float32 representing the value to clamp
|   	min		-- the minimum value
|   	mix		-- the maximum value
|
|  Returns:  	   float32 value between the min and max
*-------------------------------------------------------------------*/
func Clamp(val float32, min float32, max float32) float32 {
	//Min
	if val < min {
		val = min
	}
	//Max
	if val > max {
		val = max
	}
	return val
}
