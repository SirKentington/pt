package pt

import (
	"fmt"
	"sync/atomic"
)

type ObjectID uint64

// Object IDs are monotonically increasing integers starting at 1
var objID uint64

func UniqueObjectID() ObjectID {
	return ObjectID(atomic.AddUint64(&objID, 1))
}

// If the object's ID is unset (ie. zero), generate a new unique object ID
func (o *Object) ObjGenID() error {
	if o.GetID() != 0 {
		return fmt.Errorf("could not set already set ID")
	}
	o.ID = UniqueObjectID()
	return nil
}

func (o *Object) ObjGetID() ObjectID {
	return o.ID
}
