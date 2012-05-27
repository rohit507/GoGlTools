/* Manual conversion of the math3d library from http://oglsuperbible5.googlecode.com/svn/trunk/Src/GLTools/ */

package math3D

import (
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


