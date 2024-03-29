package object3d

import binaryreader "totala_reader/ta_files_read"

type Vertex3d struct {
	x, y, z int
}

func (v *Vertex3d) ToFloats() (fx, fy, fz float64) {
	return FixedPointToFloat(v.x), FixedPointToFloat(v.y), FixedPointToFloat(v.z)
}

func ReadVertexesFromReader(r *binaryreader.Reader, vertexArrayOffset, vertexCount int) []Vertex3d {
	var vertexArray []Vertex3d
	for vInd := 0; vInd < vertexCount; vInd++ {
		currentVertexOffset := vertexArrayOffset + (vInd * 12)
		vertexArray = append(vertexArray, Vertex3d{
			r.ReadIntFromBytesArray(currentVertexOffset, 0),
			r.ReadIntFromBytesArray(currentVertexOffset, 4),
			r.ReadIntFromBytesArray(currentVertexOffset, 8),
		})
	}
	return vertexArray
}

// Bugged for negatives.
// func intToFixedPoint(x int) (integer, real int) {
// 	return x >> 16, x ^ (65535 << 16)
// }
