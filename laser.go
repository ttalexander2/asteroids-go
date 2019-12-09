/**************************************************************************************************
|   Assignment:  Final Project Part 3:  shaders.go
|      Authors:  Thomas Alexander & Cameron Larson
				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
|       Grader:  Tito Ferra & Josh Xiong
|
|       Course:  CSc 372
|   Instructor:  L. McCann
|     Due Date:  12-9-19 3:30pm
|
|  Description:  This file holds the laser struct and it's associated Methods. This controlls the
|				 updating and rendering of the laser on the screen. The laser objects are stored in
|				 within the Ship object, and are destroyed upon collision with asteroids, or when upon
|				 leaving the bounds of the screen
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

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	laserSpeed = 0.03 //Speed of laser
	laserLen   = 10   //Len to render laser
	laserThick = 4    //Width of laser
)

type Laser struct {
	drawable uint32       //Vertex buffer array
	points   []float32    //Vertex point array
	window   *glfw.Window //Window to render to
	x        float32      //x position on the screen
	y        float32      //y position on the screen
	dX       float32      //x velocity
	dY       float32      //y velocity
	rot      float64      //current rotation
	speed    float64      //speed
	done     bool         //bool to mark removal
}

/*---------------------------------------------------------------------
|  Method	 	   Init
|
|  Purpose:  	   This Method initializes the fields in a laser
|				   start at a x,y location and fire at a specific angle
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: Laser fields are ready to be rendered
|
|  Parameters:
|      angle 	-- Angle in degrees to rotate the laser
|	   x	 	-- initial x coordinate
|	   y		-- initial y coordinate
|
|  Returns:  	    None :)
*-------------------------------------------------------------------*/
func (l *Laser) Init(angle float64, x float32, y float32) {
	l.rot = angle
	l.x = x
	l.y = y
	l.getPoints()                                      //Generate point array
	l.points = RotatePoints(l.points, l.rot, l.x, l.y) //rotate point array
	l.speed = laserSpeed
	l.done = false
	l.drawable = makeVertexArrayObj(l.points) //Generate VBO
}

/*---------------------------------------------------------------------
|  Method	 	   Update
|
|  Purpose:  	   This Method updates the positions and velocities
|				   of the laser
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: Laser fields are ready to be rendered
|
|  Parameters:	   None
|
|  Returns:  	   None
*-------------------------------------------------------------------*/
func (s *Laser) Update() {
	s.checkEdges()
	//Calculate change in x based on angle
	s.dX = float32(math.Cos(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.speed)
	s.dY = float32(math.Sin(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.speed)
	s.x += s.dX
	s.y += s.dY
	//Update point array, generate VAO (vertex array object)
	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)
}

/*---------------------------------------------------------------------
|  Method	 	   Draw
|
|  Purpose:  	   This Method renders the laser on the screen
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the laser must have a VAO (vertex array object)
|
|  Post-condition: Laser fields are ready to be rendered
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Laser) Draw() {
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(s.points)/3))
}

/*---------------------------------------------------------------------
|  Method	 	   checkEdges
|
|  Purpose:  	   This Method checks whether a has traveled past
|				   the bounds of the screen
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the laser must have a VAO (vertex array object)
|
|  Post-condition: Lasers outside screen bounds are marked for removal
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Laser) checkEdges() {
	if s.x > 1 || s.x < -1 || s.y > 1 || s.y < -1 {
		s.done = true
	}
}

/*---------------------------------------------------------------------
|  Method	 	   getPoints
|
|  Purpose:  	   This Method generates an array of points for the laser
|				   to be converted into a VAO
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: None
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Laser) getPoints() {
	s.points = []float32{
		(s.x*Width - laserThick) / Width, (s.y * Height) / Height, 0,
		(s.x*Width - laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y * Width) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width - laserThick) / Width, (s.y * Width) / Height, 0,
	}
}

/*---------------------------------------------------------------------
|  Method	 	   hits
|
|  Purpose:  	   This Method checks whether a laser has collided with
|				   an Asteroid. The algorithm uses the radius of the
|				   asteroid to measure for collision
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|				   The asteroid struct must be initialized
|
|  Post-condition: None
|
|  Parameters:
|   		a 	-- Asteroid to check for collision
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Laser) hits(a Asteroid) bool {
	var d float32
	//Calculate distance between two points
	d = float32(math.Sqrt((math.Pow(float64((s.x-a.x)*Width), 2.0) + math.Pow(float64((s.y-a.y)*Height), 2))))
	//if distance is less than radius
	if d < a.r {
		return true
	}
	return false
}
