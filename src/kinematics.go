package pt

type Kinematics struct {
	Pos Vec
	Vel Vec
	Acc Vec
}

func (o *Object) ObjGetPos() Vec {
	return o.LK.Pos
}

// Get the velocity Vec of the object
func (o *Object) ObjGetVel() Vec {
	return o.LK.Vel
}

// Get the velocity Vec of the object
func (o *Object) ObjGetAcc() Vec {
	return o.LK.Vel
}

// Set the position of the object and the relative position
// of all child objects
func (o *Object) ObjSetPos(pos Vec) {

	for _, child := range o.Children {
		child.SetPos(o.GetPosOffset(child).Add(pos))
	}
	o.LK.Pos = pos
}

// Set velocity of the object and of all child objects
func (o *Object) ObjSetVel(vel Vec) {
	o.LK.Vel = vel
	for _, child := range o.Children {
		child.SetVel(vel)
	}
}

// Set acceleration of the object and of all child objects
func (o *Object) ObjSetAcc(acc Vec) {
	o.LK.Acc = acc
	for _, child := range o.Children {
		child.SetAcc(acc)
	}
}

// Get the offset position vector of o1 as compared to o2
func (o1 *Object) GetPosOffset(o2 ObjectIface) Vec {
	return o2.GetPos().Subtract(o1.GetPos())
}

// Update this object's position using this object's velocity. Does not recursively call child objects
func (o *Object) UpdatePos() {
	o.LK.Pos = o.LK.Pos.Add(o.GetVel())
}

// Update this object's velocity using this object's acceleration. Does not recursively call child objects
func (o *Object) UpdateVel() {
	o.LK.Vel = o.LK.Vel.Add(o.GetAcc())
}

// Convenience function update velocity from acceleration
// and position from velocity
func (o *Object) UpdateKinematics() {
	o.UpdateVel()
	o.UpdatePos()
}

// Adjust the object's velocity according to the given vector v
func (o *Object) Accelerate(v Vec) {
	o.SetVel(o.GetVel().Add(v))
}

// Move an object relatively according to the given vector v.
func (o *Object) Move(v Vec) {
	o.SetPos(o.GetPos().Add(v))
}

func (o *Object) ObjGetKinematics() Kinematics {
	return o.LK
}

func (o *Object) ObjSetKinematics(lk Kinematics) {
	o.LK = lk
}
