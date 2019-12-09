package main

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	speed     = 0.005
	rotspeed  = 0.8
	maxspeed  = 0.009
	maxlasers = 1
)

type Ship struct {
	drawable  uint32
	points    []float32
	window    *glfw.Window
	x         float32
	y         float32
	dX        float32
	dY        float32
	r         float32
	dRot      float64
	rot       float64
	force     float64
	lasers    []Laser
	asteroids []Asteroid
	firing    bool
	lcount    int
	gameOver  bool
	dying     bool
}

//Update
func (s *Ship) Update() {
	if s.dying {
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)

		s.rot += s.dRot
		s.r = s.r * 0.95
		s.dX = 0
		s.dY = 0
		s.drawable = makeVertexArrayObj(s.points)
		s.points = RotatePoints(s.points, 100, s.x, s.y)
		if s.r < 0.05 {
			s.gameOver = true
		}
		return
	}

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
		if !s.firing {
			var l Laser
			l.Init(s.rot, s.x, s.y)
			s.lasers = append(s.lasers, l)
			s.lcount++
			if s.lcount > maxlasers {
				s.firing = true
			}
		}
	}
	if s.window.GetKey(glfw.KeySpace) == glfw.Release {
		s.firing = false
		s.lcount = 0
	}
	s.checkEdges()

	s.points = RotatePoints(s.points, s.dRot, s.x, s.y)

	s.dX += float32(math.Cos(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.force)
	s.dY += float32(math.Sin(s.rot*(math.Pi/180)+90*(math.Pi/180)) * s.force)
	s.dX = float32(Clamp(s.dX, -maxspeed, maxspeed))
	s.dY = float32(Clamp(s.dY, -maxspeed, maxspeed))
	s.x += s.dX
	s.y += s.dY
	s.rot += s.dRot

	s.points = TranslatePoints(s.points, s.dX, s.dY)
	s.drawable = makeVertexArrayObj(s.points)
	s.dX *= .985
	s.dY *= .985
	s.dRot *= 0.85
	s.force = 0

	if len(s.asteroids) == 0 {
		s.gameOver = true
	}

	for i := 0; i < len(s.asteroids); i++ {
		if s.hits(s.asteroids[i]) {
			s.dying = true
		}
	}

	for i := 0; i < len(s.lasers); i++ {
		s.lasers[i].Update()
		if s.lasers[i].done {
			s.lasers = append(s.lasers[:i], s.lasers[i+1:]...)
			continue
		}
		for j := len(s.asteroids) - 1; j >= 0; j-- {
			if s.lasers[i].hits(s.asteroids[j]) {
				if s.asteroids[j].r > 80 {
					var a []Asteroid = s.asteroids[j].split()
					s.asteroids = append(s.asteroids, a...)
				}
				s.asteroids = append(s.asteroids[:j], s.asteroids[j+1:]...)
				s.lasers = append(s.lasers[:i], s.lasers[i+1:]...)
				break
			}
		}
	}

}

func (s *Ship) Draw() {
	gl.BindVertexArray(s.drawable)
	gl.DrawArrays(gl.POINTS, 0, int32(len(s.points)/3))
	gl.DrawArrays(gl.LINE_LOOP, 0, int32(len(s.points)/3))
	for i := 0; i < len(s.lasers); i++ {
		s.lasers[i].Draw()
	}
}

func (s *Ship) getPoints() {
	s.points = []float32{
		(0) / Width, (0 + s.r + 15) / Height, 0, // Top
		(0 - s.r) / Width, (0 - s.r) / Height, 0, // Left
		(0 + s.r) / Width, (0 - s.r) / Height, 0, // Right
	}
}

func (s *Ship) checkEdges() {
	if s.x > 1 {
		s.x = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.x < -1 {
		s.x = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
	if s.y > 1 {
		s.y = -1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	} else if s.y < -1 {
		s.y = 1
		s.getPoints()
		s.points = TranslatePoints(s.points, s.x, s.y)
		s.points = RotatePoints(s.points, s.rot, s.x, s.y)
	}
}

func (s *Ship) hits(a Asteroid) bool {
	var d float32
	d = float32(math.Sqrt((math.Pow(float64((s.x-a.x)*Width), 2.0) + math.Pow(float64((s.y-a.y)*Height), 2))))
	if d < s.r+a.r {
		return true
	}
	return false
}
