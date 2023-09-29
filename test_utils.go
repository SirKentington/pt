package pt

type block struct {
	Object
}

func newBlock() *block {
	b := &block{}
	b.Name = "Block"
	b.Init()
	return b
}

func blockAt(pos Vec) *block {
	b := newBlock()
	b.SetPos(pos)
	return b
}

func doubleBlock() ObjectIface {
	o := NewObject("DoubleBlock")
	o.AddChildren(
		blockAt(Vec{}),
		blockAt(Vec{X: 1, Y: 1}),
	)
	return o
}

func oWithChild() ObjectIface {
	o := NewObject("ObjectWithChild")
	o.AddChildren(
		blockAt(Vec{X: 0.5, Y: 0.5}),
	)
	return o
}

func letterT() ObjectIface {
	letterT := NewObject("Letter T")
	letterT.AddChildren(
		blockAt(Vec{Y: 0, X: 0}),
		blockAt(Vec{Y: 0, X: -1}),
		blockAt(Vec{Y: 0, X: 1}),
		blockAt(Vec{Y: 1, X: 0}),
		blockAt(Vec{Y: 2, X: 0}),
		blockAt(Vec{Y: 3, X: 0}),
		blockAt(Vec{Y: 4, X: 0}),
	)
	return letterT
}

func everyPos(o ObjectIface) map[ObjectID]Vec {
	objs := make(map[ObjectID]Vec)
	objs[o.GetID()] = o.GetPos()

	for _, child := range o.GetChildren() {
		objs[child.GetID()] = child.GetPos()
	}
	return objs
}
