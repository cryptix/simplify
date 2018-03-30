package simplify

import (
	"github.com/cryptix/stl"
	"github.com/go-gl/mathgl/mgl64"
)

func NewTriangle(v1, v2, v3 mgl64.Vec3) stl.Triangle {
	t := stl.Triangle{}
	t.Vertices = [3]mgl64.Vec3{v1, v2, v3}
	return t
}

/*
func (t *Triangle) Quadric() mgl64.Mat4 {
	n := t.Normal()
	x, y, z := t.V1.X(), t.V1.Y(), t.V1.Z()
	a, b, c := n.X(), n.Y(), n.Z()
	d := -a*x - b*y - c*z
	return mgl64.Mat4{
		a * a, a * b, a * c, a * d,
		a * b, b * b, b * c, b * d,
		a * c, b * c, c * c, c * d,
		a * d, b * d, c * d, d * d,
	}
}

func (t *Triangle) Normal() mgl64.Vec3 {
	e1 := t.V2.Sub(t.V1)
	e2 := t.V3.Sub(t.V1)
	return e1.Cross(e2).Normalize()
}
*/
