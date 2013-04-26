// File generated by github.com/ungerik/pkgreflect
package vfp

import "reflect"

var Types = map[string]reflect.Type{
	"StringArray": reflect.TypeOf((*StringArray)(nil)).Elem(),
	"IntArray": reflect.TypeOf((*IntArray)(nil)).Elem(),
	"IArray": reflect.TypeOf((*IArray)(nil)).Elem(),
	"FloatArray": reflect.TypeOf((*FloatArray)(nil)).Elem(),
	"Collection": reflect.TypeOf((*Collection)(nil)).Elem(),
	"Cursor": reflect.TypeOf((*Cursor)(nil)).Elem(),
	"ShortDate": reflect.TypeOf((*ShortDate)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"Isproperty": reflect.ValueOf(Isproperty),
	"Ismethod": reflect.ValueOf(Ismethod),
	"FieldValue": reflect.ValueOf(FieldValue),
	"Fcount": reflect.ValueOf(Fcount),
	"SqlPrepare": reflect.ValueOf(SqlPrepare),
	"Reccount": reflect.ValueOf(Reccount),
	"FieldIndex": reflect.ValueOf(FieldIndex),
	"Bof": reflect.ValueOf(Bof),
	"Recno": reflect.ValueOf(Recno),
	"SqlStringConnect": reflect.ValueOf(SqlStringConnect),
	"Eof": reflect.ValueOf(Eof),
	"FieldTypeString": reflect.ValueOf(FieldTypeString),
	"Browse": reflect.ValueOf(Browse),
	"Select": reflect.ValueOf(Select),
	"Skip": reflect.ValueOf(Skip),
	"SqlExec": reflect.ValueOf(SqlExec),
	"Field": reflect.ValueOf(Field),
	"SqlDisConnect": reflect.ValueOf(SqlDisConnect),
	"Alias": reflect.ValueOf(Alias),
	"Run": reflect.ValueOf(Run),
	"Year": reflect.ValueOf(Year),
	"Addmonth": reflect.ValueOf(Addmonth),
	"Addsecond": reflect.ValueOf(Addsecond),
	"Day": reflect.ValueOf(Day),
	"Sec": reflect.ValueOf(Sec),
	"Month": reflect.ValueOf(Month),
	"Seconds": reflect.ValueOf(Seconds),
	"Addday": reflect.ValueOf(Addday),
	"Date": reflect.ValueOf(Date),
	"Addminute": reflect.ValueOf(Addminute),
	"Datetime": reflect.ValueOf(Datetime),
	"Addyear": reflect.ValueOf(Addyear),
	"Hour": reflect.ValueOf(Hour),
	"Addhour": reflect.ValueOf(Addhour),
	"Minute": reflect.ValueOf(Minute),
	"Addbs": reflect.ValueOf(Addbs),
	"Directory": reflect.ValueOf(Directory),
	"Adir": reflect.ValueOf(Adir),
	"Juststem": reflect.ValueOf(Juststem),
	"Justext": reflect.ValueOf(Justext),
	"Fullpath": reflect.ValueOf(Fullpath),
	"Justpath": reflect.ValueOf(Justpath),
	"File": reflect.ValueOf(File),
	"Filetostr": reflect.ValueOf(Filetostr),
	"Strtofile": reflect.ValueOf(Strtofile),
	"Justfname": reflect.ValueOf(Justfname),
	"Curdir": reflect.ValueOf(Curdir),
	"Floor": reflect.ValueOf(Floor),
	"Val": reflect.ValueOf(Val),
	"Atan": reflect.ValueOf(Atan),
	"Avg": reflect.ValueOf(Avg),
	"Rtod": reflect.ValueOf(Rtod),
	"Asin": reflect.ValueOf(Asin),
	"Dtor": reflect.ValueOf(Dtor),
	"Acos": reflect.ValueOf(Acos),
	"Tan": reflect.ValueOf(Tan),
	"Int": reflect.ValueOf(Int),
	"Rand": reflect.ValueOf(Rand),
	"Sin": reflect.ValueOf(Sin),
	"Sqrt": reflect.ValueOf(Sqrt),
	"Sign": reflect.ValueOf(Sign),
	"Abs": reflect.ValueOf(Abs),
	"Cos": reflect.ValueOf(Cos),
	"Sum": reflect.ValueOf(Sum),
	"Min": reflect.ValueOf(Min),
	"Ceiling": reflect.ValueOf(Ceiling),
	"Max": reflect.ValueOf(Max),
	"Lenc": reflect.ValueOf(Lenc),
	"Isdigit": reflect.ValueOf(Isdigit),
	"Replicate": reflect.ValueOf(Replicate),
	"Strtranc": reflect.ValueOf(Strtranc),
	"Islower": reflect.ValueOf(Islower),
	"Rtrim": reflect.ValueOf(Rtrim),
	"Len": reflect.ValueOf(Len),
	"Chr": reflect.ValueOf(Chr),
	"Atc": reflect.ValueOf(Atc),
	"Rightc": reflect.ValueOf(Rightc),
	"Strextract": reflect.ValueOf(Strextract),
	"Chrtranc": reflect.ValueOf(Chrtranc),
	"At": reflect.ValueOf(At),
	"Asc": reflect.ValueOf(Asc),
	"Upper": reflect.ValueOf(Upper),
	"IsNoneSingleByte": reflect.ValueOf(IsNoneSingleByte),
	"Lendb": reflect.ValueOf(Lendb),
	"Isupper": reflect.ValueOf(Isupper),
	"Mline": reflect.ValueOf(Mline),
	"Isalpha": reflect.ValueOf(Isalpha),
	"Occurs": reflect.ValueOf(Occurs),
	"Chrtran": reflect.ValueOf(Chrtran),
	"Leftc": reflect.ValueOf(Leftc),
	"Padr": reflect.ValueOf(Padr),
	"Ltrim": reflect.ValueOf(Ltrim),
	"Lower": reflect.ValueOf(Lower),
	"Transform": reflect.ValueOf(Transform),
	"Substr": reflect.ValueOf(Substr),
	"Padc": reflect.ValueOf(Padc),
	"Right": reflect.ValueOf(Right),
	"Str": reflect.ValueOf(Str),
	"Left": reflect.ValueOf(Left),
	"Strconv": reflect.ValueOf(Strconv),
	"Alltrim": reflect.ValueOf(Alltrim),
	"Substrc": reflect.ValueOf(Substrc),
	"Aline": reflect.ValueOf(Aline),
	"Padl": reflect.ValueOf(Padl),
	"Strtran": reflect.ValueOf(Strtran),
	"Isleadbyte": reflect.ValueOf(Isleadbyte),
	"Asort_string": reflect.ValueOf(Asort_string),
	"Adel": reflect.ValueOf(Adel),
	"Adiff_same": reflect.ValueOf(Adiff_same),
	"Asort_float": reflect.ValueOf(Asort_float),
	"Asort_int": reflect.ValueOf(Asort_int),
}

var Variables = map[string]reflect.Value{
	"GCursor": reflect.ValueOf(&GCursor),
	"GStatementPerConnection": reflect.ValueOf(&GStatementPerConnection),
	"GCurrentCursor": reflect.ValueOf(&GCurrentCursor),
	"GPrepareMapCursor": reflect.ValueOf(&GPrepareMapCursor),
}

