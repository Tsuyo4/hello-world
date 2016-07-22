package math

import "math"
import "math/rand"

const (
	DEG2RAD = math.Pi / 180
	RAD2DEG = 180 / math.Pi
)


func GenerateUUID() string {
	chars := `0123456789ABCDEFGHIJKLMNOPQQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	uuid  := make([]byte, 36)
	rnd   := 0
	r     := 0

	for i := 0 ; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			uuid[i] = '-';
		} else if i == 14 {
			uuid[i] = '4'
		} else {
			if rnd <= 0x02 {
				rnd = int( 0x2000000 + ( rand.Float64() * 0x1000000 ) ) | 0
			}
			r = rnd & 0xf
			rnd = rnd >> 4
			index  :=  r
			if i == 19 {
				index = ( r & 0x3 ) | 0x8
			}
			uuid[i] = chars[ index ]
		}
	}
	
	return string(uuid)
}

func Clamp( value float64, min float64, max float64) float64 {
	return math.Max( min, math.Min(max, value) )
}

func EuclideanModulo( n float64, m float64) float64 {
	return math.Mod(( math.Mod( n , m ) + m ) , m)
}

func MapLiner( x float64, a1 float64, a2 float64, b1 float64, b2 float64 ) float64 {
	return b1 + ( x - a1 ) * ( b2 - b1 ) / ( a2 - a1 )
}

func SmootherStemp( x float64, min float64, max float64) float64 {
	if x <= min {
		return 0
	}
	if x >= max {
		return 1
	}

	x = ( x - min ) / ( max - min )

	return x * x * x * ( x * ( x * 6 - 15 ) + 10 )
}

func RandInt( low int, high int ) int {
	return rand.Intn(high-low)+low
}


func RandFloat( low float64, high float64 ) float64 {
	return low+rand.Float64()*(high-low)
}

func RandFloatSpread( range_f float64) float64 {
	return range_f * ( 0.5 - rand.Float64() )
}

func DegToRad( degrees float64 ) float64 {
	return degrees * DEG2RAD
}

func RadToDeg( radians float64 ) float64 {
	return radians * RAD2DEG
}

func IsPowerOfTwo( value int ) bool {
	return ( value & ( value -1 ) ) == 0 && value != 0
}

func NearestPowerOfTwo( value float64 ) float64 {
	return math.Pow( 2, math.Floor( ( math.Log( value ) / math.Ln2 ) + .5 ) )
}

func NextPowerOfTow( value int ) int {
	value --
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	value++
	return value
}

func Round( f float64) float64 {
	return math.Floor(f + .5)
}



