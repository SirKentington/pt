package pt

import (
	"reflect"
	"testing"
)

func TestObject_GetPos(t *testing.T) {
	type fields struct {
		Name     string
		Pos      Vec
		ChildPos Vec
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{fields: fields{Pos: Vec{}}},
		{fields: fields{Pos: Vec{1, 1, 1}}},
		{fields: fields{Pos: Vec{1000000, 1000000, 1000000}}},
		{fields: fields{Pos: Vec{0.25, 0.5, 0.75}}},
		{fields: fields{Pos: Vec{}, ChildPos: Vec{1, 1, 1}}},
		{fields: fields{Pos: Vec{1, 1, 1}, ChildPos: Vec{-1, -1, -1}}},
		{fields: fields{Pos: Vec{1000000, 1000000, 1000000}, ChildPos: Vec{-100, -100, -100}}},
		{fields: fields{Pos: Vec{0.25, 0.5, 0.75}, ChildPos: Vec{-0.75, 0.5, -0.25}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewObject("parent")
			child := NewObjAtPos("child", tt.fields.ChildPos)
			o.AddChildren(child)

			// Set the object and its child to the new position
			o.SetPos(tt.fields.Pos)
			pos := o.GetPos()
			childPos := child.GetPos()
			if !reflect.DeepEqual(pos, tt.fields.Pos) {
				t.Errorf("Object.GetPos() is %v after SetPos instead of %v", pos, tt.fields.Pos)
			}
			if !reflect.DeepEqual(childPos.Subtract(tt.fields.Pos), tt.fields.ChildPos) {
			}

			// Now set the object and its child back
			newPos := Vec{}
			o.SetPos(newPos)
			pos = o.GetPos()
			childPos = child.GetPos()
			if !reflect.DeepEqual(pos, newPos) {
				t.Errorf("Object.GetPos() is %v after SetPos instead of %v", pos, newPos)
			}
			if !reflect.DeepEqual(childPos, tt.fields.ChildPos) {
				t.Errorf("Child.GetPos() is %v after SetPos instead of %v", childPos, tt.fields.ChildPos)
			}
		})
	}
}

func TestObject_SetPos(t *testing.T) {
	type fields struct {
		o ObjectIface
	}
	type args struct {
		pos Vec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{fields: fields{o: newBlock()}, args: args{pos: Vec{X: 3.5, Y: 4.5}}},
		{fields: fields{o: oWithChild()}, args: args{pos: Vec{X: 3.5, Y: 4.5}}},
		{fields: fields{o: doubleBlock()}, args: args{pos: Vec{X: 1.0, Y: 1.0}}},
		{fields: fields{o: letterT()}, args: args{pos: Vec{X: 3.5, Y: 4.5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.fields.o

			o.SetPos(tt.args.pos)
			oPos := o.GetPos()
			if !reflect.DeepEqual(tt.args.pos, oPos) {
				t.Errorf("Set pos to %v but GetPos gave %v", tt.args.pos, oPos)
			}
			origPos := everyPos(o)

			o.SetPos(tt.args.pos)
			oPos = o.GetPos()
			if !reflect.DeepEqual(tt.args.pos, oPos) {
				t.Errorf("Set pos to %v but GetPos gave %v", tt.args.pos, oPos)
			}
			newPos := everyPos(o)

			if !reflect.DeepEqual(newPos, origPos) {
				t.Errorf("Positions have changed from %v to %v", origPos, newPos)
			}
		})
	}
}
