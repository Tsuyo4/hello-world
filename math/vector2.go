package math

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) GetWidth() float64 {
	return v.X
}

func (v *Vector2) SetWidth( width float64 ) {
	v.X = width
}

func (v *Vector2) GetHeight() float64 {
	return v.Y
}

func (v *Vector2) SetHeight( height float64 ) {
	v.Y = height
}

func (v *Vector2) Set( x, y float64 ) *Vector2 {
	v.X = x
	v.Y = y

	return v
}

func (v *Vector2) SetScalar( scalar float64 ) {
	v.X = scalar
	v.Y = scalar
}

func (v *Vector2) SetX( x float64 ) *Vector2 {
	v.X = x

	return v
}

func (v *Vector2) SetY( y float64 ) *Vector2 {
	v.Y = y

	return v
}

func (v *Vector2) SetComponent( index int, value float64 ) {
	switch index {
	case 0:
		v.X = value;
	case 1:
		v.Y = value
	}
}

func (v *Vector2) GetComponent( index int ) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	}

	return 0.0
}

func (v *Vector2) Clone() *Vector2 {
	v2 := &Vector2{ X:v.X, Y:v.Y }
	return v2
}

func (v *Vector2) Copy( other *Vector2) *Vector2 {
	v.X = other.X
	v.Y = other.Y

	return v
}

func (v *Vector2) Add( other *Vector2 ) *Vector2{
	v.X += other.X;
	v.Y += other.Y;

	return v
}

func (v *Vector2) AddScalar( scalar float64) {
	v.X += scalar
	v.Y += scalar
}

func (v *Vector2) AddVectors( a *Vector2, b *Vector2 ) *Vector2 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y

	return v
}

func (v *Vector2) AddScaledVector( other Vector2, scalar float64 ) *Vector2 {
	v.X += other.X * scalar
	v.Y += other.Y * scalar

	return v
}


func (v *Vector2) Sub( other *Vector2 ) *Vector2 {
	v.X -= other.X
	v.Y -= other.Y

	return v
}

func (v *Vector2) SubScalar( scalar float64 ) *Vector2 {
	v.X -= scalar
	v.Y -= scalar

	return v
}

func (v *Vector2) SubVectors( a *Vector2, b *Vector2) *Vector2 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y

	return v
}

func (v *Vector2) Multiply( other Vector2 ) *Vector2 {
	v.X *= other.X
	v.Y *= other.Y

	return v
}

func (v *Vector2) MultiplyScalar( scalar float64 ) *Vector2 {
	v.X *= scalar
	v.Y *= scalar

	return v
}

func (v *Vector2) Divide( other Vector2 ) *Vector2 {
	v.X /= other.X
	v.Y /= other.Y

	return v
}

func (v *Vector2) DivideScalar( scalar float64) *Vector2 {
	v.X /= scalar
	v.Y /= scalar

	return v
}

func (v *Vector2) Min( other *Vector2 ) *Vector2 {
	v.X = math.Min( v.X, other.X )
	v.Y = math.Min( v.Y, other.Y )

	return v
}

func (v *Vector2) Max( other *Vector2) *Vector2 {
	v.X = math.Max( v.X, other.X )
	v.Y = math.Max( v.Y, other.Y )

	return v
}

func (v *Vector2) Clamp( min, max *Vector2) *Vector2 {
	v.X = math.Max( min.X, math.Min( max.X, v.X ))
	v.Y = math.Max( min.Y, math.Min( max.Y, v.Y ))

	return v
}

func (v *Vector2) ClampScalar( minScalar, maxScalar float64) *Vector2 {
	min := &Vector2{ X:minScalar, Y:minScalar }
	max := &Vector2{ X:maxScalar, Y:maxScalar }

	return v.Clamp(min, max)
}

func (v *Vector2) ClampLength( min float64, max float64 ) *Vector2{
	length := v.Length()

	return v.MultiplyScalar( math.Max( min, math.Min( max, length)) / length )
}


func (v *Vector2) Floor() *Vector2 {
	v.X = math.Floor( v.X )
	v.Y = math.Floor( v.Y )

	return v
}

func (v *Vector2) Ceil() *Vector2 {
	v.X = math.Ceil( v.X )
	v.Y = math.Ceil( v.Y )

	return v
}

func (v *Vector2) Round() *Vector2 {
	v.X = math.Floor( v.X + .5 )
	v.Y = math.Floor( v.Y + .5 )

	return v
}

func (v *Vector2) RoundToZero() *Vector2 {
	if v.X < 0 {
		v.X = math.Ceil( v.X )
	} else {
		v.X = math.Floor( v.X )
	}

	if v.Y < 0 {
		v.Y = math.Ceil( v.Y )
	} else {
		v.Y = math.Floor( v.Y )
	}

	return v
}

func (v *Vector2) Negate() *Vector2 {
	v.X = -v.X
	v.Y = -v.Y

	return v
}

func (v *Vector2) Dot( other *Vector2) float64 {
	return v.X * other.X + v.Y * other.Y
}

func (v *Vector2) Length() float64 {
	return math.Sqrt( v.X * v.X + v.Y * v.Y )
}

func (v *Vector2) LengthSq() float64 {
	return v.X * v.X + v.Y * v.Y
}

func (v *Vector2) Normalize() *Vector2 {
	return v.DivideScalar( v.Length() )
}

func (v *Vector2) Angle() float64 {
	angle := math.Atan2( v.Y, v.X )

	if angle < 0 {
		angle += 2 * math.Pi
	}

	return angle
}

func (v *Vector2) DistanceTo( other *Vector2 ) float64 {
	return math.Sqrt( v.DistanceToSquared(other) )
}

func (v *Vector2) DistanceToSquared( other *Vector2 ) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return dx * dx + dy * dy
}

func (v *Vector2) SetLength( length float64 ) *Vector2{
	if v.Length() == 0 {
		return v
	}
	orgLength := v.Length()
	
	v.MultiplyScalar( length )
	v.DivideScalar( orgLength )
	
	return v
}

func (v *Vector2) Lerp( other *Vector2, alpha float64 ) *Vector2 {
	v.X += ( other.X - v.X ) * alpha
	v.Y += ( other.Y - v.Y ) * alpha

	return v
}

func (v *Vector2) Equals( other *Vector2) bool {
	return (( v.X == other.X ) && ( v.Y == other.Y ))
}

func (v *Vector2) FromArray( array []float64  ) {
	v.X = array[0]
	v.Y = array[1]
}

func (v *Vector2) ToArray( array []float64 ) []float64 {
	array[0] = v.X
	array[1] = v.Y

	return array
}

func (v *Vector2) RotateAround( center Vector2, angle float64 ) *Vector2{
	c := math.Cos( angle )
	s := math.Sin( angle )

	x := v.X - center.X
	y := v.Y - center.Y

	v.X = x * c - y * s + center.X
	v.Y = y * s + y * c + center.Y

	return v
}
