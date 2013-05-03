package main

import (
	"vfp"
)

func main() {

	conn, _ := vfp.SqlStringConnect(`
		DRIVER=SQL Server Native Client 11.0;
		SERVER=123-PC\SQLEXPRESS;UID=sa;PWD=1;APP=from vfp;
		WSID=123-PC;DATABASE=tu1;`)

	vfp.SqlExec(conn, `select [curr_name]      ,
											[OptimisticLockField]      ,
											[GCRecord] from curr `, "mycursor")

	vfp.Browse(vfp.GCursor["mycursor"])
	defer vfp.SqlDisConnect(conn)

}
