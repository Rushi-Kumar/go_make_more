package micrograd

import (
	"fmt"

	"github.com/m1gwings/treedrawer/tree"
)

func trace(value *Value, t *tree.Tree) {
	if len(value.children) == 0 {
		t.AddChild(tree.NodeString("d = " + fmt.Sprintf("%f", value.data) + "| g = " + fmt.Sprintf("%f", value.grad)))
		return
	}
	T := t.AddChild(tree.NodeString("d = " + fmt.Sprintf("%f", value.data) + "| g = " + fmt.Sprintf("%f", value.grad)))
	newT := T.AddChild(tree.NodeString(value.op))
	for _, child := range value.children {
		trace(child, newT)
	}
}

func Tracer(value *Value) {
	t := tree.NewTree(tree.NodeString("Root"))
	trace(value, t)
	fmt.Println(t)
}
