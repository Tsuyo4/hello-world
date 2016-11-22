package math

import "math"

type Box3 struct {
	Min *Vector3
	Max *Vector3
}

func NewBox3() *Box3 {
	b := &Box3{
		Min:&Vector3{ X:math.Inf(0),  Y:math.Inf(0),  Z:math.Inf(0)},
		Max:&Vector3{ X:math.Inf(-1), Y:math.Inf(-1), Z:math.Inf(-1)},
	}
	return b
}

func (b *Box3) Set( min, max *Vector3 ) *Box3 {
	b.Min.Copy( min )
	b.Max.Copy( max )

	return b
}

func (b *Box3) SetFromArray( array []float64 ) {
	minX := math.Inf(0)
	minY := math.Inf(0)
	minZ := math.Inf(0)

	maxX := math.Inf(-1)
	maxY := math.Inf(-1)
	maxZ := math.Inf(-1)

	for i := 0 ; i < len(array); i += 3 {
		x := array[ i ]
		y := array[ i + 1 ]
		z := array[ i + 2 ]

		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if z < minZ {
			minZ = z
		}

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if z > maxZ {
			maxZ = z
		}
		
	}

	b.Min.Set( minX, minY, minZ )
	b.Max.Set( maxX, maxY, maxZ )
}


func (b *Box3) SetFromPoints( points []Vector3 ) *Box3 {
	b.MakeEmpty()

	for i := 0 ; i < len(points); i++ {
		b.ExpandByPoint( &points[i] )
	}

	return b
}

func (b *Box3) SetFromCenterAndSize( center, size *Vector3 ) *Box3 {
	v1 := &Vector3{}

	halfSize := v1.Copy( size ).MultiplyScalar( 0.5 )

	b.Min.Copy( center ).Sub( halfSize )
	b.Max.Copy( center ).Add( halfSize )

	return b
}

func (b *Box3) SetFromObject( ) {
	// TODO After core/Object3D
}

func (b *Box3) Clone() *Box3 {
	b1 := &Box3{}
	b1.Set( b.Min, b.Max )

	return b1
}

func (b *Box3) Copy( box *Box3 ) *Box3 {
	b.Min.Copy( box.Min )
	b.Max.Copy( box.Max )

	return b
}

func (b *Box3) MakeEmpty() *Box3 {
	b.Min.X = math.Inf(0)
	b.Min.Y = math.Inf(0)
	b.Min.Z = math.Inf(0)

	b.Max.X = math.Inf(-1)
	b.Max.Y = math.Inf(-1)
	b.Max.Z = math.Inf(-1)

	return b
}

func (b *Box3) IsEmpty() bool {
	return ( b.Max.X < b.Min.X ) || ( b.Max.Y < b.Min.Y ) || ( b.Max.Z < b.Min.Z )
}

func (b *Box3) Center() *Vector3 {
	result := &Vector3{}

	return b.CenterTarget(result)
}

func (b *Box3) CenterTarget( optionalTarget *Vector3 ) *Vector3 {
	return optionalTarget.AddVectors( b.Min, b.Max ).MultiplyScalar( 0.5 )
}

func (b *Box3) Size() *Vector3 {
	result := &Vector3{}

	return b.SizeTarget(result)
}

func (b *Box3) SizeTarget( optionalTarget *Vector3 ) *Vector3 {
	return optionalTarget.SubVectors( b.Max, b.Min )
}

func (b *Box3) ExpandByPoint( point *Vector3 ) *Box3 {
	b.Min.Min( point )
	b.Max.Max( point )

	return b
}

func (b *Box3) ExpandByVector( vector *Vector3 ) *Box3 {
	b.Min.Sub( vector )
	b.Max.Add( vector )

	return b
}

func (b *Box3) ExpandByScalar( scalar float64 ) *Box3 {
	b.Min.AddScalar( -scalar )
	b.Max.AddScalar( scalar )

	return b
}

func (b *Box3) ContainsPoint( point Vector3 ) bool {
	if point.X < b.Min.X || point.X > b.Max.X ||
		point.Y < b.Min.Y || point.Y > b.Max.Y ||
		point.Z < b.Min.Z || point.Z > b.Max.Z {
		return false
	}

	return true
}

func (b *Box3) ContainsBox( box *Box3 ) bool {
	if ( b.Min.X <= box.Min.X ) && ( box.Max.X <= b.Max.X ) &&
		( b.Min.Y <= box.Min.Y ) && ( box.Max.Y <= b.Max.Y ) &&
		( b.Min.Z <= box.Min.Z ) && ( box.Max.Z <= b.Max.Z ) {
		return true
	}
	return false
}

func (b *Box3) GetParamaeter( point *Vector3 ) *Vector3 {
	result := &Vector3{}

	return b.GetParamaeterTarget( point, result )
}

func (b *Box3) GetParamaeterTarget( point *Vector3, optionalTarget *Vector3 ) *Vector3 {
	return optionalTarget.Set(
		( point.X - b.Min.X ) / ( b.Max.X - b.Min.X ),
		( point.Y - b.Min.Y ) / ( b.Max.Y - b.Min.Y ),
		( point.Z - b.Min.Z ) / ( b.Max.Z - b.Min.Z ),
	)
}

func (b *Box3) IntersectsBox( box *Box3 ) bool {
	if box.Max.X < b.Min.X || box.Min.X > b.Max.X ||
		box.Max.Y < b.Min.Y || box.Min.Y > b.Max.Y ||
		box.Max.Z < b.Min.Z || box.Min.Z > b.Max.Z {
		return false
	}

	return true
}

func (b *Box3) IntersectsSphere() {
	// TODO After Sphere
}

func (b *Box3) IntersectsPlane() {
	// TODO After Plane
}

func (b *Box3) ClampPoint( point *Vector3 ) *Vector3 {
	result := &Vector3{}

	return b.ClampPointTarget( point, result )
}

func (b *Box3) ClampPointTarget( point, optionalTarget *Vector3 ) *Vector3 {
	return optionalTarget.Copy( point ).Clamp( b.Min, b.Max )
}

func (b *Box3) DistanceToPoint( point *Vector3 ) float64 {
	v1 := &Vector3{}

	clampedPoint := v1.Copy( point ).Clamp( b.Min, b.Max )
	
	return clampedPoint.Sub( point ).Length()
}

func (b *Box3) GetBoundingSphere() {
	// TODO After Sphere
}


func (b *Box3) Intersect( box *Box3 ) *Box3 {
	b.Min.Max( box.Min )
	b.Max.Min( box.Max )

	if b.IsEmpty() {
		b.MakeEmpty()
	}

	return b
}

func (b *Box3) Union( box *Box3 ) *Box3 {
	b.Min.Min( box.Min )
	b.Max.Max( box.Max )

	return b
}

func (b *Box3) ApplyMatrix4() {
	// TODO After Matrix4
}

func (b *Box3) Translate( offset *Vector3 ) *Box3 {
	b.Min.Add( offset )
	b.Max.Add( offset )

	return b
}

func (b *Box3) Equals( box Box3 ) bool {
	return box.Min.Equals( b.Min ) && box.Max.Equals( b.Max )
}
