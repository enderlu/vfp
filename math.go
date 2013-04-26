package vfp

import "math"
import "math/rand"
import "strconv"
import "fmt"

func Abs(zfloat float64) float64 {
	return math.Abs(zfloat)
}

func Int(zfloat float64) int {
	return int(zfloat)
}

/*
	Returns a random number between 0 and zrange.
*/
/*Example:
Rand(1 ,100) //1~100

*/
func Rand(zx int, zrange int) int {
	var r *rand.Rand
	r = rand.New(rand.NewSource(int64(zx)))
	return int(r.Intn(zrange))
}

//Returns the sine of an angle.
func Sin(zf float64) float64 {
	return math.Sin(zf)
}

//Converts degrees to radians.
func Dtor(zf float64) float64 {
	return math.Pi * zf / 180
}

//Converts radians to its equivalent in degrees.
func Rtod(zf float64) float64 {
	return 180 * zf / math.Pi
}

//Returns the cosine of a numeric expression.
func Cos(zf float64) float64 {
	return math.Cos(zf)
}

//Returns the arc cosine of a specified numeric expression.
func Acos(zf float64) float64 {
	return math.Acos(zf)
}

//Returns in radians the arc sine of a numeric expression.
func Asin(zf float64) float64 {
	return math.Asin(zf)
}

//This trigonometric function returns the tangent of an angle.
func Tan(zf float64) float64 {
	return math.Tan(zf)
}

//Returns in radians the arc tangent of a numeric expression.
func Atan(zf float64) float64 {
	return math.Atan(zf)
}

//Returns the square root of the specified numeric expression.
func Sqrt(zf float64) float64 {
	return math.Sqrt(zf)
}

//Evaluates a set of expressions and returns the expression with the maximum value.
func Max(zf1, zf2 float64, zf_args ...float64) float64 {
	if len(zf_args) == 0 {
		return math.Max(zf1, zf2)
	}
	zv := math.Max(zf1, zf2)
	for zi := 0; zi < len(zf_args); zi++ {
		zv = math.Max(zv, zf_args[zi])
	}
	return zv
}

//Evaluates a set of expressions and returns the expression with the minimum value.
func Min(zf1, zf2 float64, zf_args ...float64) float64 {
	if len(zf_args) == 0 {
		return math.Max(zf1, zf2)
	}
	zv := math.Min(zf1, zf2)
	for zi := 0; zi < len(zf_args); zi++ {
		zv = math.Min(zv, zf_args[zi])
	}
	return zv
}

//Totals all on numeric array
func Sum(zf_args []float64) float64 {
	if len(zf_args) == 0 {
		return 0
	}
	var zv float64 = 0
	for zi := 0; zi < len(zf_args); zi++ {
		zv = zv + zf_args[zi]
	}
	return zv
}

//Computes the arithmetic average of numeric expressions
func Avg(zf_args []float64) float64 {
	if len(zf_args) == 0 {
		return 0
	}

	return Sum(zf_args) / float64(len(zf_args))
}

func Val(zd string) float64 {
	zr, _ := strconv.ParseFloat(zd, 64)
	return zr
}

//Returns the next highest integer that is greater than or equal to the specified numeric expression.
/*
Ceiling(10.9)  // Displays 11

Ceiling(-10.9)  // Displays -10

Ceiling(10.0)  // Displays 10

Ceiling(-10.0) // Displays -10
*/
func Ceiling(z float64) float64 {
	if z > 0 && z > float64(Int(z)) {
		return float64(Int(z) + 1)
	}
	return float64(Int(z))
}

//Returns a numeric value of 1, –1, or 0 if the specified numeric expression evaluates to a positive, negative, or 0 value.
func Sign(z interface{}) int {
	zv := Val(fmt.Sprintf("%v", z))
	if zv > 0 {
		return 1
	}
	if zv < 0 {
		return -1
	}
	return 0

}

//Returns the nearest integer that is less than or equal to the specified numeric expression.
/*
Floor(10.9)  // Displays 10

Floor(-10.9)  // Displays -11

Floor(10.0)  // Displays 10

Floor(-10.0) // Displays -10
*/
func Floor(z float64) float64 {
	if z < 0 && z < float64(Int(z)) {
		return float64(Int(z) - 1)
	}
	return float64(Int(z))
}

//Returns a numeric expression rounded to a specified number of decimal places.
/*Example
fmt.Println("round:",
	vfp.Round(10.4545, 2),
	vfp.Round(10.4545, 1),
	vfp.Round(-10.5545, 0))//round: 10.45 10.5 -11

*/
func Round(zval float64, zdecimal float64) float64 {
	zsign := Sign(zval)

	zval = Abs(zval)
	zseed := float64(math.Pow(10, zdecimal))
	zval = zval * zseed

	zint := int(zval)

	if zval-float64(zint) >= 0.5 {
		zval += 1
	}

	zval = float64(int(zval))
	zval /= zseed

	return zval * float64(zsign)
}
