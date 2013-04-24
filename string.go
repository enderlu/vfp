package vfp

import "strconv"
import "strings"
import "unicode/utf8"

import "reflect"

func Str(znum float64, zlen int, zdecimal int) string {
	return strconv.FormatFloat(znum, 'f', zdecimal, 64)[0:zlen]
}

func Substr(zstr string, zstart, zlen int) string {
	return zstr[zstart-1 : zlen]
}

func Substrc(zstr string, zstart, zlen int) string {
	return string([]rune(zstr)[zstart-1 : zstart+zlen-1])
}

func At(zsubstring, zwholestring string) int {
	return strings.Index(zwholestring, zsubstring) + 1
}

//for  single-byte
func Len(zstr string) int {
	return len(zstr)
}

//for utf-8 = 3  single-byte
func Lenc(zstr string) int {
	return utf8.RuneCountInString(zstr)
}

//Lendb is designed for expressions containing double-byte characters
func Lendb(zstr string) int {
	//zdouble := (Len(zstr) - Lenc(zstr)) / 2
	//zsingle := Len(zstr) - zdouble*3

	//return zdouble*2 + zsingle

	znoneSingle := 0
	zSingle := 0
	for _, zv := range []rune(zstr) {
		if utf8.RuneLen(zv) != 1 {
			znoneSingle++
		} else {
			zSingle++
		}
	}
	return znoneSingle*2 + zSingle
}

func Left(zstr string, zlen int) string {
	return zstr[0:zlen]
}

func Right(zstr string, zlen int) string {
	return zstr[len(zstr)-zlen:]
}

//bugs not fixed
func Atc(zsubstring, zwholestring string) int {
	zn := strings.Index(zwholestring, zsubstring)

	zs := Left(zwholestring, zn)
	return Lenc(zs) + 1
}

func Leftc(zstr string, zlen int) string {
	return string([]rune(zstr)[0:zlen])
}

func Rightc(zstr string, zlen int) string {
	var zs []rune
	zs = []rune(zstr)
	return string(zs[len(zs)-zlen:])
}

/*Replaces each character in a character expression 
that matches a character in a second character expression 
with the corresponding character in a third character expression.

fmt.Println("Chrtran:", vfp.Chrtran("ABCDEF", "ACE", "XYZQRST")) 

Result:XBYDZF
*/
func Chrtran(zwhole, zsearch, zreplace string) string {
	var zlen int
	var zvn, zvo string
	zlen = len(zreplace)
	for zi, zrune := range zsearch {
		zvo = string(zrune)
		if zi >= zlen {
			zvn = ""
		} else {
			zvn = string(zreplace[zi])
		}

		zwhole = strings.Replace(zwhole, zvo, zvn, -1)
	}
	return zwhole
}
func Chrtranc(zwhole, zsearch, zreplace string) string {
	return Chrtran(zwhole, zsearch, zreplace)
}

/*
Returns the unicode value for the leftmost character in a character expression.
*/
func Asc(zstr string) int {
	return int([]rune(zstr)[0])
}

/*
Returns the character associated with the specified numeric unicode code.
*/
func Chr(zcode int) string {
	return string(zcode)
}

/*
Returns a character string from an expression in a format determined by a format code.
*/
func Transform(zExprArgs interface{}) string {
	return toString(reflect.ValueOf(zExprArgs))
}

func toString(val reflect.Value) string {

	var str string
	if !val.IsValid() {
		return "<zero Value>"
	}
	typ := val.Type()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return strconv.FormatFloat(real(c), 'g', -1, 64) + "+" + strconv.FormatFloat(imag(c), 'g', -1, 64) + "i"
	case reflect.String:
		return val.String()
	case reflect.Bool:
		if val.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Ptr:
		v := val
		str = typ.String() + "("
		if v.IsNil() {
			str += "0"
		} else {
			str += "&" + toString(v.Elem())
		}
		str += ")"
		return str
	case reflect.Array, reflect.Slice:
		v := val
		str += typ.String()
		str += "{"
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				str += ", "
			}
			str += toString(v.Index(i))
		}
		str += "}"
		return str
	case reflect.Map:
		t := typ
		str = t.String()
		str += "{"
		str += "<can't iterate on maps>"
		str += "}"
		return str
	case reflect.Chan:
		str = typ.String()
		return str
	case reflect.Struct:
		m, ok := val.Type().MethodByName("String")
		if ok {
			str = m.Func.Call([]reflect.Value{val})[0].String()
		} else {

			t := typ
			v := val
			str += t.String()
			str += "{"
			for i, n := 0, v.NumField(); i < n; i++ {
				if i > 0 {
					str += ", "
				}
				str += toString(v.Field(i))
			}
			str += "}"
		}
		return str
	case reflect.Interface:
		return typ.String() + "(" + toString(val.Elem()) + ")"
	case reflect.Func:
		v := val
		return typ.String() + "(" + strconv.FormatUint(uint64(v.Pointer()), 10) + ")"
	default:
		panic("Transform: can't print type " + typ.String())
	}
	return "Transform: can't happen"
}

//Retrieves a string between two delimiters. 
/*Example:
Strextract("<apple>","<",">" ,1)
*/
func Strextract(zstr, zbd, zed string, znum int) string {

	return strings.Split(strings.Split(zstr, zbd)[znum], zed)[0]

}

func Upper(zstr string) string {
	return strings.ToUpper(zstr)
}

func Lower(zstr string) string {
	return strings.ToLower(zstr)
}

//Determines whether the leftmost character of the specified character expression is a digit (0 through 9).
/*Example:
Isdigit("5443.3")//true
Isdigit("b66443.3")//false
*/
func Isdigit(zstr string) bool {
	_, zerr := strconv.ParseFloat(zstr, 64)
	if zerr != nil {
		return false
	}
	return true
}

//Determines whether the leftmost character in a character expression is alphabetic.

func Isalpha(zstr string) bool {
	zv := string(zstr[0])
	return (Asc(zv) >= Asc("A") && Asc(zv) <= Asc("Z")) ||
		(Asc(zv) >= Asc("a") && Asc(zv) <= Asc("z"))
}

//Determines whether the leftmost character of the specified character expression is a lowercase alphabetic character.

func Islower(zstr string) bool {
	zv := string(zstr[0])
	return (Asc(zv) >= Asc("a") && Asc(zv) <= Asc("z"))
}

//Determines whether the first character in a character expression is an uppercase alphabetic character.
func Isupper(zstr string) bool {
	zv := string(zstr[0])
	return Asc(zv) >= Asc("A") && Asc(zv) <= Asc("Z")
}

//Searches a character expression or memo field for a second character expression or memo 
//field and replaces each occurrence with a third character expression or memo field. 
//You can specify where the replacement begins and how many replacements are made. 
func Strtran(zstr, zsearch, zreplace string, zn int) string {
	return strings.Replace(zstr, zsearch, zreplace, zn)
}
func Strtranc(zstr, zsearch, zreplace string, zn int) string {
	return Strtran(zstr, zsearch, zreplace, zn)
}

//Returns a character string that contains a specified character expression repeated a specified number of times.
/*Example:
Replicate("HELLO ",4) // Displays HELLO HELLO HELLO HELLO
*/
func Replicate(zstr string, ztimes int) string {
	return strings.Repeat(zstr, ztimes)
}

/*
Returns a string from an expression, 
padded with spaces or characters to a specified length on the left or right sides, or both.
*/
/*Example:

*/
func Padl(zstr string, zlen int, zpadchar_arg ...string) string {

	zpadchar := " "
	if len(zpadchar_arg) > 0 {
		zpadchar = zpadchar_arg[0]
	}
	if zlen > len(zstr) {
		return Replicate(zpadchar, zlen-Lendb(zstr)) + zstr
	}
	return Substrc(zstr, 1, zlen)
}

func Padr(zstr string, zlen int, zpadchar_arg ...string) string {

	zpadchar := " "
	if len(zpadchar_arg) > 0 {
		zpadchar = zpadchar_arg[0]
	}

	if zlen > len(zstr) {
		return zstr + Replicate(zpadchar, zlen-Lendb(zstr))
	}

	return Substrc(zstr, 1, zlen)
}

func Padc(zstr string, zlen int, zpadchar_arg ...string) string {

	zpadchar := " "
	if len(zpadchar_arg) > 0 {
		zpadchar = zpadchar_arg[0]
	}
	if zlen > len(zstr) {
		zl, zr := 0, 0
		if int((zlen-Lendb(zstr))/2) != (zlen - Lendb(zstr)) {
			zl = int((zlen - Lendb(zstr)) / 2)
			zr = int((zlen-Lendb(zstr))/2) + 1
		}
		return Replicate(zpadchar, zl) + zstr + Replicate(zpadchar, zr)
	}
	return Substrc(zstr, 1, zlen)
}

/*
Removes all leading and trailing spaces or parsing characters from the specified character 
expression, or all leading and trailing zero (0) bytes from the specified binary expression.
*/
func Alltrim(zstr string, zargs ...string) string {
	zc := ""
	if len(zargs) > 0 {
		for _, zv := range zargs {
			zc += string(zv)
		}
	} else {
		zc = " "
	}
	return strings.Trim(zstr, zc)
}

/*
Removes all leading spaces or parsing characters from the specified character expression, 
or all leading zero (0) bytes from the specified binary expression.
*/
func Ltrim(zstr string, zargs ...string) string {
	zc := ""
	if len(zargs) > 0 {
		for _, zv := range zargs {
			zc += string(zv)
		}
	} else {
		zc = " "
	}
	return strings.TrimLeft(zstr, zc)
}

/*
Removes all trailing spaces or parsing characters from the specified character expression, 
or all trailing zero (0) bytes from the specified binary expression.
*/
func Rtrim(zstr string, zargs ...string) string {
	zc := ""

	if len(zargs) > 0 {
		for _, zv := range zargs {
			zc += string(zv)
		}
	} else {
		zc = " "
	}

	return strings.TrimRight(zstr, zc)
}

/*
Copies each line in a character expression or memo field to a corresponding row in an array.

Example:
	Aline("we go to school" ," ") //[we,go,to,school]
*/
func Aline(zstr string, zseperator_arg ...string) (zlines []string) {
	zseperator := "\r\n"
	if len(zseperator_arg) > 0 {
		zseperator = zseperator_arg[0]
	}
	zlines = strings.Split(zstr, zseperator)
	return
}

//Copies each line in a character expression or memo field to a corresponding row in an array.
/*Example:
vfp.Mline("apple\r\norange\r\nbanana" ,2,1)
*/
func Mline(zstr string, zlineno int, zcount_arg ...int) (zret string) {
	zcount := 0
	if len(zcount_arg) > 0 {
		zcount = zcount_arg[0]
	}
	zlines := Aline(zstr)
	zret = ""
	zlen := len(zlines)
	for zi := zlineno; zi >= zlineno && zi <= zlineno+zcount; zi++ {
		if zi >= zlen {
			break
		}
		zret += zlines[zi]
		if zcount > 0 && zi < zlineno+zcount {
			zret += "\r\n"
		}
	}
	return
}

//Returns true  if the first character are utf-8 in string
func Isleadbyte(zstr string) bool {
	return utf8.RuneLen([]rune(zstr)[0]) > 1
}

//Returns true if there are none-singlebyte characters
func IsNoneSingleByte(zstr string) bool {
	return Len(zstr) != Lenc(zstr)
}

//Converts character expressions 
//between single-byte, double-byte, UNICODE, and locale-specific representations.
func Strconv(zstr string, zconvertType int) string {
	return ""
}

//Returns the number of times a character expression occurs within another character expression.
func Occurs(zsubstring ,zwholestring string) int {
	return strings.Count(zwholestring ,zsubstring)
}
