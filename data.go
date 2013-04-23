package vfp

import "odbc"

import "reflect"

var GPrepareMapCursor map[*odbc.Statement]string = make(map[*odbc.Statement]string)
var GCursor map[string]*Cursor = make(map[string]*Cursor)

//Establishes a connection to a data source using a connection string.
func SqlStringConnect(zstr string) (conn *odbc.Connection, err *odbc.ODBCError) {
	return odbc.Connect(zstr)
}

//Prepares a SQL statement for remote execution by SQLEXEC( ).
func SqlPrepare(zconn *odbc.Connection, zselect string, zcursor string) (*odbc.Statement, *odbc.ODBCError) {
	zstmt, zerr := zconn.Prepare(zselect)
	GPrepareMapCursor[zstmt] = zcursor
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
			stmt, err = zconn.Prepare(zsql)
			zcursor = reflect.ValueOf(zargs[2]).String()
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
	return GCursor[zcursorname].Statement.NumFields()
}

type Cursor struct {
	Name      string
	Data      []*odbc.Row
	Statement *odbc.Statement
}
