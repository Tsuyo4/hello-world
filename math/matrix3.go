package math

//import "math"

type Matrix3 struct {
	Elements []float64
}

func NewMatrix3() *Matrix3 {
	e := []float64 {
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}

	m := &Matrix3{ Elements:e }

	return m
}


func (m *Matrix3) Set( n11, n12, n13, n21, n22, n23, n31, n32, n33 float64 ) *Matrix3 {

	te := m.Elements
	te[0] = n11; te[1] = n21; te[2] = n31
	te[3] = n12; te[4] = n22; te[5] = n32
	te[6] = n13; te[7] = n23; te[8] = n33

	return m
}


func (m *Matrix3) Identity() *Matrix3 {
	m.Set(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)

	return m
}

func (m *Matrix3) Clone() *Matrix3 {
	m1 := &Matrix3{}
	m1.FromArray(m.Elements)
 
	return m
}

func (m *Matrix3) Copy( matrix *Matrix3 ) *Matrix3 {
	me :=  matrix.Elements

	m.Set(
		me[0], me[3], me[6],
		me[1], me[4], me[7],
		me[2], me[5], me[8],
	)

	return m
}

// TODO After Matrix4 Uncomment
//func (m *Matrix3) SetFromMatrix4( matrix *Matrix4 ) *Matrix3 {
// 	me := matrix.Elements
// 
// 	m.Set(
// 		me[0], me[4], me[ 8],
// 		me[1], me[5], me[ 9],
// 		me[2], me[6], me[10],
// 	)
// 
// 	return m
//}

func (m *Matrix3) ApplyToVector3Array( array []float64 ) []float64 {
	v1 := &Vector3{}

	for i := 0; i < len(array); i += 3 {
		v1.FromArray( array[i:3] )
//		v1.ApplyMatrix3( m ) TODO Vecor3 Matrix
		v1.ToArray( array )
	}

	return array
}

func (m *Matrix3) ApplyToBuffer() {
	// TODO After Buffer
}

func (m *Matrix3) MultiplyScalar( scalar float64 ) *Matrix3 {
	te := m.Elements

	te[0] *= scalar; te[3] *= scalar ; te[6] *= scalar
	te[1] *= scalar; te[4] *= scalar ; te[7] *= scalar
	te[2] *= scalar; te[5] *= scalar ; te[8] *= scalar

	return m
}

func (m *Matrix3) Determinant() float64 {

	te := m.Elements

	a := te[0] ; b := te[1] ; c := te[2]
	d := te[3] ; e := te[4] ; f := te[5]
	g := te[6] ; h := te[7] ; i := te[8]

	return a * e * i - a * f * h - b * d * i + b * f * g + c * d * h - c * e * g
}

func (m *Matrix3) GetInverse( matrix *Matrix3 ) *Matrix3 {

	me := matrix.Elements
	te := m.Elements

	n11 := me[0]; n21 := me[1] ; n31 := me[2]
	n12 := me[3]; n22 := me[4] ; n32 := me[5]
	n13 := me[6]; n23 := me[7] ; n33 := me[8]

	t11 := n33 * n22 - n32 * n23
	t12 := n32 * n13 - n33 * n12
	t13 := n23 * n12 - n22 * n13

	det := n11 * t11 + n21 * t12 + n31 * t13

	if det == 0 {
		// THREE.Matrix3.getInverse(): can't invert matrix, determinant is 0

		return m.Identity()
	}

	detInv := 1 / det

	te[0] = t11 * detInv
	te[1] = ( n31 * n23 - n33 * n21 ) * detInv
	te[2] = ( n32 * n21 - n31 * n22 ) * detInv

	te[3] = t12 * detInv
	te[4] = ( n33 * n11 - n31 * n13 ) * detInv
	te[5] = ( n31 * n12 - n32 * n11 ) * detInv

	te[6] = t13 * detInv
	te[7] = ( n21 * n13 - n23 * n11 ) * detInv
	te[8] = ( n22 * n11 - n21 * n12 ) * detInv

	return m
}

func (m *Matrix3) Transpose() *Matrix3 {
	var tmp float64
	
	te := m.Elements

	tmp = te[1] ; te[1] = te[3] ; te[3] = tmp
	tmp = te[2] ; te[2] = te[6] ; te[6] = tmp
	tmp = te[5] ; te[5] = te[7] ; te[7] = tmp

	return m
}

/// deprecated
//func (m *Matrix3) FlattenToArrayOffset( array []float64 ) {
// 	return m.ToArray( array ) 
//}

// TODO After Matrix 4
//func (m *Matrix3) GetNormalMatrix( matrix4 *Matrix4 ) *Matrix3 {
// 	return m.SetFromMatrix4( matrix4 ).GetInverse( m ).Transpose()
//}

func (m *Matrix3) TransposeIntoArray( array []float64 ) *Matrix3 {
	te := m.Elements

	array[0] = te[0]
	array[1] = te[3]
	array[2] = te[6]
	array[3] = te[1]
	array[4] = te[4]
	array[5] = te[7]
	array[6] = te[2]
	array[7] = te[5]
	array[8] = te[8]

	return m
}

func (m *Matrix3) FromArray( array []float64 ) *Matrix3 {
	copy(m.Elements, array)

	return m
}

func (m *Matrix3) ToArray( array []float64 ) []float64 {
	copy(array , m.Elements)

	return array
}
