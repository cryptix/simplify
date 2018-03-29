package simplify

import "github.com/go-gl/mathgl/mgl64"

func QuadricError(a mgl64.Mat4, v mgl64.Vec3) float64 {

	return (v.X()*a.At(0, 0)*v.X() + v.Y()*a.At(1, 0)*v.X() + v.Z()*a.At(2, 0)*v.X() + a.At(3, 0)*v.X() +
		v.X()*a.At(0, 1)*v.Y() + v.Y()*a.At(1, 1)*v.Y() + v.Z()*a.At(2, 1)*v.Y() + a.At(3, 1)*v.Y() +
		v.X()*a.At(0, 2)*v.Z() + v.Y()*a.At(1, 2)*v.Z() + v.Z()*a.At(2, 2)*v.Z() + a.At(3, 2)*v.Z() +
		v.X()*a.At(0, 3) + v.Y()*a.At(1, 3) + v.Z()*a.At(2, 3) + a.At(3, 3))
}

func QuadricVector(a mgl64.Mat4) mgl64.Vec3 {
	b := mgl64.Mat4{
		a.At(0, 0), a.At(0, 1), a.At(0, 2), a.At(0, 3),
		a.At(1, 0), a.At(1, 1), a.At(1, 2), a.At(1, 3),
		a.At(2, 0), a.At(2, 1), a.At(2, 2), a.At(2, 3),
		0, 0, 0, 1,
	}
	return b.Inv().Mul4x1(mgl64.Vec4{0, 0, 0, 1}).Vec3()
}
