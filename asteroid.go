/**************************************************************************************************
|   Assignment:  Final Project Part 3:  asteroid.go
|      Authors:  Thomas Alexander & Cameron Larson
				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
|       Grader:  Tito Ferra & Josh Xiong
|
|       Course:  CSc 372
|   Instructor:  L. McCann
|     Due Date:  12-9-19 3:30pm
|
|  Description:  This file holds the asteroid struct and it's associated methods. This controlls the
|				 updating and rendering of the asteroids on the screen. The asteroid file also
|				 holds the logic to generate a (somewhat) random number of asteroids with varying
|				 shapes and sizes. The asteroids are created with somewhat random velocity and
|				 positions. Additionally, asteroids also have the ability to break apart into 2
|				 smaller asteroid objects.
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
	"math/rand"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Asteroid struct {
	drawable uint32       //Vertex buffer array
	points   []float32    //Vertex point array
	init     []float32    //Initial positions
	window   *glfw.Window //Window to render to
	x        float32      //x position on the screen
	y        float32      //y position on the screen
	dX       float32      //x velocity
	dY       float32      //y velocity
	r        float32      //radius of ship
	dRot     float64      //change in rotation
	rot      float64      //rotation of ship
	done     bool         //Marker for removal
}

/*---------------------------------------------------------------------
|  Method	 	   Update
|
|  Purpose:  	   This Method updates the positions and velocities
|				   of the asteroids
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: Asteroid fields are ready to be rendered
|
|  Parameters:	   None
|
|  Returns:  	   None
*-------------------------------------------------------------------*/
func (s *Asteroid) Update() {
	s.checkEdges() //Check window bounds

	//Update coordinates w/ velocity
	s.x += s.dX
	s.y += s.dY

	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)
}

/*---------------------------------------------------------------------
|  Method	 	   Draw
|
|  Purpose:  	   This Method renders the asteroid on the screen
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the laser must have a VAO (vertex array object)
|
|  Post-condition: Asteroid fields are ready to be rendered
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Asteroid) Draw() {
	s.drawable = makeVertexArrayObj(s.points)
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.POINTS, 0, int32(len(s.points)/3))
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(len(s.points)/3))
}

/*---------------------------------------------------------------------
|  Method	 	   checkEdges
|
|  Purpose:  	   This Method checks whether the asteroid has traveled past
|				   the bounds of the screen. If it has, it changes the
|				   location to the opposite side
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the asteroid must have a VAO (vertex array object)
|
|  Post-condition: Asteroid's location is now on the other side of the screen
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Asteroid) checkEdges() {
	//Right bound
	if s.x > 1 {
		s.getPoints()
		s.points = TranslatePoints(s.points, -2, 0)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
		s.x = -1
	} else if s.x < -1 { //Left bound
		s.getPoints()
		s.points = TranslatePoints(s.points, 2, 0)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
		s.x = 1
	}
	//Upper bound
	if s.y > 1 {
		s.y = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, 0, -2)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.y < -1 { //Lower bound
		s.y = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, 0, 2)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
}

/*---------------------------------------------------------------------
|  Method	 	   getPoints
|
|  Purpose:  	   This Method generates an array of points for the asteroid
|				   to be converted into a VAO
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: None
|
|  Parameters:	   None
|
|  Returns:  	   None
*-------------------------------------------------------------------*/
func (s *Asteroid) getPoints() {
	s.points = s.init
	s.drawable = makeVertexArrayObj(s.points)
}

/*---------------------------------------------------------------------
|  Method	 	   generatePoints
|
|  Purpose:  	   This Method generates a random array of points
|				   by creating a circular polygon with a random number
|				   of vertices, and changes the radius of each vertex
|				   to create a random asteroid shape, to be converted
|				   into a VAO
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: None
|
|  Parameters:	   None
|
|  Returns:  	   None
*-------------------------------------------------------------------*/
func (s *Asteroid) generatePoints() {
	var points []float32
	var numV int = (rand.Intn(5) + 10)
	//For each vertex
	for i := 0; i < numV*3; i += 3 {
		//Calculate angles
		var angle float64 = (2 * math.Pi) / float64(numV*3) * float64(i)
		//Add coordinate with random radius
		points = append(points, float32((rand.Float64()-0.3)*30+float64(s.r)*math.Cos(angle))/Width)
		points = append(points, float32((rand.Float64()-0.3)*30+float64(s.r)*math.Sin(angle))/Height)
		points = append(points, 0)
	}
	s.points = points
	s.init = points
}

/*---------------------------------------------------------------------
|  Function	 	   CreateAsteroids
|
|  Purpose:  	   This a list of randomly generated asteroids. The
|				   asteroids are generated at random locations,
|				   random sizes, and velocities.
|
|  Pre-condition:  Nones
|
|  Post-condition: None
|
|  Parameters:
|		num		-- Number of asteroids to create
|		minSize -- Minimum radius of asteroids
|		maxSize -- Maximum radius of asteroids
|
|  Returns:  []Asteroids -- list of generated asteroids
*-------------------------------------------------------------------*/
func CreateAsteroids(num int, minSize float32, maxSize float32) []Asteroid {
	var list []Asteroid
	//For each asteroid
	for i := 0; i < num; i++ {
		var a Asteroid
		a.done = false
		//Random position
		a.x = rand.Float32()*2 - 1.0
		a.y = rand.Float32()*2 - 1.0
		//Check to make position isn't the center
		for (a.x > -0.2 && a.x < 0.2) || (a.y > -0.2 && a.x < 0.2) {
			a.x = rand.Float32()*2 - 1.0
			a.y = rand.Float32()*2 - 1.0
		}

		//Random velocity
		a.dX = (rand.Float32() - 0.5) / 1000
		a.dY = (rand.Float32() - 0.5) / 1000
		//Random rotation
		a.r = rand.Float32()*(maxSize-minSize) + minSize
		a.rot = rand.Float64() / 3
		a.generatePoints()
		a.points = TranslatePoints(a.points, a.x, a.y)
		list = append(list, a)
	}
	return list
}

/*---------------------------------------------------------------------
|  Method	 	   split
|
|  Purpose:  	   This generates 2 new asteroids of smaller proportion
|				   to the existing asteroid. The asteroid will have
|				   a similar position, with smaller size, and differing
|				   velocity
|
|  Pre-condition:  Nones
|
|  Post-condition: None
|
|  Parameters:     None
|
|  Returns:  []Asteroid -- list of generated asteroids
*-------------------------------------------------------------------*/
func (s *Asteroid) split() []Asteroid {
	var list []Asteroid
	//First asteroid
	var a Asteroid
	a.x = s.x
	a.y = s.y
	//Create faster velocity of opposite direction
	a.dX = s.dX * -4
	a.dY = s.dY * -4
	//Smaller radius
	a.r = s.r - 40
	a.rot = rand.Float64() / 3
	a.generatePoints()
	a.points = TranslatePoints(a.points, a.x, a.y)
	a.points = RotatePoints(a.points, 76, a.x, a.y)
	list = append(list, a)
	//Second asteroi
	var b Asteroid
	b.x = s.x
	b.y = s.y
	//Create faster velocity of opposite direction
	b.dX = s.dX * 4
	b.dY = s.dY * 4
	//Smaller radius
	b.r = s.r - 40
	b.rot = rand.Float64() / 3
	b.generatePoints()
	b.points = TranslatePoints(b.points, b.x, b.y)
	b.points = RotatePoints(b.points, 128, b.x, b.y)
	list = append(list, b)
	return list
}
