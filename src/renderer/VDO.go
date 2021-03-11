package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// VertexDataObject ...
type VertexDataObject struct {
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices
	Shader        *Shader
	vao           uint32
	texture       *Texture
	elementsCount int32
	vboCount      uint32
}

// LoadVertexDataObject constructor for VertexDataObject
func LoadVertexDataObject(shader *Shader, texture *Texture) *VertexDataObject {
	var vdo VertexDataObject
	vdo.Shader = shader
	vdo.texture = texture
	gl.GenVertexArrays(1, &vdo.vao)
	return &vdo
}

// AddVBO ...
func (v *VertexDataObject) AddVBO(data []float32, size int32, drawMode uint32, datasetCount int) {
	v.bind()
	var buffer uint32

	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), drawMode)
	// stride := int32(3 * 4 * datasetCount)
	for i := 0; i < datasetCount; i++ {
		gl.EnableVertexAttribArray(v.vboCount)
		offset := i * 3 * 4
		gl.VertexAttribPointer(v.vboCount, size, gl.FLOAT, false, 0, gl.PtrOffset(offset))
		v.vboCount++
	}

	v.unbind()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// AddEBO ...
func (v *VertexDataObject) AddEBO(data []uint32, drawMode uint32) {
	v.elementsCount = int32(len(data))

	v.bind()
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data)*4, gl.Ptr(data), drawMode)

	v.unbind()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// Render ...
func (v *VertexDataObject) Render() {
	v.Shader.Use()
	v.texture.Bind()
	v.bind()
	gl.DrawElements(gl.TRIANGLES, v.elementsCount, gl.UNSIGNED_INT, nil)
}

// Bind VAO
func (v *VertexDataObject) bind() {
	gl.BindVertexArray(v.vao)
}

// Unbind VAO
func (v *VertexDataObject) unbind() {
	gl.BindVertexArray(0)
}

type BufferLayout struct {
}
