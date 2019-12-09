/**************************************************************************************************
||   Assignment:  Final Project Part 3:  main.go
||      Authors:  Thomas Alexander & Cameron Larson
||				 (ttalexander2@email.arizona.edu, cblarson@email.arizona.edu)
||       Grader:  Tito Ferra & Josh Xiong
||
||       Course:  CSc 372
||   Instructor:  L. McCann
||     Due Date:  12-9-19 3:30pm
||
||  Description:  This is the main driver of our space game, based on the classic game Asteroids.
||				 This program is mainly responsible for the initialization of the window, seting up
||			     the OpenGL graphics context, and
||
||     Language:  GoLang
|| Ex. Packages:  OpenGl, GLFW
||				 github.com/go-gl/gl/v4.1-core/gl
||				 github.com/go-gl/glfw/v3.2/glfw
||
|| Deficiencies:  I know of no unsatisfied requirements and no logic errors.
**************************************************************************************************/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	Width    = 1000 //Width of the window
	Height   = 1000 //Height of the window
	FpsCap   = 60
	shipSize = 30 //Radius of the ship
)

func main() {
	runtime.LockOSThread() //Lock thread (neccesary for OpenGl)

	rand.Seed(time.Now().UnixNano()) //Initialize random

	//Create graphics context
	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	//initialize asteroids
	var asteroids []Asteroid
	asteroids = CreateAsteroids(rand.Intn(10)+8, 80, 160)

	//Initialize player
	var player Ship
	player.window = window
	player.r = shipSize
	player.firing = false
	player.asteroids = asteroids
	player.getPoints()

	lastFrame := time.Now()
	diff := time.Now().Sub(lastFrame)
	//main game loop
	for !window.ShouldClose() {
		if (window.GetKey(glfw.KeyEscape)) == glfw.Press {
			window.SetShouldClose(true)
		}
		if player.gameOver {
			player = reload(window)
		}

		player.Update()
		for i := 0; i < len(player.asteroids); i++ {
			player.asteroids[i].Update()
		}

		render(window, program, player, player.asteroids)

		// cap framerate
		diff = time.Now().Sub(lastFrame)
		lastFrame = time.Now()
		if diff < 1/FpsCap {
			time.Sleep(1/60 - diff)
		}

	}
}

/*---------------------------------------------------------------------
|  Function 	   render
|
|  Purpose:  	   This function renders the ship, lasers, and asteroids
|			 	   to the window
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: Objects are rendered on the screen
|
|  Parameters:
|      window 	-- The glfw.Window object to draw to
|	   program 	-- Address of the current GL program
|	   s		-- Ship struct to render
|	   a		-- list of asteroids to render
|
|  Returns:  None :)
*-------------------------------------------------------------------*/
func render(window *glfw.Window, program uint32, s Ship, a []Asteroid) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) //Clear the screen
	gl.UseProgram(program)
	s.Draw() //Draw the ship
	for i := 0; i < len(a); i++ {
		a[i].Draw() // Draw the asteroid
	}
	//Swap buffers and poll events
	glfw.PollEvents()
	window.SwapBuffers()
}

/*---------------------------------------------------------------------
|  Function 	   reload
|
|  Purpose:  	   This function creates a new ship and randomly generates new
			 	   asteroids, for the purpose of starting a new game
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: The Ship struct is initialized and holds references
|				   to the asteroids.
|
|  Parameters:
|      window 	-- Pointer to the glfw.Window so the Ship knows where
|				   to render
|
|  Returns:		   Ship struct for the player
*-------------------------------------------------------------------*/
func reload(window *glfw.Window) Ship {
	//Create randomly generated asteroids
	var asteroids []Asteroid
	asteroids = CreateAsteroids(rand.Intn(10)+8, 80, 160)

	//Create player's Ship
	var player Ship
	player.window = window
	player.r = shipSize
	player.firing = false
	player.asteroids = asteroids
	//Initialize vertex array (points)
	player.getPoints()
	return player
}

/*---------------------------------------------------------------------
|  Function 	   initGlfw
|
|  Purpose:  	   This function initializes a glfw window, sets
|				   version and window parameters, and returns the
|				   initialized window
|
|  Pre-condition:  OS thread must be locked
|
|  Post-condition: Window is initialized
|
|  Parameters:	   None
|
|  Returns:  	   pointer to glfw.Window (window)
*-------------------------------------------------------------------*/
func initGlfw() *glfw.Window {
	check(glfw.Init())
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(Width, Height, "Super Awesome Space Game!", nil, nil)
	check(err)

	window.MakeContextCurrent()

	return window
}

/*---------------------------------------------------------------------
|  Function 	   initOpenGL
|
|  Purpose:  	   This function initializes open gl, and compiles the
|				   shaders written in OpenGL Shader Language
|
|  Pre-condition:  GLFW window must be initialized
|
|  Post-condition: The OpenGl graphics context is setup and ready to use
|
|  Parameters: 	   None
|
|  Returns:		   uint32 for the OpenGl program
*-------------------------------------------------------------------*/
func initOpenGL() uint32 {
	//Initialize OpenGl
	check(gl.Init())

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	//Compile Shaders
	vertexShader, err := compileShader(vertexShader, gl.VERTEX_SHADER)
	check(err)

	fragmentShaderWhite, err := compileShader(fragmentShaderWhite, gl.FRAGMENT_SHADER)
	check(err)

	//Create and link program, and attach shaders
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShaderWhite)
	gl.LinkProgram(prog)
	return prog
}

/*---------------------------------------------------------------------
|  Function 	   makeVertexArrayObj
|
|  Purpose:  	   This function generates a vertex buffer object given
|				   an array of points. This function is necessary for
|				   OpenGl to render an array of vertices
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: None
|
|  Parameters:
|      points 	-- Array of floats representing vertices. Floats must be
|				   in multiples of three, stored in the following format:
				   [x1,y1,z1,x2,y2,z2,...]
|
|  Returns:		   uint32 for vertex buffer object, ready to be rendered
|				   with OpenGL
*-------------------------------------------------------------------*/
func makeVertexArrayObj(points []float32) uint32 {
	//Generate buffer data
	var vertexBufferObj uint32
	gl.GenBuffers(1, &vertexBufferObj)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObj)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	//Convert ot vertex data
	var vertexArrayObj uint32
	gl.GenVertexArrays(1, &vertexArrayObj)
	gl.BindVertexArray(vertexArrayObj)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObj)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vertexArrayObj
}

/*---------------------------------------------------------------------
|  Function 	   compileShader
|
|  Purpose:  	   This function compiles GLSL shaders (stored in shaders.go)
|				   shaders usable to OpenGL
|
|  Pre-condition:  OpenGl context and GLFW window must be initialized
|
|  Post-condition: None
|
|  Parameters:
|      string 	-- string object holding the GLSL code
|    shaderType -- uint32 holding OpenGl shader type (vertex/fragment )
|
|  Returns:		-- uint32 for the compiled shader to be attached to OpenGL
|				-- err error returned if shader fails to compile
*-------------------------------------------------------------------*/
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	//Compile shader held in string constant
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	//Check the status of compilation
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

/*---------------------------------------------------------------------
|  Function 	   check
|
|  Purpose:  	   Check if error is thrown
|
|  Pre-condition:  None
|
|  Post-condition: None
|
|  Parameters:
|      	err 	-- error object returned by function that could throw error
|
|  Returns:		None :)
*-------------------------------------------------------------------*/
func check(err error) {
	if err != nil {
		//Something went wrong
		panic(err)
	}
}
