package variables

import (
	"fmt"
)

func Conversions() {
	var f32 float32 = 10
	var f64 float64 = 10

	var convertedFloat32 float32
	convertedFloat32 = float32(f64)

	add := (f32) + float32(f64)

	fmt.Println(f32, f64, convertedFloat32, add)

}
