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

func CloseEnoughf(a,b, epsilon float32) bool {
	return float32(math.Abs(float64(a - b))) < epsilon
}

func CloseEnoughd(a,b , epsilon float64) bool {
	return math.Abs(a - b) < epsilon
}

// and here are the other utility functions from m3d.cpp

func ProjectXYf(modelView , projection Matrix44f , viewport [4]int, pIn Vector3f) Vector2f {

	var forth, back Vector4f	

	back = [4]float32{pIn[0],pIn[1],pIn[2],1.0}

	forth = back.Transform(modelView)
	back  = forth.Transform(projection)

	if CloseEnoughf(back[3], 0.0, 0.000001) {
		div := 1.0 / back[3]
		back[0] *= div
		back[1] *= div
	}

	out := [2]float32{float32(viewport[0])+(1.0+float32(back[0]))*float32(viewport[2])/2.0 ,
					  float32(viewport[1])+(1.0+float32(back[1]))*float32(viewport[3])/2.0}

	if viewport[0] != 0 {
		out[0] -= float32(viewport[0])
	}

	if viewport[1] != 0 {
		out[1] -= float32(viewport[1])
	}

	return out
}

func ProjectXYd(modelView , projection Matrix44d , viewport [4]int, pIn Vector3d) Vector2d {

	var forth, back Vector4d	

	back = [4]float64{pIn[0],pIn[1],pIn[2],1.0}

	forth = back.Transform(modelView)
	back  = forth.Transform(projection)

	if CloseEnoughd(back[3], 0.0, 0.000001) {
		div := 1.0 / back[3]
		back[0] *= div
		back[1] *= div
	}

	out := [2]float64{float64(viewport[0])+(1.0+float64(back[0]))*float64(viewport[2])/2.0 ,
					  float64(viewport[1])+(1.0+float64(back[1]))*float64(viewport[3])/2.0}

	if viewport[0] != 0 {
		out[0] -= float64(viewport[0])
	}

	if viewport[1] != 0 {
		out[1] -= float64(viewport[1])
	}

	return out
}

func ProjectXYZf(modelView , projection Matrix44f , viewport [4]int, pIn Vector3f) Vector3f {

	var forth, back Vector4f	

	back = [4]float32{pIn[0],pIn[1],pIn[2],1.0}

	forth = back.Transform(modelView)
	back  = forth.Transform(projection)

	if CloseEnoughf(back[3], 0.0, 0.000001) {
		div := 1.0 / back[3]
		back[0] *= div
		back[1] *= div
		back[2] *= div
	}

	out := [3]float32{float32(viewport[0])+(1.0+float32(back[0]))*float32(viewport[2])/2.0 ,
					  float32(viewport[1])+(1.0+float32(back[1]))*float32(viewport[3])/2.0 , 0.0}

	if viewport[0] != 0 {
		out[0] -= float32(viewport[0])
	}

	if viewport[1] != 0 {
		out[1] -= float32(viewport[1])
	}

	out[2] = back[2]

	return out
}

func ProjectXYZd(modelView , projection Matrix44d , viewport [4]int, pIn Vector3d) Vector3d {

	var forth, back Vector4d

	back = [4]float64{pIn[0],pIn[1],pIn[2],1.0}

	forth = back.Transform(modelView)
	back  = forth.Transform(projection)

	if CloseEnoughd(back[3], 0.0, 0.000001) {
		div := 1.0 / back[3]
		back[0] *= div
		back[1] *= div
		back[2] *= div
	}

	out := [3]float64{float64(viewport[0])+(1.0+float64(back[0]))*float64(viewport[2])/2.0 ,
					  float64(viewport[1])+(1.0+float64(back[1]))*float64(viewport[3])/2.0 , 0.0}

	if viewport[0] != 0 {
		out[0] -= float64(viewport[0])
	}

	if viewport[1] != 0 {
		out[1] -= float64(viewport[1])
	}

	out[2] = back[2]

	return out
}

// Calculates the normal of a triangle specified by the three points
// p1, p2, and p3. Each pointer points to an array of three floats. The
// triangle is assumed to be wound counter clockwise. 

func FindNormal3f(p1, p2, p3 Vector3f) Vector3f {
	var v1,v2 Vector3f;		// Temporary vectors

	// Calculate two vectors from the three points. Assumes counter clockwise
	// winding!
	v1[0] = p1[0] - p2[0];
	v1[1] = p1[1] - p2[1];
	v1[2] = p1[2] - p2[2];

	v2[0] = p2[0] - p3[0];
	v2[1] = p2[1] - p3[1];
	v2[2] = p2[2] - p3[2];

	// Take the cross product of the two vectors to get
	// the normal vector.
	return v1.Cross(v2)
}

func FindNormal3d(p1, p2, p3 Vector3d) Vector3d {
	var v1,v2 Vector3d;		// Temporary vectors

	// Calculate two vectors from the three points. Assumes counter clockwise
	// winding!
	v1[0] = p1[0] - p2[0];
	v1[1] = p1[1] - p2[1];
	v1[2] = p1[2] - p2[2];

	v2[0] = p2[0] - p3[0];
	v2[1] = p2[1] - p3[1];
	v2[2] = p2[2] - p3[2];

	// Take the cross product of the two vectors to get
	// the normal vector.
	return v1.Cross(v2)
}

// Calculate the plane equation of the plane that the three specified points lay in. The
// points are given in clockwise winding order, with normal pointing out of clockwise face
// planeEq contains the A,B,C, and D of the plane equation coefficients

func GetPlaneEquation3f(p1, p2, p3 Vector3f) Vector4f {
    // Get two vectors... do the cross product
    var v1, v2 Vector3f

    // V1 = p3 - p1
    v1[0] = p3[0] - p1[0];
    v1[1] = p3[1] - p1[1];
    v1[2] = p3[2] - p1[2];

    // V2 = P2 - p1
    v2[0] = p2[0] - p1[0];
    v2[1] = p2[1] - p1[1];
    v2[2] = p2[2] - p1[2];

    // Unit normal to plane - Not sure which is the best way here
    plane :=  v1.Cross(v2)
    plane.Normalize()

    // Back substitute to get D
    return [4]float32{plane[0],plane[1],plane[2],
			 -(plane[0] * p3[0] + plane[1] * p3[1] + plane[2] * p3[2])}
}

func GetPlaneEquation4f(p1, p2, p3 Vector4f) Vector4f {
    // Get two vectors... do the cross product
    var v1, v2 Vector3f

    // V1 = p3 - p1
    v1[0] = p3[0] - p1[0];
    v1[1] = p3[1] - p1[1];
    v1[2] = p3[2] - p1[2];

    // V2 = P2 - p1
    v2[0] = p2[0] - p1[0];
    v2[1] = p2[1] - p1[1];
    v2[2] = p2[2] - p1[2];

    // Unit normal to plane - Not sure which is the best way here
    plane :=  v1.Cross(v2)
    plane.Normalize()

    // Back substitute to get D
    return [4]float32{plane[0],plane[1],plane[2],
			 -(plane[0] * p3[0] + plane[1] * p3[1] + plane[2] * p3[2])}
}

func GetPlaneEquation3d(p1, p2, p3 Vector3d) Vector4d {
    // Get two vectors... do the cross product
    var v1, v2 Vector3d

    // V1 = p3 - p1
    v1[0] = p3[0] - p1[0];
    v1[1] = p3[1] - p1[1];
    v1[2] = p3[2] - p1[2];

    // V2 = P2 - p1
    v2[0] = p2[0] - p1[0];
    v2[1] = p2[1] - p1[1];
    v2[2] = p2[2] - p1[2];

    // Unit normal to plane - Not sure which is the best way here
    plane :=  v1.Cross(v2)
    plane.Normalize()

    // Back substitute to get D
    return [4]float64{plane[0],plane[1],plane[2],
			 -(plane[0] * p3[0] + plane[1] * p3[1] + plane[2] * p3[2])}
}

// This function does a three dimensional Catmull-Rom curve interpolation. Pass four points, and a
// floating point number between 0.0 and 1.0. The curve is interpolated between the middle two points.
// Coded by RSW

func CatmullRom3f(vP0, vP1, vP2, vP3 Vector3f, t float32) Vector3f {
	t2 := t * t
	t3 := t2 * t

	return [3]float32{ 0.5 * ( ( 2.0 * vP1[0]) + (-vP0[0] + vP2[0]) * t +
                       (2.0 * vP0[0] - 5.0 *vP1[0] + 4.0 * vP2[0] - vP3[0]) * t2 +
                       (-vP0[0] + 3.0*vP1[0] - 3.0 *vP2[0] + vP3[0]) * t3),
					   0.5 * ( ( 2.0 * vP1[1]) + (-vP0[1] + vP2[1]) * t +
                       (2.0 * vP0[1] - 5.0 *vP1[1] + 4.0 * vP2[1] - vP3[1]) * t2 +
                       (-vP0[1] + 3.0*vP1[1] - 3.0 *vP2[1] + vP3[1]) * t3),
					   0.5 * ( ( 2.0 * vP1[2]) +(-vP0[2] + vP2[2]) * t +
                       (2.0 * vP0[2] - 5.0 *vP1[2] + 4.0 * vP2[2] - vP3[2]) * t2 +
                       (-vP0[2] + 3.0*vP1[2] - 3.0 *vP2[2] + vP3[2]) * t3)}
}

func CatmullRom3d(vP0, vP1, vP2, vP3 Vector3d, t float64) Vector3d {
	t2 := t * t
	t3 := t2 * t

	return [3]float64{ 0.5 * ( ( 2.0 * vP1[0]) + (-vP0[0] + vP2[0]) * t +
                       (2.0 * vP0[0] - 5.0 *vP1[0] + 4.0 * vP2[0] - vP3[0]) * t2 +
                       (-vP0[0] + 3.0*vP1[0] - 3.0 *vP2[0] + vP3[0]) * t3),
					   0.5 * ( ( 2.0 * vP1[1]) + (-vP0[1] + vP2[1]) * t +
                       (2.0 * vP0[1] - 5.0 *vP1[1] + 4.0 * vP2[1] - vP3[1]) * t2 +
                       (-vP0[1] + 3.0*vP1[1] - 3.0 *vP2[1] + vP3[1]) * t3),
					   0.5 * ( ( 2.0 * vP1[2]) +(-vP0[2] + vP2[2]) * t +
                       (2.0 * vP0[2] - 5.0 *vP1[2] + 4.0 * vP2[2] - vP3[2]) * t2 +
                       (-vP0[2] + 3.0*vP1[2] - 3.0 *vP2[2] + vP3[2]) * t3)}
}

///////////////////////////////////////////////////////////////////////////////
// Determine if the ray (starting at point) intersects the sphere centered at
// sphereCenter with radius sphereRadius
// Return value is < 0 if the ray does not intersect
// Return value is 0.0 if ray is tangent
// Positive value is distance to the intersection point
// Algorithm from "3D Math Primer for Graphics and Game Development"

func RaySphereTest3f(point, ray, sphereCenter Vector3f, sphereRadius float32) (intersect bool,dist float32) {

	//m3dNormalizeVector(ray);	// Make sure ray is unit length

	var rayToCenter Vector3f	// Ray to center of sphere
	rayToCenter[0] =  sphereCenter[0] - point[0]	
	rayToCenter[1] =  sphereCenter[1] - point[1]
	rayToCenter[2] =  sphereCenter[2] - point[2]
	
	// Project rayToCenter on ray to test
	a := rayToCenter.Dot(ray)
	
	// Distance to center of sphere
	distance2 := rayToCenter.Dot(rayToCenter)	// Or length

	
	dRet := (sphereRadius * sphereRadius) - distance2 + (a*a)
	
	if (dRet > 0.0) {			// Return distance to intersection
		dRet = a - float32(math.Sqrt(float64(dRet)))
	}
	
	return dRet <= 0 , dRet
}

func RaySphereTest3d(point, ray, sphereCenter Vector3d, sphereRadius float64) (intersect bool,dist float64) {

	//m3dNormalizeVector(ray);	// Make sure ray is unit length

	var rayToCenter Vector3d	// Ray to center of sphere
	rayToCenter[0] =  sphereCenter[0] - point[0]	
	rayToCenter[1] =  sphereCenter[1] - point[1]
	rayToCenter[2] =  sphereCenter[2] - point[2]
	
	// Project rayToCenter on ray to test
	a := rayToCenter.Dot(ray)
	
	// Distance to center of sphere
	distance2 := rayToCenter.Dot(rayToCenter)	// Or length

	
	dRet := (sphereRadius * sphereRadius) - distance2 + (a*a)
	
	if (dRet > 0.0) {			// Return distance to intersection
		dRet = a - math.Sqrt(dRet)
	}
	
	return dRet <= 0 , dRet
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// Calculate the tangent basis for a triangle on the surface of a model
// This vector is needed for most normal mapping shaders 

func CalculateTangentBasis3f(vTriangle [3]Vector3f, vTexCoords [3]Vector2f, N Vector3f) Vector3f {

    dv2v1 := vTriangle[1].Sub(vTriangle[0])
    dv3v1 := vTriangle[2].Sub(vTriangle[0])
    
    dc2c1t := vTexCoords[1][0] - vTexCoords[0][0];
    dc2c1b := vTexCoords[1][1] - vTexCoords[0][1];
    dc3c1t := vTexCoords[2][0] - vTexCoords[0][0];
    dc3c1b := vTexCoords[2][1] - vTexCoords[0][1];
    
    M := (dc2c1t * dc3c1b) - (dc3c1t * dc2c1b);
    M = 1.0 / M;
    
    dv2v1.ScaleSelf(dc3c1b)
    dv3v1.ScaleSelf(dc2c1b)
    
    vTangent :=  dv2v1.Sub(dv3v1)
    vTangent.ScaleSelf(M)  // This potentially changes the direction of the vector
    vTangent.Normalize()

    
    B := N.Cross(vTangent)
    vTangent =  B.Cross(N)
	vTangent.Normalize()
    return vTangent
}

func CalculateTangentBasis3d(vTriangle [3]Vector3d, vTexCoords [3]Vector2d, N Vector3d) Vector3d {

    dv2v1 := vTriangle[1].Sub(vTriangle[0])
    dv3v1 := vTriangle[2].Sub(vTriangle[0])
    
    dc2c1t := vTexCoords[1][0] - vTexCoords[0][0];
    dc2c1b := vTexCoords[1][1] - vTexCoords[0][1];
    dc3c1t := vTexCoords[2][0] - vTexCoords[0][0];
    dc3c1b := vTexCoords[2][1] - vTexCoords[0][1];
    
    M := (dc2c1t * dc3c1b) - (dc3c1t * dc2c1b);
    M = 1.0 / M;
    
    dv2v1.ScaleSelf(dc3c1b)
    dv3v1.ScaleSelf(dc2c1b)
    
    vTangent :=  dv2v1.Sub(dv3v1)
    vTangent.ScaleSelf(M)  // This potentially changes the direction of the vector
    vTangent.Normalize()

    
    B := N.Cross(vTangent)
    vTangent =  B.Cross(N)
	vTangent.Normalize()
    return vTangent
}

////////////////////////////////////////////////////////////////////////////
// Smoothly step between 0 and 1 between edge1 and edge 2

func SmoothStepd(edge1,edge2,x float64) float64 {
   
    t := (x - edge1) / (edge2 - edge1);
    if t > 1.0 {
        t = 1.0
	}
        
    if t < 0.0 {
        t = 0.0
	}
        
    return t * t * ( 3.0 - 2.0 * t)
}

func SmoothStepf(edge1,edge2,x float32) float32 {
   
    t := (x - edge1) / (edge2 - edge1);
    if t > 1.0 {
        t = 1.0
	}
        
    if t < 0.0 {
        t = 0.0
	}
        
    return t * t * ( 3.0 - 2.0 * t)
}

///////////////////////////////////////////////////////////////////////////
// Creae a projection to "squish" an object into the plane.
// Use m3dGetPlaneEquationf(planeEq, point1, point2, point3);
// to get a plane equation.

func MakePlanarShadowMatrix44f (pEq Vector4f, vLightPos Vector3f) Matrix44f {
	return [16]float32{  pEq[1] * -vLightPos[1] + pEq[2] * -vLightPos[2],
						-pEq[0] * -vLightPos[1],
						-pEq[0] * -vLightPos[2],
						 0.0 ,
						-pEq[1] * -vLightPos[0],
						 pEq[0] * -vLightPos[0] + pEq[2] * -vLightPos[2],
						-pEq[1] * -vLightPos[2],
						 0.0,
						-pEq[2] * -vLightPos[0],
						-pEq[2] * -vLightPos[1],
						 pEq[0] * -vLightPos[0] + pEq[1] * -vLightPos[1],
						 0.0,
						-pEq[3] * -vLightPos[0],
						-pEq[3] * -vLightPos[1],
						-pEq[3] * -vLightPos[2],
						 pEq[0] * -vLightPos[0] + pEq[1] * -vLightPos[1] + pEq[2] * -vLightPos[2] }}

func MakePlanarShadowMatrix44d (pEq Vector4d, vLightPos Vector3d) Matrix44d {
	return [16]float64{  pEq[1] * -vLightPos[1] + pEq[2] * -vLightPos[2],
						-pEq[0] * -vLightPos[1],
						-pEq[0] * -vLightPos[2],
						 0.0 ,
						-pEq[1] * -vLightPos[0],
						 pEq[0] * -vLightPos[0] + pEq[2] * -vLightPos[2],
						-pEq[1] * -vLightPos[2],
						 0.0,
						-pEq[2] * -vLightPos[0],
						-pEq[2] * -vLightPos[1],
						 pEq[0] * -vLightPos[0] + pEq[1] * -vLightPos[1],
						 0.0,
						-pEq[3] * -vLightPos[0],
						-pEq[3] * -vLightPos[1],
						-pEq[3] * -vLightPos[2],
						 pEq[0] * -vLightPos[0] + pEq[1] * -vLightPos[1] + pEq[2] * -vLightPos[2] }}

/////////////////////////////////////////////////////////////////////////////
// I want to know the point on a ray, closest to another given point in space.
// As a bonus, return the distance squared of the two points.
// In: vRayOrigin is the origin of the ray.
// In: vUnitRayDir is the unit vector of the ray
// In: vPointInSpace is the point in space
// Out: vPointOnRay is the poing on the ray closest to vPointInSpace
// Return: The square of the distance to the ray

func ClosestPointOnRay3f(vRayOrigin, vUnitRayDir, vPointInSpace Vector3f) (Vector3f,float32) {
	
	v := vPointInSpace.Sub(vRayOrigin)
	
	t := vUnitRayDir.Dot(v)
	
	var vPointOnRay Vector3f
	// This is the point on the ray
	vPointOnRay[0] = vRayOrigin[0] + (t * vUnitRayDir[0]);
	vPointOnRay[1] = vRayOrigin[1] + (t * vUnitRayDir[1]);
	vPointOnRay[2] = vRayOrigin[2] + (t * vUnitRayDir[2]);
	
	return vPointOnRay , vPointOnRay.SquaredDistanceTo(vPointInSpace)
}

func ClosestPointOnRay3d(vRayOrigin, vUnitRayDir, vPointInSpace Vector3d) (Vector3d,float64) {
	
	v := vPointInSpace.Sub(vRayOrigin)
	
	t := vUnitRayDir.Dot(v)
	
	var vPointOnRay Vector3d
	// This is the point on the ray
	vPointOnRay[0] = vRayOrigin[0] + (t * vUnitRayDir[0]);
	vPointOnRay[1] = vRayOrigin[1] + (t * vUnitRayDir[1]);
	vPointOnRay[2] = vRayOrigin[2] + (t * vUnitRayDir[2]);
	
	return vPointOnRay , vPointOnRay.SquaredDistanceTo(vPointInSpace)
}

//Calculates the signed distance of a point to a plane

func DistanceToPlanef( point Vector3f , plane Vector4f) float32 {
	return point[0]*plane[0] + point[1]*plane[1] + point[2]*plane[2] + plane[3]; }
                

func DistanceToPlaned( point Vector3d , plane Vector4d) float64 {
                	return point[0]*plane[0] + point[1]*plane[1] + point[2]*plane[2] + plane[3]; }
