
package GLFrustum

import m3d "github.com/rohit507/GoGlTools/math3D"
import glfrm "github.com/rohit507/GoGlTools/GLFrame"
import "math"

type Frustum struct {
	    // The projection matrix for this frustum
      	projMatrix m3d.Matrix44f   

        // Untransformed corners of the frustum
        nearUL, nearLL, nearUR, nearLR m3d.Vector4f
        farUL,  farLL,  farUR,  farLR m3d.Vector4f

        // Transformed corners of Frustum
        nearULT, nearLLT, nearURT, nearLRT m3d.Vector4f
        farULT,  farLLT,  farURT,  farLRT m3d.Vector4f

        // Base and Transformed plane equations
        nearPlane, farPlane, leftPlane, rightPlane m3d.Vector4f
        topPlane, bottomPlane m3d.Vector4f
}

func Init() *Frustum {
	f := new(Frustum)
	f.SetOrthographic(-1.0,1.0,-1.0,1.0,-1.0,1.0)
	return f
}

//get projection Matrix

func (f *Frustum) GetProjectionMatrix() *m3d.Matrix44f {
	return &f.projMatrix
}

// calculate the default values for various projection matrices

func (f *Frustum) SetOrthographic(xMin, xMax, yMin, yMax, zMin, zMax float32) {

	f.projMatrix = m3d.OrthographicMatrix44f(xMin, xMax, yMin, yMax, zMin, zMax)

	f.projMatrix[15] = 1.0

	// Fill in values for untransformed Frustum corners
    // Near Upper Left
	f.nearUL.Load(xMin,yMax,zMin,1.0)

    // Near Lower Left
	f.nearLL.Load(xMin,yMin,zMin,1.0)

    // Near Upper Right
	f.nearUR.Load(xMax,yMax,zMin,1.0)

    // Near Lower Right
	f.nearLR.Load(xMax,yMin,zMin,1.0)

    // Far Upper Left
	f.farUL.Load(xMin,yMax,zMax,1.0)

    // Far Lower Left
	f.farLL.Load(xMin,yMin,zMax,1.0)

    // Far Upper Right
	f.farUR.Load(xMax,yMax,zMax,1.0)

    // Far Lower Right
	f.farLR.Load(xMax,yMin,zMax,1.0)
}

func (f *Frustum) SetPerspective(fov , aspect, near, far float32) {

	ymax := near * float32(math.Tan(float64(fov * m3d.Pi / 360.0 )))
    ymin := -ymax
    xmin := ymin * aspect
    xmax := -xmin

	f.projMatrix = [16]float32{(2.0 * near)/(xmax - xmin),0.0,0.0,0.0,
							   0.0,(2.0 * near)/(ymax - ymin),0.0,0.0,
							   (xmax + xmin) / (xmax - xmin),
							   (ymax + ymin) / (ymax - ymin),
							   -((far + near)/(far - near)),
							   -1.0,0.0,0.0,-((2.0 * far * near)/(far - near)),
							   0.0 }

	yFmax := far * float32(math.Tan(float64(fov * m3d.Pi / 360.0)))
    yFmin := -yFmax
    xFmin := yFmin * aspect
    xFmax := -xFmin

	f.nearUL.Load(xmin,ymax,-near,1.0)
	f.nearLL.Load(xmin,ymin,-near,1.0)
	f.nearUR.Load(xmax,ymax,-near,1.0)
	f.nearLR.Load(xmax,ymin,-near,1.0)
	
	f.farUL.Load(xFmin,yFmax,-far,1.0)
	f.farLL.Load(xFmin,yFmin,-far,1.0)
	f.farUR.Load(xFmax,yFmax,-far,1.0)
	f.farLR.Load(xFmax,yFmin,-far,1.0)
}

// Builds a transformation matrix and transforms the corners of the Frustum,
// then derives the plane equations

func (f *Frustum) Transform(cam glfrm.Frame) {

	// Create the transformation matrix. This was the trickiest part
    // for me. The default view from OpenGL is down the negative Z
    // axis. However, building a transformation axis from these 
    // directional vectors points the frustum the wrong direction. So
    // You must reverse them here, or build the initial frustum
    // backwards - which to do is purely a matter of taste. I chose to
    // compensate here to allow better operability with some of my other
    // legacy code and projects. RSW

	forward := cam.GetForward()
	forward.ScaleSelf(-1.0)
	up := cam.GetUp()
	origin := cam.GetOrigin()

	cross := up.Cross(forward)

	rot := [16]float32{cross[0]   , cross[1]   , cross[2]   , 0.0 ,
					   up[0] 	  , up[1]      , up[2]      , 0.0 ,
					   forward[0] , forward[1] , forward[2] , 0.0 ,
					   origin[0]  , origin[1]  , origin[2]  , 1.0 }

	f.nearULT = f.nearUL.Transform(rot)	
	f.nearLLT = f.nearLL.Transform(rot)	
	f.nearURT = f.nearUR.Transform(rot)	
	f.nearLRT = f.nearLR.Transform(rot)
	
	f.farULT = f.farUL.Transform(rot)	
	f.farLLT = f.farLL.Transform(rot)	
	f.farURT = f.farUR.Transform(rot)	
	f.farLRT = f.farLR.Transform(rot)	

	f.nearPlane = m3d.GetPlaneEquation4f(f.nearULT,f.nearLLT,f.nearLRT)
	f.farPlane = m3d.GetPlaneEquation4f(f.farULT,f.farURT,f.farLRT)
	
	f.topPlane = m3d.GetPlaneEquation4f(f.nearULT,f.nearURT,f.farURT)
	f.bottomPlane = m3d.GetPlaneEquation4f(f.nearLLT,f.farLLT,f.farLRT)

	f.leftPlane = m3d.GetPlaneEquation4f(f.nearLLT,f.nearULT,f.farULT)
	f.rightPlane = m3d.GetPlaneEquation4f(f.nearLRT,f.farLRT,f.farURT)
}
