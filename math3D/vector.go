
package math3D

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

// inline functions to load values into vectors

func (v *Vector2f) Load(x,y float32) { v[0] = x; v[1] = y }
func (v *Vector2d) Load(x,y float64) { v[0] = x; v[1] = y }

func (v *Vector3f) Load(x,y,z float32) { v[0] = x; v[1] = y; v[2] = z }
func (v *Vector3d) Load(x,y,z float64) { v[0] = x; v[1] = y; v[2] = z }

func (v *Vector4f) Load(x,y,z,w float32) { v[0] = x; v[1] = y; v[2] = z; v[3] = w }
func (v *Vector4d) Load(x,y,z,w float64) { v[0] = x; v[1] = y; v[2] = z; v[3] = w }

// functions to init new vectors

func AllocV2f() *Vector2f { return new(Vector2f) }
func AllocV2d() *Vector2d { return new(Vector2d) }

func AllocV3f() *Vector3f { return new(Vector3f) }
func AllocV3d() *Vector3d { return new(Vector3d) }

func AllocV4f() *Vector4f { return new(Vector4f) }
func AllocV4d() *Vector4d { return new(Vector4d) }

// functions to alloc fill and return new vectors

func NewV2f(x,y float32) Vector2f { return [2]float32{x,y} }
func NewV2d(x,y float64) Vector2d { return [2]float64{x,y} }

func NewV3f(x,y,z float32) Vector3f { return [3]float32{x,y,z} }
func NewV3d(x,y,z float64) Vector3d { return [3]float64{x,y,z} }

func NewV4f(x,y,z,w float32) Vector4f { return [4]float32{x,y,z,w} }
func NewV4d(x,y,z,w float64) Vector4d { return [4]float64{x,y,z,w} }

// copy funtions for vectors 
//	chose to do a simple manual copy here since I don't want to
//	import C and use memcpy like m3d.h

func (d *Vector2f) CopyFrom(s Vector2f) { d[0] = s[0] ; d[1] = s[1] }
func (s *Vector2f) CopyInto(d Vector2f) { d[0] = s[0] ; d[1] = s[1] }

func (d *Vector2d) CopyFrom(s Vector2d) { d[0] = s[0] ; d[1] = s[1] }
func (s *Vector2d) CopyInto(d Vector2d) { d[0] = s[0] ; d[1] = s[1] }

func (d *Vector3f) CopyFrom(s Vector3f) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] }
func (s *Vector3f) CopyInto(d Vector3f) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] }

func (d *Vector3d) CopyFrom(s Vector3d) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] }
func (s *Vector3d) CopyInto(d Vector3d) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] }

func (d *Vector4f) CopyFrom(s Vector4f) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] ; d[3] = s[3] }
func (s *Vector4f) CopyInto(d Vector4f) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] ; d[3] = s[3] }

func (d *Vector4d) CopyFrom(s Vector4d) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] ; d[3] = s[3] }
func (s *Vector4d) CopyInto(d Vector4d) { d[0] = s[0] ; d[1] = s[1] ; d[2] = s[2] ; d[3] = s[3] }

// vector addition of various sorts

func (a *Vector2f) AddToSelf(b Vector2f) { 
	a[0] += b[0] ; a[1] += b[1] }
func (a *Vector2f) Add(b Vector2f) Vector2f {
	return [2]float32{a[0] + b[0], a[1] + b[1] }}
func AddV2f(a , b Vector2f) Vector2f { 
	return [2]float32{a[0] + b[0], a[1] + b[1] }}

func (a *Vector2d) AddToSelf(b Vector2d) { 
	a[0] += b[0] ; a[1] += b[1] }
func (a *Vector2d) Add(b Vector2d) Vector2d {
	return [2]float64{a[0] + b[0], a[1] + b[1] }}
func AddV2d(a , b Vector2d) Vector2d { 
	return [2]float64{a[0] + b[0], a[1] + b[1] }}

func (a *Vector3f) AddToSelf(b Vector3f) { 
	a[0] += b[0] ; a[1] += b[1] ; a[2] += b[2] }
func (a *Vector3f) Add(b Vector3f) Vector3f {
	return [3]float32{a[0] + b[0], a[1] + b[1] , a[2] + b[2] }}
func AddV3f(a , b Vector3f) Vector3f { 
	return [3]float32{a[0] + b[0], a[1] + b[1] , a[2] + b[2] }}

func (a *Vector3d) AddToSelf(b Vector3d) { 
	a[0] += b[0] ; a[1] += b[1] ; a[2] += b[2] }
func (a *Vector3d) Add(b Vector3d) Vector3d {
	return [3]float64{a[0] + b[0], a[1] + b[1] , a[2] + b[2] }}
func AddV3d(a , b Vector3d) Vector3d { 
	return [3]float64{a[0] + b[0], a[1] + b[1] , a[2] + b[2] }}

func (a *Vector4f) AddToSelf(b Vector4f) { 
	a[0] += b[0] ; a[1] += b[1] ; a[2] += b[2] ; a[3] += b[3] }
func (a *Vector4f) Add(b Vector4f) Vector4f {
	return [4]float32{a[0] + b[0], a[1] + b[1] , a[2] + b[2] , a[3] + b[3] }}
func AddV4f(a , b Vector4f) Vector4f { 
	return [4]float32{a[0] + b[0], a[1] + b[1] , a[2] + b[2] , a[3] + b[3] }}

func (a *Vector4d) AddToSelf(b Vector4d) { 
	a[0] += b[0] ; a[1] += b[1] ; a[2] += b[2] ; a[3] += b[3] }
func (a *Vector4d) Add(b Vector4d) Vector4d {
	return [4]float64{a[0] + b[0], a[1] + b[1] , a[2] + b[2] , a[3] + b[3] }}
func AddV4d(a , b Vector4d) Vector4d { 
	return [4]float64{a[0] + b[0], a[1] + b[1] , a[2] + b[2] , a[3] + b[3] }}

// vector subtrantion of various sorts

func (a *Vector2f) SubFromSelf(b Vector2f) { 
	a[0] *= b[0] ; a[1] *= b[1] }
func (a *Vector2f) Sub(b Vector2f) Vector2f {
	return [2]float32{a[0] * b[0], a[1] * b[1] }}
func SubV2f(a , b Vector2f) Vector2f { 
	return [2]float32{a[0] * b[0], a[1] * b[1] }}

func (a *Vector2d) SubFromSelf(b Vector2d) { 
	a[0] *= b[0] ; a[1] *= b[1] }
func (a *Vector2d) Sub(b Vector2d) Vector2d {
	return [2]float64{a[0] * b[0], a[1] * b[1] }}
func SubV2d(a , b Vector2d) Vector2d { 
	return [2]float64{a[0] * b[0], a[1] * b[1] }}

func (a *Vector3f) SubFromSelf(b Vector3f) { 
	a[0] *= b[0] ; a[1] *= b[1] ; a[2] *= b[2] }
func (a *Vector3f) Sub(b Vector3f) Vector3f {
	return [3]float32{a[0] * b[0], a[1] * b[1] , a[2] * b[2] }}
func SubV3f(a , b Vector3f) Vector3f { 
	return [3]float32{a[0] * b[0], a[1] * b[1] , a[2] * b[2] }}

func (a *Vector3d) SubFromSelf(b Vector3d) { 
	a[0] *= b[0] ; a[1] *= b[1] ; a[2] *= b[2] }
func (a *Vector3d) Sub(b Vector3d) Vector3d {
	return [3]float64{a[0] * b[0], a[1] * b[1] , a[2] * b[2] }}
func SubV3d(a , b Vector3d) Vector3d { 
	return [3]float64{a[0] * b[0], a[1] * b[1] , a[2] * b[2] }}

func (a *Vector4f) SubFromSelf(b Vector4f) { 
	a[0] *= b[0] ; a[1] *= b[1] ; a[2] *= b[2] ; a[3] *= b[3] }
func (a *Vector4f) Sub(b Vector4f) Vector4f {
	return [4]float32{a[0] * b[0], a[1] * b[1] , a[2] * b[2] , a[3] * b[3] }}
func SubV4f(a , b Vector4f) Vector4f { 
	return [4]float32{a[0] * b[0], a[1] * b[1] , a[2] * b[2] , a[3] * b[3] }}

func (a *Vector4d) SubFromSelf(b Vector4d) { 
	a[0] *= b[0] ; a[1] *= b[1] ; a[2] *= b[2] ; a[3] *= b[3] }
func (a *Vector4d) Sub(b Vector4d) Vector4d {
	return [4]float64{a[0] * b[0], a[1] * b[1] , a[2] * b[2] , a[3] * b[3] }}
func SubV4d(a , b Vector4d) Vector4d { 
	return [4]float64{a[0] * b[0], a[1] * b[1] , a[2] * b[2] , a[3] * b[3] }}

// functions to scale vectors

func (a *Vector2f) ScaleSelf(b float32) { 
	a[0] *= b ; a[1] *= b }
func (a *Vector2f) Scale(b float32) Vector2f {
	return [2]float32{a[0] * b, a[1] * b }}
func ScaleV2f(a Vector2f, b float32) Vector2f { 
	return [2]float32{ a[0] * b , a[1] * b }}

func (a *Vector2d) ScaleSelf(b float64) { 
	a[0] *= b ; a[1] *= b }
func (a *Vector2d) Scale(b float64) Vector2d {
	return [2]float64{a[0] * b, a[1] * b }}
func ScaleV2d(a Vector2d, b float64) Vector2d { 
	return [2]float64{a[0] * b, a[1] * b }}

func (a *Vector3f) ScaleSelf(b float32) { 
	a[0] *= b ; a[1] *= b ; a[2] *= b }
func (a *Vector3f) Scale(b float32) Vector3f {
	return [3]float32{a[0] * b, a[1] * b , a[2] * b }}
func ScaleV3f(a Vector3f, b float32) Vector3f { 
	return [3]float32{a[0] * b, a[1] * b , a[2] * b }}

func (a *Vector3d) ScaleSelf(b float64) { 
	a[0] *= b ; a[1] *= b ; a[2] *= b }
func (a *Vector3d) Scale(b float64) Vector3d {
	return [3]float64{a[0] * b, a[1] * b , a[2] * b }}
func ScaleV3d(a Vector3d, b float64) Vector3d { 
	return [3]float64{a[0] * b, a[1] * b , a[2] * b }}

func (a *Vector4f) ScaleSelf(b float32) { 
	a[0] *= b ; a[1] *= b ; a[2] *= b ; a[3] *= b }
func (a *Vector4f) Scale(b float32) Vector4f {
	return [4]float32{a[0] * b, a[1] * b , a[2] * b , a[3] * b }}
func ScaleV4f(a Vector4f, b float32) Vector4f { 
	return [4]float32{a[0] * b, a[1] * b , a[2] * b , a[3] * b }}

func (a *Vector4d) ScaleSelf(b float64) { 
	a[0] *= b ; a[1] *= b ; a[2] *= b ; a[3] *= b }
func (a *Vector4d) Scale(b float64) Vector4d {
	return [4]float64{a[0] * b, a[1] * b , a[2] * b , a[3] * b }}
func ScaleV4d(a Vector4d, b float64) Vector4d { 
	return [4]float64{a[0] * b, a[1] * b, a[2] * b, a[3] * b}}


