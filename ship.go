/**************************************************************************************************
|   Assignment:  Final Project Part 3:  ship.go
|      Authors:  Thomas Alexander & Cameron Larson
				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
|       Grader:  Tito Ferra & Josh Xiong
|
|       Course:  CSc 372
|   Instructor:  L. McCann
|     Due Date:  12-9-19 3:30pm
|
|  Description:  This file holds the structure for the player's ship. The file is responsible for
|				 updating and rendering the player object. Additionally, the ship listens for user
|				 input to drive the ship. Depending on input, the ship can create lasers and check
|				 for collisions with asteroids.
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
	speed     = 0.005 //Ship Speed/Acceleration
	rotspeed  = 0.8   //Ship Rotation
	maxspeed  = 0.009 //Max ship speed
	maxlasers = 1     //Max lasers to be fired at once
)

type Ship struct {
	drawable  uint32       //Vertex buffer array
	points    []float32    //Vertex point array
	window    *glfw.Window //Window to render to
	x         float32      //x position on the screen
	y         float32      //y position on the screen
	dX        float32      //x velocity
	dY        float32      //y velocity
	r         float32      //radius of ship
	dRot      float64      //change in rotation
	rot       float64      //rotation of ship
	force     float64      //Force/Thrust
	lasers    []Laser      //Lasers fired by ship
	asteroids []Asteroid   //Reference to asteroids
	firing    bool         //Fire button is being held
	lcount    int          //Number of lasers fired in a single press
	gameOver  bool         //Whether the game is over
	dying     bool         //Handling whether the player is in death animation
}

/*---------------------------------------------------------------------
|  method	 	   Update
|
|  Purpose:  	   This method updates the positions and velocities
|				   of the ship, checks for collisions between lasers
|				   and asteroids, and the ship. Additionally, this
|			       method checks for user input and updates ship
|				   values accordingly.
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   Ship and asteroids must be initialized.
|
|  Post-condition: Ship values are updated, input is handled
|
|  Parameters:	   None
|
|  Returns:  	   None
*-------------------------------------------------------------------*/
func (s *Ship) Update() {
	//Handle Death Animation
	if s.dying {
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)

		s.rot += s.dRot
		s.r = s.r * 0.95 //Decrease ship size
		s.dX = 0
		s.dY = 0
		s.drawable = makeVertexArrayObj(s.points)
		s.points = RotatePoints(s.points, 100, s.x, s.y)
		if s.r < 0.05 {
			s.gameOver = true
		}
		return
	}

	//Handle Input
	if s.window.GetKey(glfw.KeyUp) == glfw.Press {
		s.force += speed
	}
	if s.window.GetKey(glfw.KeyLeft) == glfw.Press {
		s.dRot += rotspeed
	}
	if s.window.GetKey(glfw.KeyRight) == glfw.Press {
		s.dRot += -rotspeed
	}
	if s.window.GetKey(glfw.KeySpace) == glfw.Press {
		//Create laser object
		if !s.firing {
			var l Laser
			l.Init(s.rot, s.x, s.y)
			//Add to list
			s.lasers = append(s.lasers, l)
			s.lcount++
			if s.lcount > maxlasers {
				s.firing = true
			}
		}
	}
	//Check for key release to allow player to fire again
	if s.window.GetKey(glfw.KeySpace) == glfw.Release {
		s.firing = false
		s.lcount = 0
	}
	//Check for window bounds
	s.checkEdges()

	//Rotate ship
	s.points = RotatePoints(s.points, s.dRot, s.x, s.y)

	//Calculate Thrust Direction
	s.dX += float32(math.Cos(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.force)
	s.dY += float32(math.Sin(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.force)
	s.dX = float32(Clamp(s.dX, -maxspeed, maxspeed))
	s.dY = float32(Clamp(s.dY, -maxspeed, maxspeed))

	//Update position with velocity
	s.x += s.dX
	s.y += s.dY
	s.rot += s.dRot

	//Translate vertex array, generate VAO
	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)

	//Decrease velocity and rotation
	s.dX *= .985
	s.dY *= .985
	s.dRot *= 0.85
	s.force = 0

	//Game over
	if len(s.asteroids) == 0 {
		s.gameOver = true
	}

	//Check for ship collision
	for i := 0; i < len(s.asteroids); i++ {
		if s.hits(s.asteroids[i]) {
			s.dying = true
		}
	}

	//Check for laser-asteroid collision
	for i := 0; i < len(s.lasers); i++ {
		s.lasers[i].Update()
		if s.lasers[i].done {
			//Laser off screen, remove from list
			s.lasers = append(s.lasers[:i], s.lasers[i+1:]...)
			continue
		}
		for j := len(s.asteroids) - 1; j >= 0; j-- {
			if s.lasers[i].hits(s.asteroids[j]) {
				//laser hit asteroid
				if s.asteroids[j].r > 80 {
					//Split asteroid into smaller asteroids
					var a []Asteroid = s.asteroids[j].split()
					s.asteroids = append(s.asteroids, a...)
				}
				//Remove asteroid and laser
				s.asteroids = append(s.asteroids[:j], s.asteroids[j+1:]...)
				s.lasers = append(s.lasers[:i], s.lasers[i+1:]...)
				break
			}
		}
	}

}

/*---------------------------------------------------------------------
|  method	 	   Draw
|
|  Purpose:  	   This method renders the ship on the screen
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the ship must have a VAO (vertex array object)
|
|  Post-condition: Laser fields are ready to be rendered
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Ship) Draw() {
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.POINTS, 0, int32(len(s.points)/3))
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(len(s.points)/3))
	//Draw the ship's lasers
	for i := 0; i < len(s.lasers); i++ {
		s.lasers[i].Draw()
	}
}

/*---------------------------------------------------------------------
|  Method	 	   getPoints
|
|  Purpose:  	   This Method generates an array of points for the ship
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
func (s *Ship) getPoints() {
	s.points = []float32{
		(0) / Width, (0 + s.r + 15) / Height, 0, // Top
		(0 - s.r) / Width, (0 - s.r) / Height, 0, // Left
		(0 + s.r) / Width, (0 - s.r) / Height, 0, // Right
	}
}

/*---------------------------------------------------------------------
|  Method	 	   checkEdges
|
|  Purpose:  	   This Method checks whether the ship has traveled past
|				   the bounds of the screen. If it has, it changes the
|				   location to the opposite side
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized,
|				   the ship must have a VAO (vertex array object)
|
|  Post-condition: Ship's location is now on the other side of the screen
|
|  Parameters:		None
|
|  Returns:  	    None
*-------------------------------------------------------------------*/
func (s *Ship) checkEdges() {
	//Right bounds
	if s.x > 1 {
		s.x = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.x < -1 { //Left bounds
		s.x = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
	//Upper bounds
	if s.y > 1 {
		s.y = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.y < -1 { //Lower bounds
		s.y = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
}

/*---------------------------------------------------------------------
|  Method	 	   hits
|
|  Purpose:  	   This Method checks whether the ship has collided with
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
func (s *Ship) hits(a Asteroid) bool {
	var d float32
	//Calculate distance between the two objects
	d = float32(math.Sqrt((math.Pow(float64((s.x-a.x)*Width), 2.0) + math.Pow(float64((s.y-a.y)*Height), 2))))
	//if the distance is less than the sum of their radii, a collision has occured
	if d < s.r+a.r {
		return true
	}
	return false
}
