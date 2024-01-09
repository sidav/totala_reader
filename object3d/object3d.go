package object3d

import (
	"fmt"
	"strings"
	binaryreader "totala_reader/binary_reader"
)

type Object struct {
	// metadata from the file "header"
	VersionSignature           int
	NumberOfVertexes           int
	NumberOfPrimitives         int
	OffsetToselectionPrimitive int
	XFromParent                int
	YFromParent                int
	ZFromParent                int
	OffsetToObjectName         int
	always0                    int
	OffsetToVertexArray        int
	OffsetToPrimitiveArray     int
	OffsetToSiblingObject      int
	OffsetToChildObject        int

	// the object data itself
	ObjectName string
}

func (o *Object) Print(tabAmount int) {
	spaces := strings.Repeat(" ", tabAmount)
	str := fmt.Sprintf("%+#v", o)
	str = strings.Replace(str, " ", "\n"+spaces, -1)
	str = strings.Replace(str, "{", "{\n"+spaces, -1)
	str = strings.Replace(str, "}", "\n}"+spaces, -1)
	str = strings.Replace(str, ":", ": ", -1)
	fmt.Printf(spaces+"%s\n", str)
}

func ReadObjectFromReader(r *binaryreader.Reader, modelOffset int) *Object {
	obj := &Object{
		VersionSignature:           r.ReadIntFromBytesArray(modelOffset, 0),
		NumberOfVertexes:           r.ReadIntFromBytesArray(modelOffset, 4),
		NumberOfPrimitives:         r.ReadIntFromBytesArray(modelOffset, 8),
		OffsetToselectionPrimitive: r.ReadIntFromBytesArray(modelOffset, 12),
		XFromParent:                r.ReadIntFromBytesArray(modelOffset, 16),
		YFromParent:                r.ReadIntFromBytesArray(modelOffset, 20),
		ZFromParent:                r.ReadIntFromBytesArray(modelOffset, 24),
		OffsetToObjectName:         r.ReadIntFromBytesArray(modelOffset, 28),
		always0:                    r.ReadIntFromBytesArray(modelOffset, 32),
		OffsetToVertexArray:        r.ReadIntFromBytesArray(modelOffset, 36),
		OffsetToPrimitiveArray:     r.ReadIntFromBytesArray(modelOffset, 40),
		OffsetToSiblingObject:      r.ReadIntFromBytesArray(modelOffset, 44),
		OffsetToChildObject:        r.ReadIntFromBytesArray(modelOffset, 48),
	}

	obj.ObjectName = r.ReadNullTermStringFromBytesArray(0, obj.OffsetToObjectName)
	return obj
}
