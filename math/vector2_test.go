package math

import "testing"
import "time"
import "math/rand"
import "math"

var X float64
var Y float64

func init() {
	rand.Seed(time.Now().UnixNano())
	X = math.Floor(rand.Float64()*1000)
	Y = math.Floor(rand.Float64()*1000)
}

func TestConstructor(test *testing.T) {
	a := &Vector2{}

	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	a = &Vector2{ X:X, Y:Y }
	if a.X != X || a.Y != Y {
		test.Errorf("Fail")
	}
}

func TestVector2Copy( test *testing.T ) {
	a := &Vector2{X:X, Y:Y}
	b := &Vector2{}
	
	b.Copy(a)

	if a.X != X || a.Y != Y {
		test.Errorf("Fail")
	}

	a.X = 0
	a.Y = -1

	if b.X != X || b.Y != Y {
		test.Errorf("Fail")
	}
}

func TestVector2Set( test *testing.T) {
	a := &Vector2{}
	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	a.Set(X,Y)
	if a.X != X || a.Y != Y {
		test.Errorf("Fail")
	}
}

func TestSetXSetY( test *testing.T ) {
	a := &Vector2{}
	
	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	a.SetX(1)
	a.SetY(2)
	if a.X != 1 || a.Y != 2 {
		test.Errorf("Fail")
	}
}

func TestSetGetComponent( test *testing.T ) {
	a := &Vector2{}
	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}


	a.SetComponent( 0, 1 )
	a.SetComponent( 1, 2 )
	if a.GetComponent( 0 ) != 1  || a.GetComponent( 1 ) != 2 {
		test.Errorf("Fail")
	}
}

func TestVector2Add( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y}
	b := &Vector2{ X:-X, Y:-Y}

	a.Add( b )
	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	c := &Vector2{}
	c.AddVectors( b, b )
	
		if c.X != -2*X  || c.Y != -2*Y {
		test.Errorf("Fail")
	}
	
}

func TestVector2Sub( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y }
	b := &Vector2{ X:-X, Y:-Y }

	a.Sub(b)
	if a.X != 2*X || a.Y != 2*Y {
		test.Errorf("Fail")
	}

	c := &Vector2{}
	c.SubVectors( a, a )
	if c.X != 0 || c.Y != 0 {
		test.Errorf("Fail")
	}
}

func TestVector2MultiplyDivide( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y }
	b := &Vector2{ X:-X, Y:-Y }

	a.MultiplyScalar( -2 )
	if a.X != X*-2 || a.Y != Y*-2 {
		test.Errorf("Fail")
	}

	b.MultiplyScalar( -2 )
	if b.X != X*2 || b.Y != Y*2 {
		test.Errorf("Fail")
	}

	a.DivideScalar( -2 )
	if a.X != X || a.Y != Y {
		test.Errorf("Fail")
	}

	b.DivideScalar( -2 )
	if b.X != -X || b.Y != -Y {
		test.Errorf("Fail")
	}
}

func TestVector2MinMaxClamp( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y }
	b := &Vector2{ X:-X, Y:-Y }
	c := &Vector2{}

	c.Copy( a ).Min( b )
	if c.X != -X || c.Y != -Y {
		test.Errorf("Fail")
	}

	c.Copy( a ).Max( b )
	if c.X != X || c.Y != Y {
		test.Errorf("Fail")
	}

	c.Set( -2*X, 2*Y )
	c.Clamp( b, a )
	if c.X != -X || c.Y != Y {
		test.Errorf("Fail")
	}

	c.Set( -2*X, 2*X )
	c.ClampScalar( -X, X );
	if c.X != -X ||  c.Y != X {
		test.Errorf("Fail Scalar Clamp X:")
	}
	
}

func deepEqual( v1 *Vector2, v2 *Vector2) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func TestVector2Rounding( test *testing.T ) {
	if !deepEqual( (&Vector2{X:-0.1, Y:0.1}).Floor() , &Vector2{X:-1, Y:0} ) {
		test.Errorf("Floor .1")
	}
	if !deepEqual( (&Vector2{X:-0.5, Y:0.5}).Floor() , &Vector2{X:-1, Y:0} ) {
		test.Errorf("Floor .5")
	}
	if !deepEqual( (&Vector2{X:-0.9, Y:0.9}).Floor() , &Vector2{X:-1, Y:0} ) {
		test.Errorf("Floor .5")
	}

	if !deepEqual( (&Vector2{X:-0.1, Y:0.1}).Ceil(), &Vector2{X:0, Y:1} ) {
		test.Errorf("Ceil .1")
	}
	if !deepEqual( (&Vector2{X:-0.5, Y:0.5}).Ceil(), &Vector2{X:0, Y:1} ) {
		test.Errorf("Ceil .5")
	}
	if !deepEqual( (&Vector2{X:-0.9, Y:0.9}).Ceil(), &Vector2{X:0, Y:1} ) {
		test.Errorf("Ceil .9")
	}

	if !deepEqual( (&Vector2{X:-0.1, Y:0.1}).Round(), &Vector2{X:0, Y:0} ) {
		test.Errorf("Round .1")
	}
	if !deepEqual( (&Vector2{X:-0.5, Y:0.5}).Round(), &Vector2{X:0, Y:1} ) {
		test.Errorf("Round .5")
	}
	if !deepEqual( (&Vector2{X:-0.9, Y:0.9}).Round(), &Vector2{X:-1, Y:1} ) {
		test.Errorf("Round .9")
	}

	if !deepEqual( (&Vector2{X:-0.1, Y:0.1}).RoundToZero(), &Vector2{X:0, Y:0} ) {
		test.Errorf("RoundToZero .1")
	}
	if !deepEqual( (&Vector2{X:-0.5, Y:0.5}).RoundToZero(), &Vector2{X:0, Y:0} ) {
		test.Errorf("RoundToZero .5")
	}
	if !deepEqual( (&Vector2{X:-0.9, Y:0.9}).RoundToZero(), &Vector2{X:0, Y:0} ) {
		test.Errorf("RoundToZero .9")
	}
	if !deepEqual( (&Vector2{X:-1.1, Y:1.1}).RoundToZero(), &Vector2{X:-1, Y:1} ) {
		test.Errorf("RoundToZero 1.1")
	}
	if !deepEqual( (&Vector2{X:-1.5, Y:1.5}).RoundToZero(), &Vector2{X:-1, Y:1} ) {
		test.Errorf("RoundToZero 1.5")
	}
	if !deepEqual( (&Vector2{X:-1.9, Y:1.9}).RoundToZero(), &Vector2{X:-1, Y:1} ) {
		test.Errorf("RoundToZero 1.9")
	}
}

func TestVector2Negate( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y }

	a.Negate()

	if a.X != -X || a.Y != -Y {
		test.Errorf("Fail")
	}
}

func TestVector2Dot( test *testing.T ) {
	a := &Vector2{ X:X, Y:Y }
	b := &Vector2{ X:-X, Y:-Y }
	c := &Vector2{}

	result := a.Dot(b)

	if result != (-X*X-Y*Y) {
		test.Errorf("Fail")
	}

	result = a.Dot(c)

	if result != 0 {
		test.Errorf("Fail")
	}
}


func TestVector2Length( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }
	b := &Vector2{ X:0, Y:-Y }
	c := &Vector2{}

	if a.Length() != X {
		test.Errorf("Fail")
	}
	if a.LengthSq() != X*X {
		test.Errorf("Fail")
	}
	if b.Length() != Y {
		test.Errorf("Fail")
	}
	if b.LengthSq() != Y*Y {
		test.Errorf("Fail")
	}
	if c.Length() != 0 {
		test.Errorf("Fail")
	}
	if c.Length() != 0 {
		test.Errorf("Fail")
	}

	a.Set( X, Y )
	if a.Length() != math.Sqrt( X*X+Y*Y ) {
		test.Errorf("Fail")
	}
	if a.LengthSq() != X*X + Y*Y {
		test.Errorf("Fail")
	}
	
}

func TestVector2Normalize( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }
	b := &Vector2{ X:0, Y:-Y }

	a.Normalize()
	if a.Length() != 1 {
		test.Errorf("Fail X:%g Y:%g A.Length:%g", X, Y, a.Length())
	}
	if a.X != 1 {
		test.Errorf("Fail X:%g Y:%g", X, Y)
	}

	b.Normalize()
	if b.Length() != 1 {
		test.Errorf("Fail X:%g Y:%g B.Length:%g", X, Y, b.Length())
	}
	if b.Y != -1 {
		test.Errorf("Fail X:%g Y:%g", X, Y)
	}
}

func TestVector2DistanceTo( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }
	b := &Vector2{ X:0, Y:-Y }
	c := &Vector2{}

	if a.DistanceTo(c) != X {
		test.Errorf("Fail")
	}
	if a.DistanceToSquared(c) != X*X {
		test.Errorf("Fail")
	}

	if b.DistanceTo(c) != Y {
		test.Errorf("Fail")
	}
	if b.DistanceToSquared(c) != Y*Y {
		test.Errorf("Fail")
	}
}


func TestVector2SetLength( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }

	if a.Length() != X {
		test.Errorf("Fail")
	}
	a.SetLength(Y)
	if a.Length() != Y {
		test.Errorf("Fail X:%g Y:%g", X, Y)
	}

	a = &Vector2{ 0, 0 }
	if a.Length() != 0 {
		test.Errorf("Fail")
	}
	a.SetLength(Y)
	if a.Length() != 0 {
		test.Errorf("Fail:%g",a.Length())
	}
}

func TestVector2LerpClone( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }
	b := &Vector2{ X:0, Y:-Y }

	if a.Lerp( a, 0 ).Equals( a.Lerp( a, 0.5 ) ) == false {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}
	if a.Lerp( a, 0 ).Equals( a.Lerp( a, 1 ) ) == false {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}

	if a.Clone().Lerp( b, 0 ).Equals( a ) == false {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}

	if a.Clone().Lerp( b, 0.5 ).X != X*0.5 {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}
	if a.Clone().Lerp( b, 0.5 ).Y != -Y*0.5 {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}

	if a.Clone().Lerp( b, 1 ).Equals( b ) == false {
		test.Errorf("Fail X:%f Y:%f", X, Y)
	}
	
}

func TestVector2Equals( test *testing.T ) {
	a := &Vector2{ X:X, Y:0 }
	b := &Vector2{ X:0, Y:-Y}

	if a.X == b.X {
		test.Errorf("Fail")
	}
	if a.Y == b.Y {
		test.Errorf("Fail")
	}

	if a.Equals(b) == true {
		test.Errorf("Fail")
	}
	if b.Equals(a) == true {
		test.Errorf("Fail")
	}

	a.Copy(b)

	if a.Equals(b) == false {
		test.Errorf("Fail")
	}
	if b.Equals(a) == false {
		test.Errorf("Fail")
	}
}
