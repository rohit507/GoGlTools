
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

// get axis

func (f *Frame) GetX() m3d.Vector3f {
	return f.up.Cross(f.forward)
}

func (f *Frame) GetY() m3d.Vector3f {
	return f.GetUp()
}

func (f *Frame) GetZ() m3d.Vector3f {
	return f.GetForward() 
}

// movement relative to the current frame

func (f *Frame) MoveForward(dist float32) {
	f.origin.AddToSelf(f.forward.Scale(dist))
}	

func (f *Frame) MoveUp(dist float32) {
	f.origin.AddToSelf(f.up.Scale(dist))
}

func (f *Frame) MoveRight(dist float32) {
	cross := f.up.Cross(f.forward)
	cross.ScaleSelf(dist)
	f.origin.AddToSelf(cross)
}

func (f *Frame) TranslateLocalV(disp m3d.Vector3f) {
	f.MoveForward(disp[2]) 
	f.MoveUp(disp[1])
	f.MoveRight(disp[0])
}

func (f *Frame) TranslateLocalL( x, y, z float32) {
	f.MoveForward(z) 
	f.MoveUp(y)
	f.MoveRight(x)
}

//movement relative to the world reference frame

func (f *Frame) TranslateWorldV(disp m3d.Vector3f){
	f.origin.AddToSelf(disp)
}

func (f *Frame) TranslateWorldL( x , y , z float32){
	f.origin.AddToSelf([3]float32{x,y,z})
}

// matrix assembly or something

func (f *Frame) GetMatrix(rotationOnly bool) m3d.Matrix44f {
	x := f.up.Cross(f.forward)
	var out m3d.Matrix44f

	out.SetColumn3(x,0)
	out[3] = 0.0

	out.SetColumn3(f.up,1)
	out[7] = 0.0

	out.SetColumn3(f.forward,2)
	out[11] = 0.0

	if rotationOnly {
		out[12] = 0.0
		out[13] = 0.0
		out[14] = 0.0
	} else {
		out.SetColumn3(f.origin,3)
	}

	out[15] = 1.0

	return out
}	

// this is the assembly of the camera matrix

func (f *Frame) GetCameraMatrix(rotationOnly bool) m3d.Matrix44f {

	z := f.forward.Scale(-1.0)
	x := f.up.Cross(z)

	var out m3d.Matrix44f

	out.SetColumn3(x,0)
	out[3] = 0.0

	out.SetColumn3(f.up,1)
	out[7] = 0.0

	out.SetColumn3(z,2)
	out[11] = 0.0

	out[12] = 0.0
	out[13] = 0.0
	out[14] = 0.0
	out[15] = 1.0

	if rotationOnly {
		return out
	}

	trans := m3d.TranslationMatrix44f(- f.origin[0] , - f.origin[1] , - f.origin[2])
	out = out.Multiply(trans)


	return out
}

// local rotations

func (f *Frame) RotateLocalY(ang float32) {
	rot := m3d.RotationMatrix44f(ang , f.up[0] ,f.up[1] ,f.up[2])
	rotf := f.forward.Transform(rot)
	f.forward.CopyFrom(rotf)
}

func (f *Frame) RotateLocalZ(ang float32) {
	rot := m3d.RotationMatrix44f(ang , f.forward[0] , f.forward[1] , f.forward[2] )
	rotu := f.up.Transform(rot)
	f.up.CopyFrom(rotu)
}

func (f *Frame) RotateLocalX(ang float32) {
	x := f.up.Cross(f.forward)

	rot := m3d.RotationMatrix44f(ang , x[0] , x[1] , x[2])
	
	rotu := f.up.Transform(rot)
	rotf := f.forward.Transform(rot)

	f.up.CopyFrom(rotu)
	f.forward.CopyFrom(rotf)
}

func (f *Frame) RotateWorld(ang , x , y, z float32) {
	
	rot := m3d.RotationMatrix44f(ang , x , y , z)
	
	rotu := f.up.Transform(rot)
	rotf := f.forward.Transform(rot)

	f.up.CopyFrom(rotu)
	f.forward.CopyFrom(rotf)
}

// normalize the whole kit cat and kaboodle

func (f *Frame) Normalize() {

	c := f.up.Cross(f.forward)
	f.forward = c.Cross(f.up)	

	f.forward.Normalize()
	f.up.Normalize()

}

// convert coordinate systems

func (f *Frame) LocalToWorld(local m3d.Vector3f, rotOnly bool) m3d.Vector3f {

	rot := f.GetMatrix(true)

	out := local.Transform(rot)

	if ! rotOnly {
		out.AddToSelf(f.origin)
	}

	return out

}

func (f *Frame) WorldToLocal(world m3d.Vector3f) m3d.Vector3f {

	out := world.Sub(f.origin)

	rot := f.GetMatrix(true)

	rot = rot.Inverse()

	return out.Transform(rot)

}
	
func (f *Frame) TransformPoint( in m3d.Vector3f) m3d.Vector3f {

	rot := f.GetMatrix(false) 

	return in.Transform(rot)

}

func (f *Frame) RotateVector( in m3d.Vector3f) m3d.Vector3f {

	rot := f.GetMatrix(true)

	return in.Transform(rot)

}	
