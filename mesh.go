package simplify

import "github.com/cryptix/stl"

type Mesh struct {
	Triangles []stl.Triangle
}

func NewMesh(triangles []stl.Triangle) *Mesh {
	return &Mesh{triangles}
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}

func (m *Mesh) Simplify(factor float64) *Mesh {
	return Simplify(m, factor)
}
