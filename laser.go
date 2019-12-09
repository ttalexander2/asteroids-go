package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	laserSpeed = 0.03
	laserLen   = 10
	laserThick = 4
)

type Laser struct {
	drawable uint32
	points   []float32
	window   *glfw.Window
	x        float32
	y        float32
	dX       float32
	dY       float32
	rot      float64
	speed    float64
	done     bool
}

func (l *Laser) Init(angle float64, x float32, y float32) {
	l.rot = angle
	l.x = x
	l.y = y
	l.getPoints()
	l.points = RotatePoints(l.points, l.rot, l.x, l.y)
	l.speed = laserSpeed
	l.done = false
	l.drawable = makeVertexArrayObj(l.points)
}

//Update
func (s *Laser) Update() {
	s.checkEdges()
	s.dX = float32(math.Cos(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.speed)
	s.dY = float32(math.Sin(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.speed)
	s.x += s.dX
	s.y += s.dY
	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)
}

func (s *Laser) Draw() {
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(s.points)/3))
}

func (s *Laser) checkEdges() {
	if s.x > 1 || s.x < -1 || s.y > 1 || s.y < -1 {
		s.done = true
	}
}

func (s *Laser) getPoints() {
	s.points = []float32{
		(s.x*Width - laserThick) / Width, (s.y * Height) / Height, 0, // Top
		(s.x*Width - laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y * Width) / Height, 0,
		(s.x*Width + laserThick) / Width, (s.y*Width + laserLen) / Height, 0,
		(s.x*Width - laserThick) / Width, (s.y * Width) / Height, 0,
	}
}

func (s *Laser) hits(a Asteroid) bool {
	var d float32
	d = float32(math.Sqrt((math.Pow(float64((s.x-a.x)*Width), 2.0) + math.Pow(float64((s.y-a.y)*Height), 2))))
	if d < a.r {
		return true
	}
	return false
}
