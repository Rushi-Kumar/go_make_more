/**
 * Micrograd is a tiny autograd engine, developed with reference to Andrej Karpathy's micrograd.
 *
**/

package micrograd

import (
	"math"
)

type Value struct {
	data     float64
	grad     float64
	op       string
	children []*Value
	backward func() // TODO implement
}

func NewValue(data float64, children ...*Value) *Value {
	return &Value{data: data, grad: 0, children: children, backward: nil}
}

func (v *Value) BackPropagate() {
	topolgicalOrder := make([]*Value, 0)
	visited := make(map[*Value]bool)
	var dfs func(*Value)
	dfs = func(v *Value) {
		if visited[v] {
			return
		}
		visited[v] = true
		for _, child := range v.children {
			dfs(child)
		}
		topolgicalOrder = append(topolgicalOrder, v)
	}
	dfs(v)
	v.grad = 1
	for i := len(topolgicalOrder) - 1; i >= 0; i-- {
		if topolgicalOrder[i].backward != nil {
			topolgicalOrder[i].backward()
		}
	}
}

func Add(x, y *Value) *Value {
	out := NewValue(x.data+y.data, x, y)
	backward := func() {
		x.grad += 1.0 * out.grad
		y.grad += 1.0 * out.grad
	}
	out.backward = backward
	out.op = "+"
	return out
}

func Mul(x, y *Value) *Value {
	out := NewValue(x.data*y.data, x, y)
	backward := func() {
		x.grad += y.data * out.grad
		y.grad += x.data * out.grad
	}
	out.backward = backward
	out.op = "*"
	return out
}

func Neg(x *Value) *Value {
	out := NewValue(-x.data, x)
	backward := func() {
		x.grad += -1.0 * out.grad
	}
	out.backward = backward
	out.op = "neg"
	return out
}

func Sub(x, y *Value) *Value {
	out := Add(x, Neg(y))
	return out
}

func Pow(x *Value, y *Value) *Value {
	out := NewValue(math.Pow(x.data, y.data), x, y)
	backward := func() {
		x.grad += y.data * math.Pow(x.data, y.data-1) * out.grad
	}
	out.backward = backward
	out.op = "pow"
	return out
}

func Exp(x *Value) *Value {
	out := NewValue(math.Exp(x.data), x)
	backward := func() {
		x.grad += out.grad * out.data
	}
	out.backward = backward
	out.op = "exp"
	return out
}

func Div(x, y *Value) *Value {
	out := Mul(x, Pow(y, NewValue(-1)))
	return out
}

func Tanh(x *Value) *Value {
	out := NewValue(math.Tanh(x.data), x)
	backward := func() {
		x.grad += x.grad + (1-out.data*out.data)*out.grad
	}
	out.backward = backward
	out.op = "tanh"
	return out
}
