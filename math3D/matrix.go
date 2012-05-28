
package math3D

// all the matrix functions I can get from math.3d
//	the original version used memcpy and i'd like 
//	to avoid that, plus I don't know if gc does loop
//	unrolling right

func (s *Matrix33f) CopyTo(d Matrix33f) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] 
	d[3] = s[3] ; d[4] = s[4] ; d[5] = s[5] 
	d[6] = s[6] ; d[7] = s[7] ; d[8] = s[8] 
}

func (s *Matrix33d) CopyTo(d Matrix33d) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] 
	d[3] = s[3] ; d[4] = s[4] ; d[5] = s[5] 
	d[6] = s[6] ; d[7] = s[7] ; d[8] = s[8] 
}

func (s *Matrix44f) CopyTo(d Matrix44f) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2]; d[3] = d[3]  
	d[4] = s[4] ; d[5] = s[5] ; d[6] = s[6]; d[7] = d[7]  
	d[8] = s[8] ; d[9] = s[9] ; d[10] = s[10]; d[11] = d[11]  
	d[12] = s[12] ; d[13] = s[13] ; d[14] = s[14]; d[15] = d[15]  
}

func (s *Matrix44d) CopyTo(d Matrix44d) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2]; d[3] = d[3]  
	d[4] = s[4] ; d[5] = s[5] ; d[6] = s[6]; d[7] = d[7]  
	d[8] = s[8] ; d[9] = s[9] ; d[10] = s[10]; d[11] = d[11]  
	d[12] = s[12] ; d[13] = s[13] ; d[14] = s[14]; d[15] = d[15]  
}

func (d *Matrix33f) CopyFrom(s Matrix33f) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] 
	d[3] = s[3] ; d[4] = s[4] ; d[5] = s[5] 
	d[6] = s[6] ; d[7] = s[7] ; d[8] = s[8] 
}

func (d *Matrix33d) CopyFrom(s Matrix33d) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] 
	d[3] = s[3] ; d[4] = s[4] ; d[5] = s[5] 
	d[6] = s[6] ; d[7] = s[7] ; d[8] = s[8] 
}

func (d *Matrix44f) CopyFrom(s Matrix44f) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2]; d[3] = d[3]  
	d[4] = s[4] ; d[5] = s[5] ; d[6] = s[6]; d[7] = d[7]  
	d[8] = s[8] ; d[9] = s[9] ; d[10] = s[10]; d[11] = d[11]  
	d[12] = s[12] ; d[13] = s[13] ; d[14] = s[14]; d[15] = d[15]  
}

func (d *Matrix44d) CopyFrom(s Matrix44d) {
	d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2]; d[3] = d[3]  
	d[4] = s[4] ; d[5] = s[5] ; d[6] = s[6]; d[7] = d[7]  
	d[8] = s[8] ; d[9] = s[9] ; d[10] = s[10]; d[11] = d[11]  
	d[12] = s[12] ; d[13] = s[13] ; d[14] = s[14]; d[15] = d[15]  
}


// get identity matrices of various sorts

func Identity33f() Matrix33f { return [9]float32{1.0,0.0,0.0, // this is stil in column major
												 0.0,1.0,0.0,
												 0.0,0.0,1.0}}

func Identity33d() Matrix33d { return [9]float64{1.0,0.0,0.0, // this is stil in column major
												 0.0,1.0,0.0,
												 0.0,0.0,1.0}}

func Identity44f() Matrix44f { return [16]float32{1.0,0.0,0.0,0.0, // this is stil in column major
												  0.0,1.0,0.0,0.0,
												  0.0,0.0,1.0,0.0,
												  0.0,0.0,0.0,1.0}}

func Identity44d() Matrix44d { return [16]float64{1.0,0.0,0.0,0.0, // this is stil in column major
												  0.0,1.0,0.0,0.0,
												  0.0,0.0,1.0,0.0,
												  0.0,0.0,0.0,1.0}}

// setters and getters for the various matrix columns

func (m *Matrix33f) GetColumn(column int) Vector3f { return [3]float32{ m[3 * column ],
																		m[3 * column + 1],
																		m[3 * column + 2]}}

func (m *Matrix33d) GetColumn(column int) Vector3d { return [3]float64{ m[3*column],
																		m[3*column + 1],
																		m[3*column + 2]}}

func (m *Matrix44f) GetColumn(column int) Vector4f { return [4]float32{ m[4*column],
																		m[4*column + 1],
																		m[4*column + 2],
																		m[4*column + 3]}}

func (m *Matrix44d) GetColumn(column int) Vector4d { return [4]float64{ m[4*column],
																		m[4*column + 1],
																		m[4*column + 2],
																		m[4*column + 3]}}

func (m *Matrix33f) SetColumn(v Vector3f, col int) { m[3 * col] = v[0] 
													 m[3 * col + 1] = v[1]
													 m[3 * col + 2] = v[2] }

func (m *Matrix33d) SetColumn(v Vector3d, col int) { m[4 * col] = v[0] 
													 m[4 * col + 1] = v[1]
													 m[4 * col + 2] = v[2] }

func (m *Matrix44f) SetColumn(v Vector4f, col int) { m[4 * col] = v[0] 
													 m[4 * col + 1] = v[1]
													 m[4 * col + 2] = v[2]
													 m[4 * col + 3] = v[3] }

func (m *Matrix44d) SetColumn(v Vector4d, col int) { m[4 * col] = v[0] 
													 m[4 * col + 1] = v[1]
													 m[4 * col + 2] = v[2]
													 m[4 * col + 3] = v[3] }

// functions to extract the rotation matrices from 4x4matrices

func (m *Matrix44f) ExtractRotationMatrix() Matrix33f {
	return [9]float32{m[0],m[1],m[2],m[4],m[5],m[6],m[8],m[9],m[10]}
}

func (m *Matrix44d) ExtractRotationMatrix() Matrix33d {
	return [9]float64{m[0],m[1],m[2],m[4],m[5],m[6],m[8],m[9],m[10]}
}

// functions that inject a rotation matrix into a 4x4 matrix

func (m *Matrix44f) InjectRotationMatrix(r Matrix33f) {
	m[0] = r[0] ; m[1] = r[1] ; m[2] = r[2]  
	m[4] = r[3] ; m[5] = r[4] ; m[6] = r[5]
	m[8] = r[6] ; m[9] = r[7] ; m[10] = r[8]
}

func (m *Matrix44d) InjectRotationMatrix(r Matrix33d) {
	m[0] = r[0] ; m[1] = r[1] ; m[2] = r[2]  
	m[4] = r[3] ; m[5] = r[4] ; m[6] = r[5]
	m[8] = r[6] ; m[9] = r[7] ; m[10] = r[8]
}

// matrix multiplication functions. 

func (a *Matrix33f) Multiply(b *Matrix33f) Matrix33f {
	
	var ai0 , ai1, ai2 float32
	var out Matrix33f

	for i := 0 ; i < 3 ; i++ {
		ai0 = a[i] 
		ai1 = a[3+i]
		ai2 = a[6+i]
   		out[0+i]  = ai0 * b[0] + ai1 * b[1] + ai2 * b[2]
   		out[3+i]  = ai0 * b[3] + ai1 * b[4] + ai2 * b[5]
   		out[6+i]  = ai0 * b[6] + ai1 * b[7] + ai2 * b[8]
	}

	return out
}

func (a *Matrix33d) Multiply(b *Matrix33d) Matrix33d {
	
	var ai0 , ai1, ai2 float64
	var out Matrix33d

	for i := 0 ; i < 3 ; i++ {
		ai0 = a[i] 
		ai1 = a[3+i]
		ai2 = a[6+i]
   		out[0+i]  = ai0 * b[0] + ai1 * b[1] + ai2 * b[2]
   		out[3+i]  = ai0 * b[3] + ai1 * b[4] + ai2 * b[5]
   		out[6+i]  = ai0 * b[6] + ai1 * b[7] + ai2 * b[8]
	}

	return out
}

func (a *Matrix44f) Multiply(b *Matrix44f) Matrix44f {
	
	var ai0 , ai1, ai2 , ai3 float32
	var out Matrix44f

	for i := 0 ; i < 4 ; i++ {
		ai0 = a[i] 
		ai1 = a[4+i]
		ai2 = a[8+i]
		ai3 = a[12+i]
   		out[0+i]  = ai0 * b[0] + ai1 * b[1] + ai2 * b[2] + ai3 * b[3]
   		out[4+i]  = ai0 * b[4] + ai1 * b[5] + ai2 * b[6] + ai3 * b[7]
   		out[8+i]  = ai0 * b[8] + ai1 * b[9] + ai2 * b[10] + ai3 * b[11]
   		out[12+i] = ai0 * b[12] + ai1 * b[13] + ai2 * b[14] + ai3 * b[15]
	}

	return out
}

func (a *Matrix44d) Multiply(b *Matrix44d) Matrix44d {
	
	var ai0 , ai1, ai2 , ai3 float64
	var out Matrix44d

	for i := 0 ; i < 4 ; i++ {
		ai0 = a[i] 
		ai1 = a[4+i]
		ai2 = a[8+i]
		ai3 = a[12+i]
   		out[0+i]  = ai0 * b[0] + ai1 * b[1] + ai2 * b[2] + ai3 * b[3]
   		out[4+i]  = ai0 * b[4] + ai1 * b[5] + ai2 * b[6] + ai3 * b[7]
   		out[8+i]  = ai0 * b[8] + ai1 * b[9] + ai2 * b[10] + ai3 * b[11]
   		out[12+i] = ai0 * b[12] + ai1 * b[13] + ai2 * b[14] + ai3 * b[15]
	}

	return out
}


