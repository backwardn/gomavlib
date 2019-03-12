// autogenerated with dialgen. do not edit.

package icarous

import (
	"github.com/gswly/gomavlib"
)

// icarous.xml

type MessageIcarousHeartbeat struct {
	Status uint8
}

func (*MessageIcarousHeartbeat) GetId() uint32 {
	return 42000
}

type MessageIcarousKinematicBands struct {
	Numbands int8
	Type1    uint8
	Min1     float32
	Max1     float32
	Type2    uint8
	Min2     float32
	Max2     float32
	Type3    uint8
	Min3     float32
	Max3     float32
	Type4    uint8
	Min4     float32
	Max4     float32
	Type5    uint8
	Min5     float32
	Max5     float32
}

func (*MessageIcarousKinematicBands) GetId() uint32 {
	return 42001
}

var Dialect = []gomavlib.Message{
	// icarous.xml
	&MessageIcarousHeartbeat{},
	&MessageIcarousKinematicBands{},
}
