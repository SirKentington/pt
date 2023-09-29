package pt

import (
	"reflect"
	"testing"
)

func newObjectWithChildren(name string, children ...ObjectIface) *Object {
	o := NewObject("Testing")
	o.AddChildren(children...)
	return o
}

func TestObject_GenID(t *testing.T) {
	tests := []struct {
		name    string
		o       *Object
		wantErr bool
	}{
		{o: NewObject("T"), wantErr: true},
		{o: &Object{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.GenID(); (err != nil) != tt.wantErr {
				t.Errorf("Object.GenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObject_GetChildren(t *testing.T) {

	children := []ObjectIface{}
	for i := 0; i < 4; i++ {
		children = append(children, NewObject("Child"))
	}

	tests := []struct {
		name string
		o    *Object
		want []ObjectIface
	}{
		{o: NewObject("T"), want: []ObjectIface{}},
		{o: newObjectWithChildren("Testing", children...), want: children},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByID(tt.o.GetChildren()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Object.GetChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_RemoveChildren(t *testing.T) {

	children := []ObjectIface{}
	for i := 0; i < 4; i++ {
		children = append(children, NewObject("Child"))
	}

	type args struct {
		children []ObjectIface
	}
	tests := []struct {
		name string
		o    *Object
		args args
		want []ObjectIface
	}{
		{o: newObjectWithChildren("Testing", children...), args: args{children}, want: []ObjectIface{}},
		{o: newObjectWithChildren("Testing", children...), args: args{children[0:1]}, want: children[1:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.RemoveChildren(tt.args.children...)
			got := tt.o.GetChildren()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Object.RemoveChildren().GetChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_Move(t *testing.T) {

	o1 := NewObject("Testing")
	o1.SetPos(Vec{123.4, 123.4, 123.4})

	type args struct {
		vel Vec
	}
	tests := []struct {
		name    string
		o       *Object
		args    args
		wantPos Vec
	}{
		{o: NewObject("Testing"), args: args{vel: Vec{1.0, 1.0, 0.0}}, wantPos: Vec{1.0, 1.0, 0.0}},
		{o: o1, args: args{vel: Vec{-123.4, 0, -100.0}}, wantPos: Vec{0.0, 123.4, 23.4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.Move(tt.args.vel)
			if tt.o.GetPos().Thousandths() != tt.wantPos.Thousandths() {
				t.Errorf("Got position %v but wanted position %v", tt.o.GetPos(), tt.wantPos)
			}
		})
	}
}

func TestObject_Accelerate(t *testing.T) {

	o1 := NewObject("Testing")
	o1.SetVel(Vec{1.0, 1.0, 1.0})

	type args struct {
		acc Vec
	}
	tests := []struct {
		name    string
		o       *Object
		args    args
		wantVel Vec
	}{
		{o: NewObject("Testing"), args: args{acc: Vec{10.0, 9.0, 8.0}}, wantVel: Vec{10.0, 9.0, 8.0}},
		{o: o1, args: args{acc: Vec{2.0, 2.0, 2.0}}, wantVel: Vec{3.0, 3.0, 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.Accelerate(tt.args.acc)
			if tt.o.GetVel().Thousandths() != tt.wantVel.Thousandths() {
				t.Errorf("Got velocity %v but wanted velocity %v", tt.o.GetVel(), tt.wantVel)
			}
		})
	}
}

func TestObject_Update(t *testing.T) {

	o1 := NewObject("Testing")
	o1.SetVel(Vec{1.0, 10.0, 100.0})

	o2 := NewObject("Testing")
	o2.SetPos(Vec{1000.0, 100.0, 10.0})
	o2.SetAcc(Vec{20.0, 21.0, 22.0})

	tests := []struct {
		name    string
		o       *Object
		wantErr bool
		wantPos Vec
		wantVel Vec
	}{
		{o: NewObject("Testing"), wantPos: Vec{}, wantVel: Vec{}},
		{o: o1, wantPos: Vec{1.0, 10.0, 100.0}, wantVel: Vec{1.0, 10.0, 100.0}},
		{o: o2, wantPos: Vec{1020.0, 121.0, 32.0}, wantVel: Vec{20.0, 21.0, 22.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.Update(); (err != nil) != tt.wantErr {
				t.Errorf("Object.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.o.GetVel().Thousandths() != tt.wantVel.Thousandths() {
				t.Errorf("Got velocity %v but wanted velocity %v", tt.o.GetVel(), tt.wantVel)
			}
			if tt.o.GetPos().Thousandths() != tt.wantPos.Thousandths() {
				t.Errorf("Got position %v but wanted position %v", tt.o.GetPos(), tt.wantPos)
			}
		})
	}
}
