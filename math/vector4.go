package math

import "math"

type Vector4 struct {
	X float64
	Y float64
	Z float64
	W float64
}


func NewVector4( x, y, z float64) *Vector4 {
	v := &Vector4{ X:x, Y:y , Z:z, W:1 }
	return v
}

func (v *Vector4) Set( x, y, z, w float64 ) *Vector4 {
	v.X = x
	v.Y = y
	v.Z = z
	v.W = w

	return v
}

func (v *Vector4) SetScalar( scalar float64 ) *Vector4 {
	v.X = scalar
	v.Y = scalar
	v.Z = scalar
	v.W = scalar

	return v
}

func (v *Vector4) SetX( x float64 ) *Vector4 {
	v.X = x

	return v
}

func (v *Vector4) SetY( y float64 ) *Vector4 {
	v.Y = y

	return v
}

func (v *Vector4) SetZ( z float64 ) *Vector4 {
	v.Z = z

	return v
}

func (v *Vector4) SetW( w float64 ) *Vector4 {
	v.W = w

	return v
}


func (v *Vector4) SetComponent( index int, value float64) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	case 3:
		v.W = value
	}
}

func (v *Vector4) GetComponent( index int, value float64 ) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	case 3:
		return v.W
	}

	return 0
}

func (v *Vector4) Clone() *Vector4 {
	v1 := &Vector4{ X:v.X, Y:v.Y, Z:v.Z, W:v.W }

	return v1
}

func (v *Vector4) Copy( other *Vector4 ) *Vector4 {
	v.X = other.X
	v.Y = other.Y
	v.Z = other.Z
	v.W = other.W

	return v
}

func (v *Vector4) Add( other *Vector4 ) *Vector4 {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
	v.W += other.W
	
	return v
}

func (v *Vector4) AddScalar( scalar float64 ) *Vector4 {
	v.X += scalar
	v.Y += scalar
	v.Z += scalar
	v.W += scalar

	return v
}

func (v *Vector4) AddVectors( a, b Vector4) *Vector4 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	v.W = a.W + b.W

	return v
}

func (v *Vector4) AddScaledVector( other Vector4, s float64 ) *Vector4 {
	v.X += other.X * s
	v.Y += other.Y * s
	v.Z += other.Z * s
	v.W += other.W * s

	return v
}

func (v *Vector4) Sub( other *Vector4 ) *Vector4 {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
	v.W -= other.W
	
	return v
}

func (v *Vector4) SubScalar( scalar float64 ) *Vector4 {
	v.X -= scalar
	v.Y -= scalar
	v.Z -= scalar
	v.W -= scalar

	return v
}

func (v *Vector4) SubVectors( a, b *Vector4) *Vector4 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	v.W = a.W - b.W

	return v
}

func (v *Vector4) MultiplyScalar( scalar float64 ) *Vector4 {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	v.W *= scalar

	return v
}

func (v *Vector4) ApplyMatrix4() {
	// TODO After Matrix4
}

func (v *Vector4) DivideScalar( scalar float64 ) *Vector4 {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
	v.W /= scalar

	return v
}

func (v *Vector4) SetAxisAngleFromQuaternion() {
	// TODO After Quaternion
}

func (v *Vector4) SetAxisAngleFromRotationMatrix() {
	// TODO After Matrix
}

func (v *Vector4) Min( other *Vector4 ) *Vector4 {
	v.X = math.Min( v.X, other.X )
	v.Y = math.Min( v.Y, other.Y )
	v.Z = math.Min( v.Z, other.Z )
	v.W = math.Min( v.W, other.W )

	return v
}

func (v *Vector4) Max( other *Vector4) *Vector4 {
	v.X = math.Max( v.X, other.X )
	v.Y = math.Max( v.Y, other.Y )
	v.Z = math.Max( v.Z, other.Z )
	v.W = math.Max( v.W, other.W )

	return v
}

func (v *Vector4) Clamp( min, max *Vector4 ) *Vector4 {
	v.X = math.Max( min.X, math.Min( max.X, v.X) )
	v.Y = math.Max( min.Y, math.Min( max.Y, v.Y) )
	v.Z = math.Max( min.Z, math.Min( max.Z, v.Z) )
	v.W = math.Max( min.W, math.Min( max.W, v.W) )

	return v
}

func (v *Vector4) ClampScalar( minVal, maxVal float64) *Vector4 {
	min := &Vector4{ X:minVal, Y:minVal, Z:minVal, W:minVal }
	max := &Vector4{ X:maxVal, Y:maxVal, Z:maxVal, W:maxVal }

	return v.Clamp( min, max)
}

func (v *Vector4) Floor() *Vector4 {
	v.X = math.Floor( v.X )
	v.Y = math.Floor( v.Y )
	v.Z = math.Floor( v.Z )
	v.W = math.Floor( v.W )

	return v
}

func (v *Vector4) Ceil() *Vector4 {
	v.X = math.Ceil( v.X )
	v.Y = math.Ceil( v.Y )
	v.Z = math.Ceil( v.Z )
	v.W = math.Ceil( v.W )

	return v
}

func (v *Vector4) Round() *Vector4 {
	v.X = math.Floor( v.X + .5 )
	v.Y = math.Floor( v.Y + .5 )
	v.Z = math.Floor( v.Z + .5 )
	v.W = math.Floor( v.W + .5 )

	return v
}

func (v *Vector4) RoundToZero() *Vector4 {
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

	if v.W < 0 {
		v.W = math.Ceil( v.W )
	} else {
		v.W = math.Floor( v.W )
	}

	return v
}


func (v *Vector4) Negate() *Vector4 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W

	return v
}

func (v *Vector4) Dot( other *Vector4 ) float64 {
	return v.X * other.X + v.Y * other.Y + v.Z * other.Z + v.W * other.W
}

func (v *Vector4) LengthSq() float64 {
	return v.X * v.X + v.Y * v.Y + v.Z * v.Z + v.W * v.W
}

func (v *Vector4) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z + v.W * v.W)
}

func (v *Vector4) LengthManhattan() float64 {
	return math.Abs( v.X ) + math.Abs( v.Y ) + math.Abs( v.Z ) + math.Abs( v.W )
}

func (v *Vector4) Normalize() *Vector4 {
	return v.DivideScalar( v.Length() )
}

func (v *Vector4) SetLength( length float64 ) *Vector4 {
	orgLength := v.Length()

	v.MultiplyScalar( length )
	v.DivideScalar( orgLength )

	return v
}

func (v *Vector4) Lerp( other *Vector4, alpha float64 ) *Vector4 {
	v.X += ( other.X - v.X ) * alpha
	v.Y += ( other.Y - v.Y ) * alpha
	v.Z += ( other.Z - v.Z ) * alpha
	v.W += ( other.W - v.W ) * alpha

	return v
}

func (v *Vector4) LerpVectors( v1, v2 *Vector4, alpha float64) *Vector4 {
	return v.SubVectors(v2,v1).MultiplyScalar( alpha ).Add( v1 )
}

func (v *Vector4) Equals( other *Vector4 ) bool {
	return ( (v.X == other.X) && (v.Y == other.Y) && (v.Z == other.Z) && (v.W == other.W) )
}

func (v *Vector4) FromArray( array []float64 ) *Vector4 {
	v.X = array[0]
	v.Y = array[1]
	v.Z = array[2]
	v.W = array[3]

	return v
}

func (v *Vector4) ToArray( array []float64) []float64 {
	array[0] = v.X
	array[1] = v.Y
	array[2] = v.Z
	array[3] = v.W

	return array
}

func (v *Vector4) FromAttribute() {
	// TODO After Attribute
}
