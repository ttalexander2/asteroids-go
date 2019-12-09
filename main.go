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
	Width    = 1000
	Height   = 1000
	rows     = 10
	columns  = 10
	FpsCap   = 60
	shipSize = 30
)

func main() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UnixNano())

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	var asteroids []Asteroid
	asteroids = CreateAsteroids(rand.Intn(10)+8, 80, 160)

	var player Ship
	player.window = window
	player.r = shipSize
	player.firing = false
	player.asteroids = asteroids
	player.getPoints()

	lastFrame := time.Now()
	diff := time.Now().Sub(lastFrame)
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

func render(window *glfw.Window, program uint32, s Ship, a []Asteroid) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	s.Draw()
	for i := 0; i < len(a); i++ {
		a[i].Draw()
	}
	glfw.PollEvents()
	window.SwapBuffers()
}

func reload(window *glfw.Window) Ship {
	var asteroids []Asteroid
	asteroids = CreateAsteroids(rand.Intn(10)+8, 80, 160)

	var player Ship
	player.window = window
	player.r = shipSize
	player.firing = false
	player.asteroids = asteroids
	player.getPoints()
	return player
}

// initGlfw initializes glfw and returns a Window to use.
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

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	check(gl.Init())

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShader, gl.VERTEX_SHADER)
	check(err)

	fragmentShaderWhite, err := compileShader(fragmentShaderWhite, gl.FRAGMENT_SHADER)
	check(err)

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShaderWhite)
	gl.LinkProgram(prog)
	return prog
}

// makeVertexArrayObj initializes and returns a vertex array from the points provided.
func makeVertexArrayObj(points []float32) uint32 {
	var vertexBufferObj uint32
	gl.GenBuffers(1, &vertexBufferObj)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObj)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vertexArrayObj uint32
	gl.GenVertexArrays(1, &vertexArrayObj)
	gl.BindVertexArray(vertexArrayObj)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObj)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vertexArrayObj
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
