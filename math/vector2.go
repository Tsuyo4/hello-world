package math

type Vector2 struct {
	X float64
	Y float64
}


func (v *Vector2) Copy( other *Vector2) *Vector2 {
	v.X = other.X	
	v.Y = other.Y

	return v
}

func (v *Vector2) Set( x, y float64 ) *Vector2 {
	v.X = x
	v.Y = y

	return v	
}
