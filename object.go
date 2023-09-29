package pt

import (
	"fmt"
	"sort"

	"github.com/barkimedes/go-deepcopy"
	"github.com/fogleman/gg"
)

/****************************************
 * Functions implementing the interface *
 ***************************************/
// Generic object implementing the ObjectIface interface
type Object struct {
	Name       string
	ID         ObjectID
	LK         Kinematics // Structure for linear kinematics (angular kinematics may be added later)
	HB         Hitbox
	UpdateFunc func() error
	Children   map[ObjectID]ObjectIface // If this object contains any children, they go in this map
}

func (rdr *Object) CopyTo(wrtr *Object) {
	wrtr.Name = rdr.Name
	wrtr.LK = deepcopy.MustAnything(rdr.LK).(Kinematics)
	wrtr.HB = deepcopy.MustAnything(rdr.HB).(Hitbox)
	wrtr.Children = deepcopy.MustAnything(rdr.Children).(map[ObjectID]ObjectIface)
}

func (o *Object) ObjGetName() string {
	return o.Name
}

func NewObject(name string) *Object {
	o := &Object{Name: name}
	o.Init()
	return o
}

func NewObjAtPos(name string, pos Vec) *Object {
	o := NewObject(name)
	o.SetPos(pos)
	return o
}

func (o *Object) ObjInit() {
	o.GenID()
	o.Children = make(map[ObjectID]ObjectIface)
}

func (o *Object) ObjAddChildren(children ...ObjectIface) {
	for _, child := range children {
		o.Children[child.GetID()] = child
	}
}

func (o *Object) ObjRemoveChildren(children ...ObjectIface) {
	for _, child := range children {
		delete(o.Children, child.GetID())
	}
}

// Recursively
func (o *Object) ObjGetChildren() []ObjectIface {
	children := []ObjectIface{}
	for _, child := range o.Children {
		children = append(children, child)
		children = append(children, child.GetChildren()...)
	}
	return children
}

func (o *Object) ObjUpdate() error {
	o.UpdateKinematics()
	for _, child := range o.Children {
		child.Update()
	}
	return nil
}

func (o *Object) ObjSetChildren(children []ObjectIface) {
	o.Children = make(map[ObjectID]ObjectIface)
	o.AddChildren(children...)
}

/*************************************************************
 * Additional Helper Functions Using the Interface Functions *
 *************************************************************/

func sortByID(objects []ObjectIface) []ObjectIface {
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].GetID() < objects[j].GetID()
	})
	return objects
}

func (o *Object) ObjDrawPos(dc *gg.Context, at Vec) {
	pos := o.GetPos()
	x := fmt.Sprintf("%.0f", pos.X)
	y := fmt.Sprintf("%.0f", pos.Y)
	s := fmt.Sprintf("%s,%s", x, y)
	dc.SetRGB255(255, 255, 255)
	dc.DrawString(s, at.X, at.Y)
	dc.Stroke()
}
