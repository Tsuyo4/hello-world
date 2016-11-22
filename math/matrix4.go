package math

//import "math"

type Matrix4 struct {
	Elements []float64
}

func NewMatrix4() *Matrix4 {
	e:= []float64 {
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	m := &Matrix4{ Elements:e }

	return m
}

func (m *Matrix4) Set( n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64 ) *Matrix4 {
	te := m.Elements

	te[0] = n11; te[4] = n12; te[8]  = n13; te[12] = n14
	te[1] = n21; te[5] = n22; te[9]  = n23; te[13] = n24
	te[2] = n31; te[6] = n32; te[10] = n33; te[14] = n34
	te[3] = n41; te[7] = n42; te[11] = n43; te[15] = n44

	return m
}

func (m *Matrix4) Identity() *Matrix4 {
	m.Set(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	return m
}

func (m *Matrix4) Clone() *Matrix4 {
	m1 := &Matrix4{}
	m1.FromArray{m.Elements}

	return m
}

func (m *Matrix4) Copy(matrix *Matrix4) *Matrix4 {
	copy(m.Elements, matrix.Elements)

	return m
}

func (m *Matrix4) CopyPosition(matrix *Matrix4) *Matrix4 {
	te := m.Elements
	me := matrix.Elements

	te[12] = me[12]
	te[13] = me[13]
	te[14] = me[14]

	return m
}

func (m *Matrix4) ExtractBasis( xAxis, yAxis, zAxis *Vector3 ) *Matrix4 {
	
	xAxis.SetFromMatrixColumn( m, 0 )
	yAxis.SetFromMatrixColumn( m, 1 )
	zAxis.SetFromMatrixColumn( m, 2 )

	return m
}

func (m *Matrix4) MakeBasis( xAxis, yAxis, zAxis *Vector3 ) *Matrix4 {
	m.Set(
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		0,       0,       0,       1,
	)

	return m
}


func (m *Matrix4) ExtractRotation( matrix *Matrix4 ) *Matrix4 {
	v1 := &Vector3{}
	
	te := m.Elements
	me := m.Elements

	scaleX := 1 / v1.SetFromMatrixColumn(matrix, 0).Length()
	scaleY := 1 / v1.SetFromMatrixColumn(matrix, 1).Length()
	scaleZ := 1 / v1.SetFromMatrixColumn(matrix, 2).Length()

	te[0] = me[0] * scaleX
	te[1] = me[1] * scaleX
	te[2] = me[2] * scaleX

	te[4] = me[4] * scaleY
	te[5] = me[5] * scaleY
	te[6] = me[6] * scaleY

	te[8]  = me[8]  * scaleZ
	te[9]  = me[9]  * scaleZ
	te[10] = me[10] * scaleZ

	return m
}

func (m *Matrix4) MakeRotationFromEuler( euler *Euler) *Matrix4 {
	te := m.Elements
 
	x := euler.X
	y := euler.y
	z := euler.z
 
	a := math.Cos(x)
	b := math.Sin(x)
	
	c := math.Cos(y)
	d := math.Sin(y)
 
	e := math.Cos(z)
	f := math.Sin(z)

	if euler.GetOrder() == EulerRotationOrderXYZ {
		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[0] = c * e
		te[4] = -c * f
		te[8] = d

		te[1] = af + be * d
		te[5] = ae - bf * d
		te[9] = -b * c

		te[2] = bf - ae * d
		te[6] = be + af * d
		te[10] = a * c
	} else if euler.GetOrder() == EulerRotationOrderYXZ {
		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[0] = ce + df * b
		te[4] = de + b  - cf
		te[8] = a *d

		te[1] = a * f
		te[5] = a * e
		te[9] = -b

		te[2] = cf * b - de
		te[6] = df + ce * b
		te[10] = a * c
	} else if euler.GetOrder() == EulerRotationOrderZXY {
		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[0] = ce - df * b
		te[4] = -a * f
		te[8] = de + cf * b

		te[1] = cf + de * b
		te[5] = a * e
		te[9] = df - ce * b

		te[2] = -a * d
		te[6] = b
		te[10] = a * c
	} else if euler.GetOrder() == EulerRotationOrderZYX {
		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[0] = c * e
		te[4] = be * d - af
		te[8] = ae * d + bf

		te[1] = c * f
		te[5] = bf * d + ae
		te[9] = af * d - be

		te[2] = -d
		te[6] = b * c
		te[10] = a * c
	} else if euler.GetOrder() == EulerRotationOrderYZX {
		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[0] = c * e
		te[4] = bd - ac * f
		te[8] = bc * f + ad

		te[1] = f
		te[5] = a * e
		te[9] = -b * e

		te[2] = -d*e
		te[6] = ad * f + bc
		te[10] = ac - bd
	} else if euler.GetOrder() == EulerRotationOrderXZY {
		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[0] = c * e
		te[4] = -f
		te[8] = d * e

		te[1] = ac * f + bd
		te[5] = a * e
		te[9] = ad * f - bc

		te[2] = bc * f - ad
		te[6] = b * e
		te[10] = bd * f + ac
	}

	te[3] = 0
	te[7] = 0
	te[11] = 0

	te[12] = 0
	te[13] = 0
	te[14] = 0
	te[15] = 1
	
	return m
}
