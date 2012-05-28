/* Manual conversion of the math3d library from http://oglsuperbible5.googlecode.com/svn/trunk/Src/GLTools/ */

package math3D

import (
	"math"
)

type (

	Vector2f [2]float32 	// 3D points = 3D Vectors, but we need
 	Vector2d [2]float64		// 2D representations sometimes... (x,y) order

	Vector3f [3]float32		// Vector of three floats (x, y, z)
	Vector3d [3]float64		// Vector of three doubles (x, y, z)

	Vector4f [4]float32		// Lesser used... Do we really need these?
	Vector4d [4]float64		// Yes, occasionaly we do need a trailing w component

	// 3x3 matrix - column major. X vector is 0, 1, 2, etc.
	//		0	3	6	
	//		1	4	7
	//		2	5	8

	Matrix33f [9]float32		// A 3 x 3 matrix, column major (floats) - OpenGL Style
	Matrix33d [9]float64		// A 3 x 3 matrix, column major (doubles) - OpenGL Style


	// 4x4 matrix - column major. X vector is 0, 1, 2, etc.
	//	0	4	8	12
	//	1	5	9	13
	//	2	6	10	14
	//	3	7	11	15

	Matrix44f	[16]float32		// A 4 X 4 matrix, column major (floats) - OpenGL style
	Matrix44d	[16]float64		// A 4 x 4 matrix, column major (doubles) - OpenGL style

)

const (
	Pi = math.Pi
	TwoPi = Pi * 2
	PiOver180 = Pi / 180
	InvPiOver180 = 1 / PiOver180
)

// Defined as macros in m3d.h, should probably convert to
// 	somthing similar once funtionality is added in the 
//  meantime i'll make iteasy for the compiler to smooch 
//  things together

func DegToRad(x float32) float32 { return x * PiOver180 }
func RadToDeg(x float32) float32 { return x * InvPiOver180 }

func HrToDeg(x float32)  float32 { return x * (1.0 / 15.0) }
func HrToRad(x float32)  float32 { return x * (1.0 / 15.0) * PiOver180 }

func DegToHr(x float32)  float32 { return x * 15.0 }
func RadtoHr(x float32)  float32 { return x * 15 * InvPiOver180 }

// here are the inline functions  from m3d.h

func IsPow2 (x uint32) bool {

	for x > 1 {
		x <<= 1
	}

	return x == 1
}

	//yay OOifying things :P
	// also code duplication beyond belief

func (v *Vector2f) X() float32 { return v[0] }
func (v *Vector2f) Y() float32 { return v[1] }

func (v *Vector2f) SetX(i float32) { v[0] = i }
func (v *Vector2f) SetY(i float32) { v[1] = i }

func (v *Vector2d) X() float64 { return v[0] }
func (v *Vector2d) Y() float64 { return v[1] }

func (v *Vector2d) SetX(i float64) { v[0] = i }
func (v *Vector2d) SetY(i float64) { v[1] = i }

func (v *Vector3f) X() float32 { return v[0] }
func (v *Vector3f) Y() float32 { return v[1] }
func (v *Vector3f) Z() float32 { return v[2] }

func (v *Vector3f) SetX(i float32) { v[0] = i }
func (v *Vector3f) SetY(i float32) { v[1] = i }
func (v *Vector3f) SetZ(i float32) { v[2] = i }

func (v *Vector3d) X() float64 { return v[0] }
func (v *Vector3d) Y() float64 { return v[1] }
func (v *Vector3d) Z() float64 { return v[2] }

func (v *Vector3d) SetX(i float64) { v[0] = i }
func (v *Vector3d) SetY(i float64) { v[1] = i }
func (v *Vector3d) SetZ(i float64) { v[2] = i }

func (v *Vector4f) X() float32 { return v[0] }
func (v *Vector4f) Y() float32 { return v[1] }
func (v *Vector4f) Z() float32 { return v[2] }
func (v *Vector4f) W() float32 { return v[3] }

func (v *Vector4f) SetX(i float32) { v[0] = i }
func (v *Vector4f) SetY(i float32) { v[1] = i }
func (v *Vector4f) SetZ(i float32) { v[2] = i }
func (v *Vector4f) SetW(i float32) { v[3] = i }

func (v *Vector4d) X() float64 { return v[0] }
func (v *Vector4d) Y() float64 { return v[1] }
func (v *Vector4d) Z() float64 { return v[2] }
func (v *Vector4d) W() float64 { return v[3] }

func (v *Vector4d) SetX(i float64) { v[0] = i }
func (v *Vector4d) SetY(i float64) { v[1] = i }
func (v *Vector4d) SetZ(i float64) { v[2] = i }
func (v *Vector4d) SetW(i float64) { v[3] = i }


