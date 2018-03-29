package simplify

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

func Less(a, b mgl64.Vec3) bool {
	if a.X() != b.X() {
		return a.X() < b.X()
	}
	if a.Y() != b.Y() {
		return a.Y() < b.Y()
	}
	return a.Z() < b.Z()
}

type PairKey struct {
	A, B mgl64.Vec3
}

func MakePairKey(a, b *Vertex) PairKey {
	if Less(a.Vec3, b.Vec3) {
		a, b = b, a
	}
	return PairKey{a.Vec3, b.Vec3}
}

type Pair struct {
	A, B        *Vertex
	Index       int
	Removed     bool
	CachedError float64
}

func NewPair(a, b *Vertex) *Pair {
	if Less(b.Vec3, a.Vec3) {
		a, b = b, a
	}
	return &Pair{a, b, -1, false, -1}
}

func (p *Pair) Quadric() mgl64.Mat4 {
	return p.A.Quadric.Add(p.B.Quadric)
}

func (p *Pair) Vector() mgl64.Vec3 {
	q := p.Quadric()
	if math.Abs(q.Det()) > 1e-3 {
		v := QuadricVector(q)
		if !math.IsNaN(v.X()) && !math.IsNaN(v.Y()) && !math.IsNaN(v.Z()) {
			return v
		}
	}
	// cannot compute best vector with matrix
	// look for best vector along edge
	const n = 32
	a := p.A.Vec3
	b := p.B.Vec3
	d := b.Sub(a)
	bestE := -1.0
	bestV := mgl64.Vec3{}
	for i := 0; i <= n; i++ {
		t := float64(i) / n
		v := a.Add(d.Mul(t))
		e := QuadricError(q, v)
		if bestE < 0 || e < bestE {
			bestE = e
			bestV = v
		}
	}
	return bestV
}

func (p *Pair) Error() float64 {
	if p.CachedError < 0 {
		p.CachedError = QuadricError(p.Quadric(), p.Vector())
	}
	return p.CachedError
}
