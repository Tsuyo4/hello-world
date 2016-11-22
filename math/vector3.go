package math

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector3) Set( x, y, z float64) *Vector3{
	v.X = x
	v.Y = y
	v.Z = z

	return v
}

func (v *Vector3) SetScalar( scalar float64 ) *Vector3 {
	v.X = scalar
	v.Y = scalar
	v.Z = scalar

	return v
}

func (v *Vector3) SetX( x float64 ) *Vector3 {
	v.X = x

	return v
}

func (v *Vector3) SetY( y float64 ) *Vector3 {
	v.Y = y

	return v
}

func (v *Vector3) SetZ( z float64 ) *Vector3 {
	v.Z = z

	return v
}

func (v *Vector3) SetComponent( index int, value float64 ) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	}
}

func (v *Vector3) GetComponent( index int ) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	}

	return 0.0
}

func (v *Vector3) Clone() *Vector3 {
	v2 := &Vector3{ X:v.X, Y:v.Y, Z:v.Z }
	return v2
}

func (v *Vector3) Copy( other *Vector3 ) *Vector3 {
	v.X = other.X
	v.Y = other.Y
	v.Z = other.Z

	return v
}

func (v *Vector3) Add( other *Vector3) *Vector3 {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z

	return v
}

func (v *Vector3) AddScalar( scalar float64 ) *Vector3{
	v.X += scalar
	v.Y += scalar
	v.Z += scalar

	return v
}

func (v *Vector3) AddVectors( a, b *Vector3 ) *Vector3 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z

	return v
}

func (v *Vector3) AddScaledVector( other *Vector3, scalar float64 ) *Vector3 {
	v.X += v.X * scalar
	v.Y += v.Y * scalar
	v.Z += v.Z * scalar

	return v
}

func (v *Vector3) Sub( other *Vector3) *Vector3 {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z

	return v
}

func (v *Vector3) SubScalar( scalar float64 ) *Vector3 {
	v.X -= scalar
	v.Y -= scalar
	v.Z -= scalar

	return v
}

func (v *Vector3) SubVectors( a, b *Vector3) *Vector3 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z

	return v
}

func (v *Vector3) Multiply( other *Vector3 ) *Vector3 {

	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z

	return v
}

func (v *Vector3) MultiplyScalar( scalar float64 ) *Vector3 {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar

	return v
}

func (v *Vector3) MultiplyVectors( a, b *Vector3 ) *Vector3 {
	v.X = a.X * b.X
	v.Y = a.Y * b.Y
	v.Z = a.Z * b.Z

	return v
}

func (v *Vector3) ApplyEuler() {
	// TODO After Quaternion 
}

func (v *Vector3) ApplyAxisAngle() {
	// TODO After Quaternion
}

func (v *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {
	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements

	v.X = e[0] * x + e[3] * y + e[6] * z
	v.Y = e[1] * x + e[4] * y + e[7] * z
	v.Z = e[2] * x + e[5] * y + e[8] * z

	return v
}

func (v *Vector3) ApplyMatrix4() {
	// TODO After Matrix4
}

func (v *Vector3) ApplyProjection() {
	// TODO After Matrix4
}

func (v *Vector3) ApplyQuaternion() {
	// TODO After Quaternion
}

func (v *Vector3) Project() {
	// TODO After Matrix 
}

func (v *Vector3) Unproject() {
	// TODO After Matrix 
}

func (v *Vector3) TransformDirection() {
	// TODO After Matrix4
}

func (v *Vector3) Divide( other *Vector3) *Vector3 {
	v.X /= other.X
	v.Y /= other.Y
	v.Z /= other.Z

	return v
}

func (v *Vector3) DivideScalar( scalar float64 ) *Vector3 {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar

	return v
}

func (v *Vector3) Min( other *Vector3 ) *Vector3 {
	v.X = math.Min( v.X, other.X )
	v.Y = math.Min( v.Y, other.Y )
	v.Z = math.Min( v.Z, other.Z )

	return v
}

func (v *Vector3) Max( other *Vector3 ) *Vector3 {
	v.X = math.Max( v.X, other.X )
	v.Y = math.Max( v.Y, other.Y )
	v.Z = math.Max( v.Z, other.Z )

	return v
}

func (v *Vector3) Clamp( min, max *Vector3 ) *Vector3 {
	v.X = math.Max( min.X, math.Min( max.X, v.X ))
	v.Y = math.Max( min.Y, math.Min( max.Y, v.Y ))
	v.Z = math.Max( min.Z, math.Min( max.Z, v.Z ))

	return v
}

func (v *Vector3) ClampLength( min, max float64 ) *Vector3 {
	length := v.Length()

	return v.MultiplyScalar( math.Max( min, math.Min( max, length ) ) / length )
}

func (v *Vector3) Floor() *Vector3 {
	v.X = math.Floor( v.X )
	v.Y = math.Floor( v.Y )
	v.Z = math.Floor( v.Z )

	return v
}

func (v *Vector3) Ceil() *Vector3 {
	v.X = math.Ceil( v.X )
	v.Y = math.Ceil( v.Y )
	v.Z = math.Ceil( v.Z )

	return v
}

func (v *Vector3) Round() *Vector3 {
	v.X = math.Floor( v.X + .5 )
	v.Y = math.Floor( v.Y + .5 )
	v.Z = math.Floor( v.Z + .5 )

	return v
}

func (v *Vector3) RoundToZero() *Vector3 {
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

	if v.Z < 0 {
		v.Z = math.Ceil( v.Z )
	} else {
		v.Z = math.Floor( v.Z )
	}

	return v
}


func (v *Vector3) Negate() *Vector3 {
	v.X = - v.X
	v.Y = - v.Y
	v.Z = - v.Z

	return v
}

func (v *Vector3) Dot( other *Vector3) float64 {
	return v.X * other.X + v.Y * other.Y + v.Z * other.Z
}

func (v *Vector3) LengthSq() float64 {
	return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v *Vector3) LengthManhattan() float64 {
	return math.Abs( v.X ) + math.Abs( v.Y ) + math.Abs( v.Z )
}

func (v *Vector3) Normalize() *Vector3 {
	return v.DivideScalar( v.Length() )
}

func (v *Vector3) SetLength( length float64 ) *Vector3 {
	orgLength := v.Length();

	v.MultiplyScalar( length )
	v.DivideScalar( orgLength )

	return v
}

func (v *Vector3) Lerp( other *Vector3, alpha float64 ) *Vector3 {
	v.X += ( other.X - v.X ) * alpha
	v.Y += ( other.Y - v.Y ) * alpha
	v.Z += ( other.Z - v.Z ) * alpha

	return v
}

func (v *Vector3) LerpVectors( v1, v2 *Vector3, alpha float64 ) *Vector3 {
	return v.SubVectors( v2, v1).MultiplyScalar( alpha ).Add( v1 )
}

func (v *Vector3) Cross( other *Vector3 ) *Vector3 {
	x := v.X
	y := v.Y
	z := v.Z

	v.X = y * v.Z - z * v.Y
	v.Y = z * v.X - x * v.Z
	v.Z = x * v.Y - y * v.X

	return v
}

func (v *Vector3) CrossVectors( a, b *Vector3 ) *Vector3 {
	ax := a.X
	ay := a.Y
	az := a.Z

	bx := b.X
	by := b.Y
	bz := b.Z

	v.X = ay * bz - az * by
	v.Y = az * bx - ax * bz
	v.Z = ax * by - ay * bx

	return v
}

func (v *Vector3) ProjectOnVector( vector *Vector3) *Vector3 {
	scalar := vector.Dot(v) / vector.LengthSq()

	return v.Copy(vector).MultiplyScalar(scalar)
}

func (v *Vector3) ProjectOnPlane( planeNormal *Vector3) *Vector3 {
	v1 := &Vector3{}

	v1.Copy( v ).ProjectOnVector( planeNormal )

	return v.Sub( v1 )
}

func (v *Vector3) Reflect( normal *Vector3 ) *Vector3 {
	v1 := &Vector3{}

	return v.Sub( v1.Copy( normal ).MultiplyScalar( 2 * v.Dot( normal) ) )
}

func (v *Vector3) AngleTo( other *Vector3) float64 {
	theta := v.Dot(other) / ( math.Sqrt( v.LengthSq() * other.LengthSq() ) )

	return math.Acos( Clamp( theta, -1, 1 ) )
}

func (v *Vector3) DistanceTo( other *Vector3 ) float64 {
	return math.Sqrt( v.DistanceToSquared(other) )
}

func (v *Vector3) DistanceToSquared( other *Vector3 ) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z

	return dx * dx + dy * dy + dz * dz
}

func (v *Vector3) SetFromSpherical() {
	// TODO After Spherical
}

func (v *Vector3) SetFromMatrixPosition() {
	// TODO After Matrix 
}

func (v *Vector3) SetFromMatrixScale() {
	// TODO After Matrix 
}

func (v *Vector3) SetFromMatrixColumn() {
}

func (v *Vector3) Equals( other *Vector3 ) bool {
	return ( ( v.X == other.X ) && ( v.Y == other.Y ) && ( v.Z == other.Z ) )
}

func (v *Vector3) FromArray( array []float64 ) *Vector3 {
	v.X = array[0]
	v.Y = array[1]
	v.X = array[2]

	return v
}

func (v *Vector3) ToArray( array []float64 ) []float64 {
	array[0] = v.X
	array[1] = v.Y
	array[2] = v.Z

	return array
}

func (v *Vector3) FromAttribute() {
	// TODO After Attribute
}
