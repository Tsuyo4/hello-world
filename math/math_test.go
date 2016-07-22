package math

import "testing"

func TestMath(test *testing.T) {

	if Clamp( 10, 0, 5 ) != 5 {
		test.Error(`Clamp Error`);
	}

}
