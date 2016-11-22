package math

import "math"

type Box2 struct {
	Min *Vector2
	Max *Vector2
}

func NewBox2() *Box2 {
	b := &Box2{
		Min:&Vector2{X:math.Inf(0),  Y:math.Inf(0)},
		Max:&Vector2{X:math.Inf(-1), Y:math.Inf(-1)},
	}

	return b
}

func (b *Box2) Set(min, max *Vector2) *Box2 {
	b.Min.Copy( min )
	b.Max.Copy( max )

	return b
}

func (b *Box2) SetFromPoint( points []Vector2 ) *Box2 {
	b.MakeEmpty()
	
	for i := 0 ; i < len(points); i++  {
		b.ExpandByPoint( &points[i] )
	}

	return b
}

func (b *Box2) SetFromCenterAndSize( center, size *Vector2) *Box2 {
	v1 := &Vector2{}
	halfSize := v1.Copy( size ).MultiplyScalar( 0.5 )
	b.Min.Copy( center ).Sub( halfSize )
	b.Max.Copy( center ).Add( halfSize )

	return b
}

func (b *Box2) Clone() *Box2 {
	b1 := &Box2{}
	b1.Set( b.Min, b.Max )

	return b1
}

func (b *Box2) Copy( box *Box2 ) *Box2 {
	b.Min.Copy( box.Min )
	b.Max.Copy( box.Max )

	return b
}

func (b *Box2) MakeEmpty() *Box2 {
	b.Min.X = math.Inf(0)
	b.Min.Y = math.Inf(0)

	b.Max.X = math.Inf(-1)
	b.Max.Y = math.Inf(-1)

	return b
}

func (b *Box2) IsEmpty() bool {
	return ( b.Max.X < b.Min.X ) || ( b.Max.Y < b.Min.Y )
}

func (b *Box2) SizeTarget( optionalTarget *Vector2 ) *Vector2 {
	return optionalTarget.SubVectors( b.Max, b.Min )
}

func (b *Box2) Size() *Vector2 {
	result := &Vector2{}
	
	return result.SubVectors( b.Max, b.Min )
}

func (b *Box2) ExpandByPoint( point *Vector2 ) *Box2 {
	b.Min.Min( point )
	b.Max.Max( point )

	return b
}

func (b *Box2) ExpandByVector( vector *Vector2) *Box2 {
	b.Min.Sub( vector )
	b.Max.Add( vector )

	return b
}

func (b *Box2) ExpandByScalar( scalar float64 ) *Box2 {
	b.Min.AddScalar( -scalar )
	b.Max.AddScalar( scalar)

	return b
}

func (b *Box2) ContainsPoint( point *Vector2) bool {
	if point.X < b.Min.X || point.X > b.Max.X || point.Y < b.Min.Y || point.Y > b.Max.Y {
		return false
	}
	return true
}

func (b *Box2) ContainsBox ( box *Box2) bool {
	if ( b.Min.X <= box.Min.X ) && ( box.Max.X <= b.Max.X ) &&
		( b.Min.Y <= box.Min.Y ) && ( box.Max.X <= b.Max.Y ) {
		return true
	}

	return false
}

func (b *Box2) GetParamaeter( point *Vector2 ) *Vector2 {
	result := &Vector2{}

	return result.Set(
		( point.X - b.Min.X ) / ( b.Max.X - b.Min.X ),
		( point.Y - b.Min.Y ) / ( b.Max.Y - b.Min.Y ))
}

func (b *Box2) GetParamaeterTarget( point, optionalTarget  *Vector2) *Vector2 {
	return optionalTarget.Set(
		( point.X - b.Min.X ) / ( b.Max.X - b.Min.X ),
		( point.Y - b.Min.Y ) / ( b.Max.Y - b.Min.Y ))
}

func (b *Box2) IntersectsBox( box *Box2 ) bool {
	if box.Max.X < b.Min.X || box.Min.X > b.Max.X || box.Max.Y < b.Min.Y || box.Min.Y > b.Max.Y {
		return false
	}
	
	return true
}

func (b *Box2) ClampPoint( point *Vector2 ) *Vector2 {
	result := &Vector2{}

	return result.Copy( point ).Clamp( b.Min, b.Max )
}

func (b *Box2) ClampPointTarget( point, optionalTarget *Vector2 ) *Vector2 {	
	return optionalTarget.Copy( point ).Clamp( b.Min, b.Max )
}

func (b *Box2) DistanceToPoint( point *Vector2 ) float64 {
	v1 := &Vector2{}

	clampPoint := v1.Copy( point ).Clamp( b.Min, b.Max)

	return clampPoint.Sub( point ).Length()
}

func (b *Box2) Intersect( box *Box2 ) *Box2 {
	b.Min.Max( box.Min )
	b.Max.Min( box.Max )

	return b
}

func (b *Box2) Union( box *Box2 ) *Box2 {
	b.Min.Min( box.Min )
	b.Max.Max( box.Max )

	return b
}

func (b *Box2) Translate( offset *Vector2 ) *Box2 {
	b.Min.Add( offset )
	b.Max.Add( offset )

	return b
}

func (b *Box2) Equals( box *Box2 ) bool {
	return box.Min.Equals( b.Min ) && box.Max.Equals( b.Max )
}
