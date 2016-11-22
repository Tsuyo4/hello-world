package math

//import "math"

const (
	EulerRotationOrderXYZ int = iota+1
	EulerRotationOrderYZX
	EulerRotationOrderZXY
	EulerRotationOrderXZY
	EulerRotationOrderYXZ
	EulerRotationOrderZYX

	EulerDefalutOrder = EulerRotationOrderXYZ
)


type Euler struct {
	x float64
	y float64
	z float64

	order int
	callback func()
}

func NewEuler() *Euler {
	e := &Euler{}

	e.order = EulerDefalutOrder

	return e
}

func (e *Euler) GetX() float64 {
	return e.x
}

func (e *Euler) SetX( x float64 ) {
	e.x = x
	e.OnChangeCallback()
}

func (e *Euler) GetY() float64 {
	return e.y
}

func (e *Euler) SetY( y float64 ) {
	e.y = y
	e.OnChangeCallback()
}

func (e *Euler) GetZ() float64 {
	return e.z
}

func (e *Euler) SetZ( z float64 ) {
	e.z = z
	e.OnChangeCallback()
}

func (e *Euler) GetOrder() int {
	return e.order
}

func (e *Euler) SetOrder( order int ) {
	e.order = order
}

func (e *Euler) Set( x, y, z float64, order int ) *Euler {
	e.x = x
	e.y = y
	e.z = z
	e.order = order

	e.OnChangeCallback()

	return e
}

func (e *Euler) Clone() *Euler {
	e1 := &Euler{
		x:e.x,
		y:e.y,
		z:e.z,
		order:e.order,
	}

	return e1
}

func (e *Euler) Copy( euler *Euler ) *Euler {
	e.x = euler.x
	e.y = euler.y
	e.z = euler.z
	e.order = euler.order

	e.OnChangeCallback()

	return e
}

func (e *Euler) SetFromRotationMatrix() {
	// TODO After Matrix
}

func (e *Euler) SetFromQuaternion() {
	// TODO After Quaternion
}

func (e *Euler) SetFromVector3( v *Vector3) *Euler {
	return e.SetFromVector3Order( v, e.order )
}

func (e *Euler) SetFromVector3Order( v *Vector3, order int ) *Euler {
	return e.Set( v.X, v.Y, v.Z, order )
}

func (e *Euler) Reorder() {
	// TODO After Quaternion
}

func (e *Euler) Equals( euler *Euler) bool {
	return ( euler.x == e.x ) && ( euler.y == e.y ) && ( euler.z == e.z ) && ( euler.order == e.order )
}

func (e *Euler) FromArray( array []float64 ) *Euler {
	e.x = array[0]
	e.y = array[1]
	e.z = array[2]
	if len(array) >= 4 {
		e.order = int(array[3])
 	}

	e.OnChangeCallback()

	return e
}

func (e *Euler) OnChange( callback func() ) *Euler {
	e.callback = callback

	return e
}

func (e *Euler) OnChangeCallback() {
	if e.callback != nil {
		e.callback()
	}
}
