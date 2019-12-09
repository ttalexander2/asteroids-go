package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Asteroid struct {
	drawable uint32
	points   []float32
	init     []float32
	window   *glfw.Window
	x        float32
	y        float32
	dX       float32
	dY       float32
	r        float32
	dRot     float64
	rot      float64
	done     bool
}

//Update
func (s *Asteroid) Update() {
	s.checkEdges()

	//s.points = RotatePoints(s.points, s.rot, s.dX, s.dX)

	s.x += s.dX
	s.y += s.dY

	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)
}

//Draw
func (s *Asteroid) Draw() {
	s.drawable = makeVertexArrayObj(s.points)
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.POINTS, 0, int32(len(s.points)/3))
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(len(s.points)/3))
}

func (s *Asteroid) checkEdges() {
	if s.x > 1 {
		s.getPoints()
		s.points = TranslatePoints(s.points, -2, 0)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
		s.x = -1
	} else if s.x < -1 {
		s.getPoints()
		s.points = TranslatePoints(s.points, 2, 0)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
		s.x = 1
	}
	if s.y > 1 {
		s.y = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, 0, -2)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.y < -1 {
		s.y = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, 0, 2)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
}

func (s *Asteroid) getPoints() {
	s.points = s.init
	s.drawable = makeVertexArrayObj(s.points)
}

func (s *Asteroid) generatePoints() {
	var points []float32
	var numV int = (rand.Intn(5) + 10)
	for i := 0; i < numV*3; i += 3 {
		var angle float64 = (2 * math.Pi) / float64(numV*3) * float64(i)
		points = append(points, float32((rand.Float64()-0.3)*30+float64(s.r)*math.Cos(angle))/Width)
		points = append(points, float32((rand.Float64()-0.3)*30+float64(s.r)*math.Sin(angle))/Height)
		points = append(points, 0)
	}
	s.points = points
	s.init = points
}

func CreateAsteroids(num int, minSize float32, maxSize float32) []Asteroid {
	var list []Asteroid
	for i := 0; i < num; i++ {
		var a Asteroid
		a.done = false
		a.x = rand.Float32()*2 - 1.0
		a.y = rand.Float32()*2 - 1.0
		for (a.x > -0.2 && a.x < 0.2) || (a.y > -0.2 && a.x < 0.2) {
			a.x = rand.Float32()*2 - 1.0
			a.y = rand.Float32()*2 - 1.0
		}
		a.dX = (rand.Float32() - 0.5) / 1000
		a.dY = (rand.Float32() - 0.5) / 1000
		a.r = rand.Float32()*(maxSize-minSize) + minSize
		a.rot = rand.Float64() / 3
		a.generatePoints()
		a.points = TranslatePoints(a.points, a.x, a.y)
		list = append(list, a)
	}
	return list
}

func (s *Asteroid) split() []Asteroid {
	var list []Asteroid
	var a Asteroid
	a.x = s.x
	a.y = s.y

	a.dX = s.dX * -4
	a.dY = s.dY * -4
	a.r = s.r - 40
	a.rot = rand.Float64() / 3
	a.generatePoints()
	a.points = TranslatePoints(a.points, a.x, a.y)
	a.points = RotatePoints(a.points, 76, a.x, a.y)
	list = append(list, a)

	var b Asteroid
	b.x = s.x
	b.y = s.y

	b.dX = s.dX * 4
	b.dY = s.dY * 4
	b.r = s.r - 40
	b.rot = rand.Float64() / 3
	b.generatePoints()
	b.points = TranslatePoints(b.points, b.x, b.y)
	b.points = RotatePoints(b.points, 128, b.x, b.y)
	list = append(list, b)
	return list
}
