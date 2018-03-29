package simplify

import "github.com/go-gl/mathgl/mgl64"

type Vertex struct {
	mgl64.Vec3
	Quadric mgl64.Mat4
}

func NewVertex(v mgl64.Vec3) *Vertex {
	return &Vertex{Vec3: v}
}
