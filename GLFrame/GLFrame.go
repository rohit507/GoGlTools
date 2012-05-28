
package GLFrame

import m3d "github.com/rohit507/GoGlTools/math3D"

type Frame struct {
	origin , forward , up m3d.Vector3f
}

func Init() *Frame {
	out := new(Frame)
	out.origin = m3d.NewV3f( 0.0 , 0.0 , 0.0 )
	out.up = m3d.NewV3f( 0.0 , 1.0 , 0.0 )
	out.forward = m3d.NewV3f( 0.0 ,  1.0 , -1.0 ) 
	return out
}

// Direct Origin Setters and getters

func (f *Frame) SetOriginV(point m3d.Vector3f) {
	f.origin.CopyFrom(point)
}

func (f *Frame) SetOriginL(x , y, z float32) {
	f.origin.Load(x,y,z)
}

func (f *Frame) GetOrigin() m3d.Vector3f {
	return m3d.NewV3f(f.origin[0] , f.origin[1] , f.origin[2])
}

// Direct Forward setters and getters

func (f *Frame) SetForwardV(point m3d.Vector3f) {
	f.forward.CopyFrom(point)
}

func (f *Frame) SetForwardL(x , y, z float32) {
	f.forward.Load(x,y,z)
}

func (f *Frame) GetForward() m3d.Vector3f {
	return m3d.NewV3f(f.forward[0] , f.forward[1] , f.forward[2])
}

// Direct Up setters and getters

func (f *Frame) SetUpV(point m3d.Vector3f) {
	f.up.CopyFrom(point)
}

func (f *Frame) SetUpL(x , y, z float32) {
	f.up.Load(x,y,z)
}

func (f *Frame) GetUp() m3d.Vector3f {
	return m3d.NewV3f(f.up[0] , f.up[1] , f.up[2])
}

