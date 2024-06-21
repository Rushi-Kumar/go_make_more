package micrograd

type Layer struct {
	neurons []*Neuron
}

func NewLayer(in, out int, activation ActivationFunc) *Layer {
	neurons := make([]*Neuron, out)
	for i := 0; i < out; i++ {
		neurons[i] = NewNeuron(in, NewValue(0), activation)
	}
	return &Layer{neurons: neurons}
}

func (l *Layer) Forward(inputs []*Value) []*Value {
	out := make([]*Value, len(l.neurons))
	for i, neuron := range l.neurons {
		out[i] = neuron.Forward(inputs)
	}
	return out
}
