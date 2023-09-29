package pt

import (
	"fmt"

	"github.com/fogleman/gg"
)

type ObjectIface interface {
	// Accounting functions
	Init()
	GetName() string
	GetID() ObjectID
	GenID() error

	// Movement Functions (position, velocity, acceleration)
	GetPos() Vec
	SetPos(Vec)
	Move(Vec)
	GetPosOffset(ObjectIface) Vec
	GetVel() Vec
	SetVel(Vec)
	GetAcc() Vec
	SetAcc(Vec)
	GetKinematics() Kinematics
	SetKinematics(Kinematics)

	// Child functions
	GetChildren() []ObjectIface
	SetChildren([]ObjectIface)
	AddChildren(...ObjectIface)
	RemoveChildren(...ObjectIface)

	// Game loop functions
	Update() error
	Draw(drawContext *gg.Context)
}

func (o *Object) Init() {
	o.ObjInit()
}

func (o *Object) GetName() string {
	return o.ObjGetName()
}

func (o *Object) GetID() ObjectID {
	return o.ObjGetID()
}

func (o *Object) GenID() error {
	return o.ObjGenID()
}

// Get the velocity Vec of the object
func (o *Object) GetPos() Vec {
	return o.ObjGetPos()
}

// Get the velocity Vec of the object
func (o *Object) SetPos(pos Vec) {
	o.ObjSetPos(pos)
}

// Get the velocity Vec of the object
func (o *Object) GetVel() Vec {
	return o.ObjGetVel()
}

// Get the velocity Vec of the object
func (o *Object) SetVel(vel Vec) {
	o.ObjSetVel(vel)
}

// Get the velocity Vec of the object
func (o *Object) GetAcc() Vec {
	return o.ObjGetAcc()
}

// Get the velocity Vec of the object
func (o *Object) SetAcc(acc Vec) {
	o.ObjSetAcc(acc)
}

func (o *Object) GetChildren() []ObjectIface {
	return o.ObjGetChildren()
}

func (o *Object) SetChildren(children []ObjectIface) {
	o.ObjSetChildren(children)
}

func (o *Object) AddChildren(children ...ObjectIface) {
	o.ObjAddChildren(children...)
}

func (o *Object) RemoveChildren(children ...ObjectIface) {
	o.ObjRemoveChildren(children...)
}

func (o *Object) GetKinematics() Kinematics {
	return o.GetKinematics()
}

func (o *Object) SetKinematics(lk Kinematics) {
	o.SetKinematics(lk)
}

func (o *Object) Update() error {
	fmt.Println("Update for", o.GetName())
	return o.ObjUpdate()
}

// Intended to be implemented by the specific object
func (o *Object) Draw(drawContext *gg.Context) {
}
