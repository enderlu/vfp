package vfp

import "odbc"

import "reflect"

var GmapCursor map[*odbc.Statement]string = make(map[*odbc.Statement]string)

//Establishes a connection to a data source using a connection string.
func SqlStringConnect(zstr string) (conn *odbc.Connection, err *odbc.ODBCError) {
	return odbc.Connect(zstr)
}

//Prepares a SQL statement for remote execution by SQLEXEC( ).
func SqlPrepare(zconn *odbc.Connection, zselect string, zcursor string) (*odbc.Statement, *odbc.ODBCError) {
	zstmt, zerr := zconn.Prepare(zselect)
	GmapCursor[zstmt] = zcursor
	return zstmt, zerr
}

//Sends a SQL statement to the data source, where the statement is processed.

func SqlExec(zargs ...interface{}) (rows []*odbc.Row, err *odbc.ODBCError) {
	var stmt *odbc.Statement

	if len(zargs) >= 1 && len(zargs) <= 3 {
		if reflect.ValueOf(zargs[0]).Type().String() == "*odbc.Statement" {
			stmt = zargs[0].(*odbc.Statement)
		}

		if len(zargs) == 2 {
			panic("arguments count  need 3 or 1 ")
		}

		if reflect.ValueOf(zargs[0]).Type().String() == "*odbc.Connection" {
			var zconn *odbc.Connection
			zconn = zargs[0].(*odbc.Connection)
			zsql := reflect.ValueOf(zargs[1]).String()
			stmt, err = zconn.Prepare(zsql)
		}

	} else {
		panic("No SQL statement")
	}

	stmt.Execute()
	rows, err = stmt.FetchAll()
	return
}
