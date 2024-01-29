package model

import (
	"totala_reader/ta_files_read/object3d"
	"totala_reader/ta_files_read/texture"
)

// Non-TA format (floats everywhere) for ease of rendering
type Model struct {
	Vertices                              [][3]float64
	ObjectName                            string
	XFromParent, YFromParent, ZFromParent float64
	Primitives                            []*ModelSurface
	SelectionPrimitive                    *ModelSurface
	ChildObject                           *Model
	SiblingObject                         *Model
}

func NewModelFrom3doObject3d(obj *object3d.Object, allTextures []*texture.GafEntry) *Model {
	model := &Model{
		ObjectName: obj.ObjectName,
	}
	model.XFromParent, model.YFromParent, model.ZFromParent = obj.ParentOffsetAsFloats()
	for _, v := range obj.Vertexes {
		x, y, z := v.ToFloats()
		model.Vertices = append(model.Vertices, [3]float64{x, y, z})
	}
	for _, p := range obj.Primitives {
		newSurf := newModelSurfaceFrom3doPrimitive(p, allTextures)
		model.Primitives = append(model.Primitives, newSurf)
		// TODO: have selection primitive saved separately from other ones
		if p == obj.SelectionPrimitive {
			model.SelectionPrimitive = newSurf
		}
	}
	// Calculate and store centers for all surfaces
	model.calcCenterOfAllSurfaces()
	// Calculate UV-mapping
	model.performUvMappingOnAllSurfaces()
	if obj.ChildObject != nil {
		model.ChildObject = NewModelFrom3doObject3d(obj.ChildObject, allTextures)
	}
	if obj.SiblingObject != nil {
		model.SiblingObject = NewModelFrom3doObject3d(obj.SiblingObject, allTextures)
	}
	return model
}

func (m *Model) calcCenterOfAllSurfaces() {
	for _, surf := range m.Primitives {
		for _, index := range surf.VertexIndices {
			surf.CenterCoords[0] += m.Vertices[index][0]
			surf.CenterCoords[1] += m.Vertices[index][1]
			surf.CenterCoords[2] += m.Vertices[index][2]
		}
		for i := range surf.CenterCoords {
			surf.CenterCoords[i] /= float64(len(surf.VertexIndices))
		}
	}
}
