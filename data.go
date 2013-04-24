package vfp

import "odbc"

import "reflect"
import "time"
import "fmt"

var GPrepareMapCursor map[*odbc.Statement]string = make(map[*odbc.Statement]string)
var GCursor map[string]*Cursor = make(map[string]*Cursor)
var GStatementPerConnection map[*odbc.Connection][]*odbc.Statement = make(map[*odbc.Connection][]*odbc.Statement)

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

		if len(zargs) == 2 {
			panic("arguments count  need 3 or 1 ")
		}

		if reflect.ValueOf(zargs[0]).Type().String() == "*odbc.Connection" {
			var zconn *odbc.Connection
			zconn = zargs[0].(*odbc.Connection)
			zsql := reflect.ValueOf(zargs[1]).String()
			zcursor = reflect.ValueOf(zargs[2]).String()
			stmt, err = SqlPrepare(zconn, zsql, zcursor)
			delete(GPrepareMapCursor, stmt)
		}

	} else {
		panic("No SQL statement")
	}

	if zcursor == "" {
		panic("Cursor name is empty")
	}

	stmt.Execute()
	rows, err = stmt.FetchAll()

	GCursor[zcursor] = &Cursor{Name: zcursor, Data: rows, Statement: stmt}

	return
}

//Returns the name of a field, referenced by number, in a table.
func Field(zi int, zcursorname string) (zname *odbc.Field) {

	zname, _ = GCursor[zcursorname].Statement.FieldMetadata(zi)

	return
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
}

func (c *Cursor) String() string {
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
			zstr += Padr(FieldValue(zfd.Type, k, row), zlen)
			zstr += ZVERT
		}
		zstr += "\r\n"
		zstr += ZLINE
	}
	zstr += "FCount(" + Transform(Fcount(c.Name)) + ")\r\n"

	return zstr
}

//get field value according to it's type ,and return value as string 
func FieldValue(t int, k int, r *odbc.Row) (zret string) {
	zret = ""
	switch t {
	case -9:
		zret = r.GetString(k)
	case -11:
		zret = string(r.Get(k).([]byte))
	case 4:
		zret = Transform(r.GetInt(k))
	case 2:
		zret = Transform(r.Get(k).([]uint8))
	case 6:
		zret = Transform(r.Get(k).(float64))

	case 93:
		zret = Transform(r.Get(k).(time.Time))
	default:
		zret = r.GetString(k)
	}
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

func Browse(c *Cursor) {
	fmt.Println(c)
}
