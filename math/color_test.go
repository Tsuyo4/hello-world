package math

import "testing"

func TestColorConstructor(test *testing.T) {
	c := &Color{}
	if c.R != 0 && c.G != 0 && c.B != 0 {
		test.Error("Color Constructor Error");
	}	
}

func TestColorConstructorFromRGB(test *testing.T) {
	c := NewColor(1, 1, 1)

	if c.R != 1 && c.G != 1 && c.B != 1 {
		test.Error("Color Constructor Error From RGB");
	}
}

func TestColorCopyHex(test *testing.T) {
	c  := &Color{}
	c2 := NewColorHex(0xF5FFFA);
	c.Copy(c2);
	if c.GetHex() != c2.GetHex() {
		test.Errorf("Hex c:%x Hex c2:%x", c.GetHex() , c2.GetHex())
	}
}

func TestColorCopyColorString(test *testing.T) {
	c  := &Color{}
	c2 := NewColorString("ivory")
	c.Copy(c2);
	if c.GetHex() != c2.GetHex() {
		test.Errorf("Hex c:%x Hex c2:%x", c.GetHex() , c2.GetHex())
	}
}

func TestSetRGB(test *testing.T) {
	c := &Color{}
	c.SetRGB(1, 0.2, 0.1)
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.2 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.1 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestCopyGammaToLinear(test *testing.T) {
	c  := &Color{}
	c2 := &Color{}
	c2.SetRGB( 0.3, 0.5, 0.9)
	c.CopyGammaToLinear(c2, 2.0)
	if c.R != 0.09 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.25 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.81 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestCopyLinearToGamma(test *testing.T) {
	c  := &Color{}
	c2 := &Color{}
	c2.SetRGB(0.09, 0.25, 0.81)
	c.CopyLinearToGamma(c2, 2.0)
	if c.R != 0.3 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.5 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.9 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestConvertGammaToLinear(test *testing.T) {
	c := &Color{}
	c.SetRGB(0.3, 0.5, 0.9)
	
	c.ConvertGammaToLinear();
	if c.R != 0.09 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.25 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.81 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestConvertLinearToGamma(test *testing.T) {
	c := &Color{}
	c.SetRGB(4, 9, 16)
	c.ConvertLinearToGamma()

	if c.R != 2 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 3 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 4 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetWithNum(test *testing.T) {
	c := &Color{}
	c.Set(0xFF0000)

	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetWithString(test *testing.T) {
	c := &Color{}
	c.Set("silver")

	if c.GetHex() != 0xC0C0C0 {
		test.Errorf("Hex C:", c.GetHex())
	}
	
}

func TestClone(test *testing.T) {
	c  := NewColorString("teal")
	c2 := c.Clone()

	if c2.GetHex() != 0x008080 {
		test.Errorf("Hex C2:%x", c2.GetHex())
	}
}

func TestLerp(test *testing.T) {
	c  := &Color{}
	c2 := &Color{1, 1, 1}

	c.SetRGB(0, 0, 0)
	c.Lerp(c2, 0.2)
	if c.R != 0.2 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.2 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.2 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBRed(test *testing.T) {
	c := Color{}
	c.SetStyle("rgb(255,0,0)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBARed(test *testing.T) {
	c := Color{}
	c.SetStyle("rgba(255,0,0,0.5)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBRedWithSpaces(test *testing.T) {
	c := Color{}
	c.SetStyle("rgb( 255,  0,    0 )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBARedWithSpace(test *testing.T) {
	c := Color{}
	c.SetStyle("rgba( 255,  0, 0   , 1 )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBPercent(test *testing.T) {
	c := Color{}
	c.SetStyle("rgb( 100%, 50%, 10%)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.5 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.1 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBAPercent(test *testing.T) {
	c := Color{}
	c.SetStyle("rgba( 100%, 50%, 10%, 0.5)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.5 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.1 {
		test.Errorf("Blue: %d", c.B)
	}
}


func TestStyleRGBPercentWithSpace(test *testing.T) {
	c := Color{}
	c.SetStyle("rgb( 100% ,50%   ,      10% )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.5 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.1 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestStyleRGBAPercentWithSpace(test *testing.T) {
	c := Color{}
	c.SetStyle("rgba( 100% ,50%   ,   10%, 0.5 )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0.5 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0.1 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetStyleHSLRed(test *testing.T) {
	c := Color{}
	c.SetStyle("hsl(360,100%,50%)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetStyleHSLARed(test *testing.T) {
	c := Color{}
	c.SetStyle("hsla(360,100%,50%,0.5)")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetStyleHSLRedSpaces(test *testing.T) {
	c := Color{}
	c.SetStyle("hsl(360,  100% ,  50% )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetStyleHSLARedSpaces(test *testing.T) {
	c := Color{}
	c.SetStyle("hsla(360,  100% ,  50%,   0.5 )")
	if c.R != 1 {
		test.Errorf("Red: %d", c.R)
	}
	if c.G != 0 {
		test.Errorf("Green: %d", c.G)
	}
	if c.B != 0 {
		test.Errorf("Blue: %d", c.B)
	}
}

func TestSetStyleHexSkyBlue(test *testing.T) {
	c := Color{}
	c.SetStyle("#87CEEB")

	if c.GetHex() != 0x87CEEB {
		test.Errorf("Hex c:%x",c.GetHex())
	}
	
}


func TestSetStyleHexSkyBlueMiexed(test *testing.T) {
	c := &Color{}
	c.SetStyle("#87cEeB")
	
	if c.GetHex() != 0x87CEEB {
		test.Errorf("Hex c:%x",c.GetHex())
	}
}

func TestStyleHEx2Olive(test *testing.T) {
	c := &Color{}
	c.SetStyle("#F00")

	if c.GetHex() != 0xFF0000 {
		test.Errorf("Hex c:%x",c.GetHex())
	}
}

func TestSetStyleHex2OliveMixed(test *testing.T) {
	c := &Color{}
	c.SetStyle("#f00")

	if c.GetHex() != 0xFF0000 {
		test.Errorf("Hex c:%x",c.GetHex())
	}
}

func TestSetStyleColorName(test *testing.T) {
	c := &Color{}
	c.SetStyle("powderblue")

	if c.GetHex() != 0xB0E0E6 {
		test.Errorf("Hex c:%x",c.GetHex())
	}
}

func TestGetHex(test *testing.T) {
	c   := NewColorString("red")
	res := c.GetHex()
	if res != 0xFF0000 {
		test.Errorf("Hex c:%x",res)
	}
}

func TestSetHex(test *testing.T) {
	c := &Color{}
	c.SetHex(0xFA8072)
	if c.GetHex() != 0xFA8072 {
		test.Errorf("Hex c:%x",c.GetHex())
	}
}

func TestGetHexString(test *testing.T) {
	c   := NewColorString("tomato")
	res := c.GetHexString()

	if res != "ff6347" {
		test.Errorf("Hex :%s",res)
	}
}

func TestGetStyle(test *testing.T) {
	c   := NewColorString("plum")
	res := c.GetStyle()
	if res != "rgb(221,160,221)" {
		test.Errorf("Style:%s",res)
	}
}

func TestGetHSL(test *testing.T) {
	c   := NewColorHex( 0x80FFFF )
	hsl := c.GetHSL()

	if hsl.H != 0.5 {
		test.Errorf("hue: %d", hsl.H)
	}

	if hsl.S != 1.0 {
		test.Errorf("saturation: %d", hsl.S)
	}

	if Round(hsl.L*100)/100  != 0.75 {
		test.Errorf("lightness: %d", hsl.L)
	}
}

func TestSetHSL(test *testing.T) {
	c := &Color{}
	c.SetHSL(0.75, 1.0, 0.25)
	hsl := c.GetHSL()

	if hsl.H != 0.75 {
		test.Errorf("hue: %d", hsl.H)
	}

	if hsl.S != 1.0 {
		test.Errorf("saturation: %d", hsl.S)
	}

	if hsl.L != 0.25 {
		test.Errorf("lightness: %d", hsl.L)
	}
}
