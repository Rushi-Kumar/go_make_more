/**
 * Micrograd is a tiny autograd engine, developed with reference to Andrej Karpathy's micrograd.
 *
**/

package micrograd

import (
	"testing"
)

func TestExpression(t *testing.T) {
	x := NewValue(3)
	y := NewValue(4)
	z := Mul(Add(x, y), NewValue(10)) // (x+y) * 10
	if z.data != 70 {
		t.Errorf("Expected 70, got %v", z.data)
	}

	x = Pow(NewValue(2), NewValue(3)) // 2^3
	y = NewValue(8)
	z = Sub(x, y) // 2^3 - 32
	if z.data != 0 {
		t.Errorf("Expected 0, got %v", x.data)
	}

	x = NewValue(3)
	y = NewValue(4)
	z = Add(x, y)
	if z.children[0] != x || z.children[1] != y {
		t.Errorf("Expected x and y to be children of z")
	}

	x = NewValue(3)
	y = NewValue(6)
	z = Div(y, x)
	if z.data != 2 {
		t.Errorf("Expected 2, got %v", z.data)
	}

	x = NewValue(3)
	y = NewValue(4)
	z = Sub(y, x)
	if z.data != 1 {
		t.Errorf("Expected 1, got %v", z.data)
	}
	Tracer(z)
}
