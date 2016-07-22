package math

import "testing"

func TestConstructor(test *testing.T) {
	a := &Vector2{}

	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	a = &Vector2{ X:1, Y:1 }
	if a.X != 1 || a.Y != 1 {
		test.Errorf("Fail")
	}
}

func TestVector2Copy( test *testing.T ) {
	a := &Vector2{X:1, Y:2}
	b := &Vector2{}
	
	b.Copy(a)

	if a.X != 1 || a.Y != 2 {
		test.Errorf("Fail")
	}
}

func TestVector2Set( test *testing.T) {
	a := &Vector2{}
	if a.X != 0 || a.Y != 0 {
		test.Errorf("Fail")
	}

	a.Set( 1, 2)
	if a.X != 1 || a.Y != 2 {
		test.Errorf("Fail")
	}	
}
