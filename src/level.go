package pt

import (
	"sort"

	"github.com/fogleman/gg"
)

type Level struct {
	Object
}

func NewLevel() *Level {
	l := &Level{}
	l.Name = "Level"
	l.Init()
	return l
}

func (l *Level) ChildrenByZ() []ObjectIface {
	children := l.GetChildren()

	// Sort by ID first
	sort.Slice(children, func(i, j int) bool {
		return children[i].GetID() < children[j].GetID()
	})

	// Now stable sort by Z to keep the ID ordering
	sort.SliceStable(children, func(i, j int) bool {
		return children[i].GetPos().Z < children[j].GetPos().Z
	})
	return children
}

func (l *Level) Draw(drawContext *gg.Context) {
	for _, child := range l.ChildrenByZ() {
		child.Draw(drawContext)
	}
}

func (l *Level) Update() error {
	panic("Levels should always implement Update")
}

type LevelIface interface {
	ChildrenByZ() []ObjectIface
	Draw(drawContext *gg.Context)
	Update() error
}
