package vfp

import odbc "github.com/weigj/go-odbc"

import "reflect"

import "fmt"

var GPrepareMapCursor map[*odbc.Statement]string = make(map[*odbc.Statement]string)
var GCursor map[string]*Cursor = make(map[string]*Cursor)
var GStatementPerConnection map[*odbc.Connection][]*odbc.Statement = make(map[*odbc.Connection][]*odbc.Statement)
var GCurrentCursor *Cursor

//Establishes a connection to a data source using a connection string.
func SqlStringConnect(zstr string) (conn *odbc.Connection, err *odbc.ODBCError) {
	return odbc.Connect(zstr)
}

//Terminates a connection to a data source.
func SqlDisConnect(conn *odbc.Connection) {

	zr, ok := GStatementPerConnection[conn]
	if ok {
		for _, zstmt := range zr {
			zstmt.Close()
		}
	} else {
		panic("Connection is invalid")
	}

	conn.Close()
}

//Prepares a SQL statement for remote execution by SQLEXEC( ).
func SqlPrepare(zconn *odbc.Connection, zselect string, zcursor string) (*odbc.Statement, *odbc.ODBCError) {
	zstmt, zerr := zconn.Prepare(zselect)
	GPrepareMapCursor[zstmt] = zcursor
	GStatementPerConnection[zconn] = append(GStatementPerConnection[zconn], zstmt)
	return zstmt, zerr
}

//Sends a SQL statement to the data source, where the statement is processed.

func SqlExec(zargs ...interface{}) (rows []*odbc.Row, err *odbc.ODBCError) {
	var stmt *odbc.Statement
	var zcursor string = ""
	if len(zargs) >= 1 && len(zargs) <= 3 {
		if reflect.ValueOf(zargs[0]).Type().String() == "*odbc.Statement" {
			stmt = zargs[0].(*odbc.Statement)
			zcursor = GPrepareMapCursor[stmt]
			delete(GPrepareMapCursor, stmt)
		}

		if reflect.ValueOf(zargs[0]).Type().String() == "*odbc.Connection" {
			var zconn *odbc.Connection
			zconn = zargs[0].(*odbc.Connection)
			zsql := reflect.ValueOf(zargs[1]).String()
			if len(zargs) == 3 {
				zcursor = reflect.ValueOf(zargs[2]).String()
			}
			stmt, err = SqlPrepare(zconn, zsql, zcursor)

			delete(GPrepareMapCursor, stmt)
			if err != nil {
				return nil, err
			}

		}

	} else {
		panic("Arguments are wrong!")
	}

	err = stmt.Execute()
	if err != nil {
		return nil, err
	}
	if zcursor != "" {
		rows, err = stmt.FetchAll()
		if err == nil {
			GCursor[zcursor] = &Cursor{Name: zcursor, Data: rows, Statement: stmt, Recno: 1}
		}
	} else {
		stmt.Close()
		return nil, nil
	}

	Select(zcursor)
	return
}

func Usein(zalias_arg ...string) bool {
	zalias := ""
	if len(zalias_arg) == 0 {
		zalias = Alias()
	} else {
		zalias = zalias_arg[0]
	}
	if Used(zalias) {
		delete(GCursor, zalias)
		return true
	}

	return false
}
func Used(zalias string) bool {
	_, ok := GCursor[zalias]
	return ok
}

//Returns the name of a field, referenced by number, in a table.
func Field(zi int, zcursorname string) (zname *odbc.Field) {

	zname, _ = GCursor[zcursorname].Statement.FieldMetadata(zi)

	return
}

func FieldIndex(zf string, zcursorname_arg ...string) int {
	zcursorname := ""
	if len(zcursorname_arg) == 0 {
		zcursorname = Alias()
	} else {
		zcursorname = zcursorname_arg[0]
	}
	zfound := -1
	for zi := 0; zi < Fcount(zcursorname); zi++ {
		zname, _ := GCursor[zcursorname].Statement.FieldMetadata(zi)
		if Lower(zname.Name) == Lower(zf) {
			zfound = zi
			break
		}
	}
	return zfound
}

//Returns the number of fields in a table.
func Fcount(zcursorname string) int {
	zn, _ := GCursor[zcursorname].Statement.NumFields()
	return zn
}

type Cursor struct {
	Name      string
	Data      []*odbc.Row
	Statement *odbc.Statement
	Recno     int
}

func (c *Cursor) String() string {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()
	zstr := ""
	zstr += "\r\n\r\nCursor " + c.Name + "'s Records:\r\n"
	zflag := "#Record   "
	var ZDISLEN float64 = 5
	zfc := Fcount(c.Name)

	ztotallen := 0
	zlen := 0
	for k := 0; k < zfc; k++ {
		zfd := Field(k, c.Name)
		zlen = int(Max(float64(zfd.Size)+ZDISLEN, float64(Lendb(zfd.Name+"("+FieldTypeString(zfd.Type)+")"))+ZDISLEN))
		ztotallen += zlen
	}
	ZVERT := "| "
	ztotallen += (zfc+1)*Lendb(ZVERT) + Lendb(zflag) + 1

	ZLINE := Replicate("-", ztotallen) + "\r\n"

	zstr += ZLINE
	zstr += ZVERT + zflag + ZVERT
	for k := 0; k < zfc; k++ {
		zfd := Field(k, c.Name)
		zlen = int(Max(float64(zfd.Size)+ZDISLEN, float64(Lendb(zfd.Name+"("+FieldTypeString(zfd.Type)+")"))+ZDISLEN))
		zstr += Padr(zfd.Name+"("+FieldTypeString(zfd.Type)+")", zlen)
		zstr += ZVERT
	}
	zstr += "\r\n"
	zstr += ZLINE

	for i, row := range c.Data {
		zstr += ZVERT
		zstr += Padr("#"+Transform(i+1), Lendb(zflag))
		zstr += ZVERT
		for k := 0; k < zfc; k++ {
			zfd := Field(k, c.Name)
			zlen = int(Max(float64(zfd.Size)+ZDISLEN, float64(Lendb(zfd.Name+"("+FieldTypeString(zfd.Type)+")"))+ZDISLEN))
			zstr += Padr(FieldValue(k, row), zlen)
			zstr += ZVERT
		}
		zstr += "\r\n"
		zstr += ZLINE
	}
	zstr += "FCount(" + Transform(Fcount(c.Name)) + ")\r\n"

	return zstr
}

//get field value according to it's type ,and return value as string
//source could be *odbc.Row or cursor name
func FieldValue(index interface{}, source_arg ...interface{}) (zret string) {

	var r *odbc.Row
	zret = ""
	zfromcursor := false
	zname := ""
	var source interface{}

	if len(source_arg) == 0 {
		source = Alias()
	} else {
		source = source_arg[0]
	}

	switch source.(type) {
	case *odbc.Row:
		r = source.(*odbc.Row)
	case string:
		zname = source.(string)
		r = GCursor[zname].Data[Recno(zname)-1]
		zfromcursor = true
	default:
		r = nil
	}

	k := -1

	if zfromcursor {
		switch index.(type) {
		case string:
			k = FieldIndex(index.(string), zname)
		default:
			k = int(Val(fmt.Sprintf("%v", index)))

		}
	} else {
		k = int(Val(fmt.Sprintf("%v", index)))
	}

	if r != nil {
		zret = fmt.Sprintf("%v", r.Get(k))
	}
	//zret = ""
	//switch t {
	//case -9:
	//	zret = r.GetString(k)
	//case -11:
	//	zret = string(r.Get(k).([]byte))
	//case 4:
	//	zret = Transform(r.GetInt(k))
	//case 2:
	//	zret = Transform(r.Get(k).([]uint8))
	//case 6:
	//	zret = Transform(r.Get(k).(float64))

	//case 93:
	//	zret = Transform(r.Get(k).(time.Time))
	//default:
	//	zret = r.GetString(k)
	//}
	return
}

func FieldTypeString(t int) (zret string) {
	zret = ""
	switch t {
	case -9:
		zret = "string"
	case -11:
		zret = "oid"
	case 4:
		zret = "int"
	case 2:
		zret = "numeric"
	case 6:
		zret = "float64"
	case 93:
		zret = "datetime"
	default:
		zret = Transform(t)
	}
	return
}

func Browse(zcursorname ...string) {
	var zstr []string
	if len(zcursorname) == 0 {

		zstr = Aline(GCurrentCursor.String())
		//fmt.Println(GCurrentCursor)
	} else {
		zstr = Aline(GCursor[zcursorname[0]].String())
		//fmt.Println(GCursor[zcursorname[0]])
	}

	for _, zl := range zstr {
		fmt.Println(zl)
	}
}

//Returns the number of records in the current or specified table.
func Reccount(zcursorname ...string) int {
	if len(zcursorname) == 0 {
		return len(GCurrentCursor.Data)
	}
	return len(GCursor[zcursorname[0]].Data)
}

//Returns the current record number in the current or specified table.
func Recno(zcursorname ...string) int {
	if len(zcursorname) == 0 {
		return GCurrentCursor.Recno
	}
	return GCursor[zcursorname[0]].Recno

}

func Select(zcursorname string) *Cursor {
	GCurrentCursor = GCursor[zcursorname]
	return GCurrentCursor
}

//Moves the record pointer forward or backward in a table.
func Skip(zn int, zcursorname ...string) *odbc.Row {
	var c *Cursor
	if len(zcursorname) == 0 {
		c = GCurrentCursor
	} else {
		c = GCursor[zcursorname[0]]
	}

	c.Recno += zn
	if c.Recno < 1 {
		c.Recno = 1
	}
	if c.Recno > Reccount(c.Name) {
		c.Recno = Reccount(c.Name)
	}

	return c.Data[c.Recno-1]
}

//Determines whether the record pointer is positioned past the last record in the current or specified table.
func Eof(zcursorname ...string) bool {
	var c string
	if len(zcursorname) == 0 {
		c = GCurrentCursor.Name
	} else {
		c = GCursor[zcursorname[0]].Name
	}

	return Reccount(c) == Recno(c)
}

//Determines whether the record pointer is positioned at the beginning of a table.
func Bof(zcursorname ...string) bool {
	var c string
	if len(zcursorname) == 0 {
		c = GCurrentCursor.Name
	} else {
		c = GCursor[zcursorname[0]].Name
	}

	return Recno(c) == 1
}

//Returns the table alias of the current or specified work area.
func Alias() string {
	if GCurrentCursor != nil {
		return GCurrentCursor.Name
	}
	return ""
}

//Moves the record pointer to the specified record number. There are multiple versions of the syntax.
//if zrec over range 1~reccount ,it means go top or go bottom
func GoRec(zrec int, zcursorname ...string) *odbc.Row {
	zname := ""
	if len(zcursorname) == 0 {
		zname = Alias()
	} else {
		zname = zcursorname[0]
	}

	c := GCursor[zname]
	c.Recno = zrec

	if c.Recno < 1 {
		c.Recno = 1
	}
	if c.Recno > Reccount(c.Name) {
		c.Recno = Reccount(c.Name)
	}

	return c.Data[c.Recno-1]

}
