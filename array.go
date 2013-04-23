package vfp

import "sort"

func Asort_float(za []float64) []float64 {
	sort.Float64s(za)
	return za
}

func Asort_string(za []string) []string {
	sort.Strings(za)
	return za
}

func Asort_int(za []int) []int {
	sort.Ints(za)
	return za
}

//Get the difference part and same part of float64 array with ascending index
func Adiff_same(zf_args_t []float64) ([]float64, []float64) {
	zf_args := make([]float64, len(zf_args_t))
	copy(zf_args, zf_args_t)
	if len(zf_args) == 0 {
		return zf_args, zf_args
	}
	Asort_float(zf_args)

	zsame := make([]float64, 0)
	zdiff := make([]float64, 0)
	var ze float64 = 0
	zchange_flag := false
	for zi, zv := range zf_args {
		if ze != zv || (zi == 0) {
			ze = zv

			zdiff = append(zdiff, ze)

			zchange_flag = true
		} else {
			if zchange_flag {
				zchange_flag = false
				zsame = append(zsame, ze)
				zchange_flag = false
			}
		}

	}
	return zdiff, zsame
}

//Copies elements from one array to another array.
//func Acopy()
type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }



/*Example:
zfl := FloatArray{3, 23, 34, 11, 56}
zfs := StringArray{"suck", "fuck", "btich"}
fmt.Println("old", zfl)
fmt.Println("adel", Adel(&zfl, 3), Adel(&zfs, 2))
fmt.Println("new", zfl, zfs)
*/
func Adel(za IArray, zi int) bool {
	return za.Del(zi)
}

type IArray interface {
	Del(zindex int) bool
}

type FloatArray []float64

func (p *FloatArray) Del(zindex int) bool {

	if zindex >= 0 && zindex < len(*p) {
		*p = append((*p)[:zindex], (*p)[zindex+1:]...)
		return true
	}

	return false
}

type StringArray []string

func (p *StringArray) Del(zindex int) bool {
	if zindex >= 0 && zindex < len(*p) {
		*p = append((*p)[:zindex], (*p)[zindex+1:]...)
		return true
	}

	return false
}
