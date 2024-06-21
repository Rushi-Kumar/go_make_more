package micrograd

import "math/rand"

type ActivationFunc func(*Value) *Value

type Neuron struct {
	weights    []*Value
	bias       *Value
	activation ActivationFunc
}

func NewNeuron(shape int, bias *Value, activationFunc ActivationFunc) *Neuron {
	weights := make([]*Value, shape)
	for i := 0; i < shape; i++ {
		weights[i] = NewValue(rand.Float64())
	}
	return &Neuron{weights: weights, bias: bias, activation: activationFunc}
}

func (n *Neuron) Forward(inputs []*Value) *Value {
	out := n.bias
	for i := 0; i < len(n.weights); i++ {
		out = Add(out, Mul(n.weights[i], inputs[i]))
	}
	return out
}
