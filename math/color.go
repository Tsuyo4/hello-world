package math

import "fmt"
import "regexp"
import "strconv"
import "math"

type Color struct
{
	R float64
	G float64
	B float64
}

type HSL struct
{
	H float64
	S float64
	L float64
}

func NewColor( r float64, g float64, b float64) *Color {
	color := &Color{}

	color.R = r
	color.G = g
	color.B = b

	return color
}

func NewColorHex( value int32 ) *Color {
	color := &Color{}
	color.SetHex(value)

	return color
}

func NewColorString( value string ) *Color {
	color := &Color{}
	color.SetStyle( value )

	return color
}

func (color *Color) Set( value interface{} ) *Color {

	if v, ok := value.(*Color); ok {
		color.Copy(v)
	} else if v, ok := value.(int); ok {
		color.SetHex(int32(v))
	} else if v, ok := value.(string); ok {
		color.SetStyle(v)
	}

	return color
}


func (color *Color) Copy( color2 *Color ) *Color {
	color.R = color2.R
	color.G = color2.G
	color.B = color2.B

	return color
}

func (color *Color) SetHex( hex int32) *Color {

	color.R = float64( hex >> 16 & 255 ) / 255.0
	color.G = float64( hex >> 8  & 255 ) / 255.0
	color.B = float64( hex & 255 ) / 255.0

	return color
}

func (color *Color) SetRGB( r float64, g float64, b float64) *Color {
	color.R = r;
	color.G = g;
	color.B = b;

	return color
}

func hue2rgb( p float64, q float64, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0 / 6.0 {
		return p + ( q - p ) * 6.0 * t
	}
	if t < 1.0 / 2.0 {
		return q
	}
	if t < 2.0 / 3.0 {
		return p + ( q - p ) * 6.0 * ( 2.0 / 3.0 - t )
	}
	return p
}

func (color *Color) SetHSL( h float64, s float64, l float64) {
	h = EuclideanModulo(h, 1)
	s = Clamp(s, 0, 1)
	l = Clamp(l, 0, 1)

	if s == 0  {
		color.R = 1
		color.G = 1
		color.B = 1
	} else {
		p := l + s - ( l * s )
		if l <= 0.5 {
			p = l * ( 1 + s )
		}
		q := ( 2 * l ) - p
		color.R = hue2rgb( q, p, h + 1.0 / 3.0)
		color.G = hue2rgb( q, p, h)
		color.B = hue2rgb( q, p, h - 1.0 / 3.0)
	}
}

func (color *Color) GetHex() int32 {
	return int32(color.R * 255) << 16 ^ int32( color.G * 255 ) << 8 ^ int32( color.B * 255 ) << 0
}

func (color *Color) SetStyle( style string) {
	rgbhslReg := regexp.MustCompile(`^((?:rgb|hsl)a?)\(\s*([^\)]*)\)`)

	if rgbhslReg.MatchString( style ) {
		group      := rgbhslReg.FindStringSubmatch(style)
		name       := group[1]
		components := group[2]

		switch( name ) {
		case "rgb", "rgba":
			rgb1Reg := regexp.MustCompile(`^(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*(,\s*([0-9]*\.?[0-9]+)\s*)?$`)
			if rgb1Reg.MatchString( components ) {
				colorArray := rgb1Reg.FindStringSubmatch( components )
				r_value, _ := strconv.ParseInt( colorArray[1], 10, 32)
				color.R = math.Min( 255, float64(r_value) ) / 255
				g_value, _ := strconv.ParseInt( colorArray[2], 10, 32)
				color.G = math.Min( 255, float64(g_value) ) / 255
				b_value, _ := strconv.ParseInt( colorArray[3], 10, 32)
				color.B = math.Min( 255, float64(b_value) ) / 255

				return
			}
			
			rgb2Reg := regexp.MustCompile(`^(\d+)\%\s*,\s*(\d+)\%\s*,\s*(\d+)\%\s*(,\s*([0-9]*\.?[0-9]+)\s*)?$`)
			if rgb2Reg.MatchString( components ) {
				colorArray := rgb2Reg.FindStringSubmatch( components )
				r_value, _ := strconv.ParseInt( colorArray[1], 10, 32)
				color.R = math.Min( 100, float64(r_value) ) / 100.0
				g_value, _ := strconv.ParseInt( colorArray[2], 10, 32)
				color.G = math.Min( 100, float64(g_value) ) / 100.0
				b_value, _ := strconv.ParseInt( colorArray[3], 10, 32)
				color.B = math.Min( 100, float64(b_value) ) / 100.0

				return
			}
			break;

		case "hsl", "hsla":
			hslReg := regexp.MustCompile(`^([0-9]*\.?[0-9]+)\s*,\s*(\d+)\%\s*,\s*(\d+)\%\s*(,\s*([0-9]*\.?[0-9]+)\s*)?`)
			if hslReg.MatchString( components ) {
				colorArray := hslReg.FindStringSubmatch( components )
				h_value, _ := strconv.ParseFloat( colorArray[1], 64 )
				h := h_value / 360
				s_value, _ := strconv.ParseInt( colorArray[2], 10, 32)
				s := float64(s_value) / 100
				l_value, _ := strconv.ParseInt( colorArray[3], 10, 32)
				l := float64(l_value) / 100

				color.SetHSL(h, s, l);

				return
			}
			
			
		}
		
	} else {
		colorCodeReg := regexp.MustCompile(`^\#([A-Fa-f0-9]+)$`)
		if colorCodeReg.MatchString( style ) {
			colorArray := colorCodeReg.FindStringSubmatch( style )
			hex  := colorArray[1]
			size := len(hex)

			if( size == 3 ) {
				r_value, _ := strconv.ParseInt( string(hex[0]) + string(hex[0]), 16, 32 )
				color.R    = float64(r_value) / 255.0
				g_value, _ := strconv.ParseInt( string(hex[1]) + string(hex[1]), 16, 32 )
				color.G    = float64(g_value) / 255.0
				b_value, _ := strconv.ParseInt( string(hex[2]) + string(hex[2]), 16, 32 )
				color.B    = float64(b_value) / 255.0
				return
			} else if size == 6 {
				r_value, _ := strconv.ParseInt( string(hex[0]) + string(hex[1]), 16, 32 )
				color.R    = float64(r_value) / 255.0
				g_value, _ := strconv.ParseInt( string(hex[2]) + string(hex[3]), 16, 32 )
				color.G    = float64(g_value) / 255.0
				b_value, _ := strconv.ParseInt( string(hex[4]) + string(hex[5]), 16, 32 )
				color.B    = float64(b_value) / 255.0
				return
			}
		}
	}

	if len(style) > 0 {
		if ColorKeywords[style] != 0 {
			hex := ColorKeywords[style]

			color.SetHex(hex);
		}
		// TODO else 後のエラー処理

	}
	
}

func (color *Color) CopyGammaToLinear( color2 *Color, gammaFactor float64) *Color {
	color.R = math.Pow( color2.R, gammaFactor )
	color.G = math.Pow( color2.G, gammaFactor )
	color.B = math.Pow( color2.B, gammaFactor )

	return color
}

func (color *Color) CopyLinearToGamma( color2 *Color, gammaFactor float64) *Color {
	safeInverse := 1.0
	if gammaFactor > 0 {
		safeInverse = 1.0 / gammaFactor
	}

	color.R = math.Pow( color2.R, safeInverse)
	color.G = math.Pow( color2.G, safeInverse)
	color.B = math.Pow( color2.B, safeInverse)

	return color
}


func (color *Color) ConvertGammaToLinear() *Color {
	r := color.R
	g := color.G
	b := color.B

	color.R = r * r
	color.G = g * g
	color.B = b * b

	return color
}


func (color *Color) ConvertLinearToGamma() *Color {
	color.R = math.Sqrt(color.R)
	color.G = math.Sqrt(color.G)
	color.B = math.Sqrt(color.B)

	return color
}

func (color *Color) Clone() *Color {
	c := &Color{}
	c.Copy(color)

	return c
}

func (color *Color) Lerp( color2 *Color, alpha float64) *Color {
	color.R += ( color2.R - color.R ) * alpha
	color.G += ( color2.G - color.G ) * alpha
	color.B += ( color2.B - color.B ) * alpha

	return color
}

func (color *Color) GetHexString() string {
	return fmt.Sprintf("%06x", color.GetHex())
}

func (color *Color) GetStyle() string {
	return fmt.Sprintf("rgb(%d,%d,%d)",int32(color.R*255), int32(color.G*255), int32(color.B*255))
}

func (color *Color) GetHSL() *HSL {
	hsl := &HSL{}

	r := color.R
	g := color.G
	b := color.B

	max := math.Max( r, math.Max( g, b) )
	min := math.Min( r, math.Min( g, b) )

	hue        := 0.0
	saturation := 0.0
	lightness  := ( min + max ) / 2.0
	
	if min != max {
		delta := max - min

		saturation = delta / ( 2 - max - min )
		if lightness <= 0.5 {
			saturation = delta / ( max + min )
		}

		switch max {
		case r:
			hue = ( g - b )/  delta
			if g < b {
				hue += 6
			}
		case g:
			hue = ( b - r ) / delta + 2
		case b:
			hue = ( r - g ) / delta + 4
		}

		hue /= 6
	}

	hsl.H = hue
	hsl.S = saturation
	hsl.L = lightness

	return hsl
}

func (color *Color) OffsetHSL( h, s, l float64) *Color {

	var hsl = color.GetHSL()

	hsl.H += h
	hsl.S += s
	hsl.L += l

	color.SetHSL(hsl.H, hsl.S, hsl.L)

	return color
}


func (color *Color) Add( color2 *Color) *Color {
	color.R += color2.R
	color.G += color2.G
	color.B += color2.B

	return color
}

func (color *Color) AddColors( color2 , color3 *Color) *Color {
	color.R += color2.R + color3.R
	color.G += color2.G + color3.G
	color.B += color2.B + color3.B

	return color
}


func (color *Color) AddScalar( s float64) *Color {
	color.R += s
	color.G += s
	color.B += s

	return color
}

func (color *Color) Multiply( color2 *Color) *Color {
	color.R *= color2.R
	color.G *= color2.G
	color.B *= color2.B

	return color
}


func (color *Color) Equals( color2 *Color) bool {
	return (color.R == color2.R) && (color.G == color2.G) && (color.B == color2.B)
	
}

func (color *Color) FromArray( array [3]float64) {
	color.R = array[0]
	color.G = array[1]
	color.B = array[2]
}

func (color *Color) ToArray( array []float64, offset int ) []float64 {

	array[ offset ]   = color.R
	array[ offset+1 ] = color.G
	array[ offset+2 ] = color.B

	return array
}


var ColorKeywords map[string]int32 = map[string]int32 {"aliceblue": 0xF0F8FF, "antiquewhite": 0xFAEBD7, "aqua": 0x00FFFF, "aquamarine": 0x7FFFD4, "azure": 0xF0FFFF,
"beige": 0xF5F5DC, "bisque": 0xFFE4C4, "black": 0x000000, "blanchedalmond": 0xFFEBCD, "blue": 0x0000FF, "blueviolet": 0x8A2BE2,
"brown": 0xA52A2A, "burlywood": 0xDEB887, "cadetblue": 0x5F9EA0, "chartreuse": 0x7FFF00, "chocolate": 0xD2691E, "coral": 0xFF7F50,
"cornflowerblue": 0x6495ED, "cornsilk": 0xFFF8DC, "crimson": 0xDC143C, "cyan": 0x00FFFF, "darkblue": 0x00008B, "darkcyan": 0x008B8B,
"darkgoldenrod": 0xB8860B, "darkgray": 0xA9A9A9, "darkgreen": 0x006400, "darkgrey": 0xA9A9A9, "darkkhaki": 0xBDB76B, "darkmagenta": 0x8B008B,
"darkolivegreen": 0x556B2F, "darkorange": 0xFF8C00, "darkorchid": 0x9932CC, "darkred": 0x8B0000, "darksalmon": 0xE9967A, "darkseagreen": 0x8FBC8F,
"darkslateblue": 0x483D8B, "darkslategray": 0x2F4F4F, "darkslategrey": 0x2F4F4F, "darkturquoise": 0x00CED1, "darkviolet": 0x9400D3,
"deeppink": 0xFF1493, "deepskyblue": 0x00BFFF, "dimgray": 0x696969, "dimgrey": 0x696969, "dodgerblue": 0x1E90FF, "firebrick": 0xB22222,
"floralwhite": 0xFFFAF0, "forestgreen": 0x228B22, "fuchsia": 0xFF00FF, "gainsboro": 0xDCDCDC, "ghostwhite": 0xF8F8FF, "gold": 0xFFD700,
"goldenrod": 0xDAA520, "gray": 0x808080, "green": 0x008000, "greenyellow": 0xADFF2F, "grey": 0x808080, "honeydew": 0xF0FFF0, "hotpink": 0xFF69B4,
"indianred": 0xCD5C5C, "indigo": 0x4B0082, "ivory": 0xFFFFF0, "khaki": 0xF0E68C, "lavender": 0xE6E6FA, "lavenderblush": 0xFFF0F5, "lawngreen": 0x7CFC00,
"lemonchiffon": 0xFFFACD, "lightblue": 0xADD8E6, "lightcoral": 0xF08080, "lightcyan": 0xE0FFFF, "lightgoldenrodyellow": 0xFAFAD2, "lightgray": 0xD3D3D3,
"lightgreen": 0x90EE90, "lightgrey": 0xD3D3D3, "lightpink": 0xFFB6C1, "lightsalmon": 0xFFA07A, "lightseagreen": 0x20B2AA, "lightskyblue": 0x87CEFA,
"lightslategray": 0x778899, "lightslategrey": 0x778899, "lightsteelblue": 0xB0C4DE, "lightyellow": 0xFFFFE0, "lime": 0x00FF00, "limegreen": 0x32CD32,
"linen": 0xFAF0E6, "magenta": 0xFF00FF, "maroon": 0x800000, "mediumaquamarine": 0x66CDAA, "mediumblue": 0x0000CD, "mediumorchid": 0xBA55D3,
"mediumpurple": 0x9370DB, "mediumseagreen": 0x3CB371, "mediumslateblue": 0x7B68EE, "mediumspringgreen": 0x00FA9A, "mediumturquoise": 0x48D1CC,
"mediumvioletred": 0xC71585, "midnightblue": 0x191970, "mintcream": 0xF5FFFA, "mistyrose": 0xFFE4E1, "moccasin": 0xFFE4B5, "navajowhite": 0xFFDEAD,
"navy": 0x000080, "oldlace": 0xFDF5E6, "olive": 0x808000, "olivedrab": 0x6B8E23, "orange": 0xFFA500, "orangered": 0xFF4500, "orchid": 0xDA70D6,
"palegoldenrod": 0xEEE8AA, "palegreen": 0x98FB98, "paleturquoise": 0xAFEEEE, "palevioletred": 0xDB7093, "papayawhip": 0xFFEFD5, "peachpuff": 0xFFDAB9,
"peru": 0xCD853F, "pink": 0xFFC0CB, "plum": 0xDDA0DD, "powderblue": 0xB0E0E6, "purple": 0x800080, "red": 0xFF0000, "rosybrown": 0xBC8F8F,
"royalblue": 0x4169E1, "saddlebrown": 0x8B4513, "salmon": 0xFA8072, "sandybrown": 0xF4A460, "seagreen": 0x2E8B57, "seashell": 0xFFF5EE,
"sienna": 0xA0522D, "silver": 0xC0C0C0, "skyblue": 0x87CEEB, "slateblue": 0x6A5ACD, "slategray": 0x708090, "slategrey": 0x708090, "snow": 0xFFFAFA,
"springgreen": 0x00FF7F, "steelblue": 0x4682B4, "tan": 0xD2B48C, "teal": 0x008080, "thistle": 0xD8BFD8, "tomato": 0xFF6347, "turquoise": 0x40E0D0,
"violet": 0xEE82EE, "wheat": 0xF5DEB3, "white": 0xFFFFFF, "whitesmoke": 0xF5F5F5, "yellow": 0xFFFF00, "yellowgreen": 0x9ACD32 };




