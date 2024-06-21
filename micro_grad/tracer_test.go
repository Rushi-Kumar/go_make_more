package micrograd

import (
	"testing"
)

func TestTrace(t *testing.T) {
	x1 := NewValue(2.0)
	x2 := NewValue(0.0)
	w1 := NewValue(-3.0)
	w2 := NewValue(1.0)
	b := NewValue(6.8813735870195432)
	n := Add(Add(Mul(w1, x1), Mul(w2, x2)), b)
	e := Exp(Mul(NewValue(2), n))
	o := Div(Sub(e, NewValue(1)), Add(e, NewValue(1)))

	// o := Tanh(n)
	o.BackPropagate()
	Tracer(o)
}
