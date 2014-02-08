package main

import (
	"fmt"
	"github.com/enderlu/vfp"
	"os"
)

var Browse = vfp.Browse
var SqlStringConnect = vfp.SqlStringConnect
var SqlExec = vfp.SqlExec
var Select = vfp.Select
var SqlDisConnect = vfp.SqlDisConnect

func main() {

	zss := 0
	zcount := 1000
	if len(os.Args) >= 2 {
		zcount = int(vfp.Val(os.Args[1]))
	}
	conn, _ := SqlStringConnect(`
		DRIVER=SQL Server Native Client 11.0;
		SERVER=123-PC\SQLEXPRESS;UID=sa;PWD=1;APP=from vfp;
		WSID=123-PC;DATABASE=tu1;`)

	SqlExec(conn, `delete from [order]`)
	zs := vfp.Seconds()
	for zi := 1; zi <= zcount; zi++ {
		zsql := fmt.Sprintf(`insert into [order] values('C%v' ,'%v' ,'%v' ,%v ,%v)`,
			vfp.Padl(vfp.Transform(zi), 9, "0"), vfp.Date(), vfp.Datetime(), vfp.Rand(zi, 100), vfp.Rand(zi, 700))
		_, err1 := SqlExec(conn, zsql)
		if err1 != nil {
			fmt.Println(err1)
			fmt.Println(zsql)
		} else {
			zss++
		}
	}

	fmt.Println("插入所花时间:", zss, "条", vfp.Seconds()-zs)
	zs = vfp.Seconds()

	_, err := SqlExec(conn, ` select top 10 * from [order] order by client desc `, "order")
	Select("order")
	Browse()
	fmt.Println("提取", vfp.Reccount(), "所花时间:", vfp.Seconds()-zs)

	if err != nil {
		println(err.Error())
	}

	zs = vfp.Seconds()
	_, err = SqlExec(conn, `SELECT TOP 1000 cast([Oid] as nvarchar(72)) as oid 
      ,[curr_name]
      ,[OptimisticLockField]
      ,[GCRecord]
  FROM [tu1].[dbo].[curr]
 `, "curr")
	Select("curr")
	Browse()
	fmt.Println("提取", vfp.Alias(), vfp.Reccount(), "条记录所花时间:", vfp.Seconds()-zs)

	if err != nil {
		println(err.Error())
	}

	vfp.Skip(1)

	fmt.Println(vfp.FieldValue("curr_name"))
	vfp.Skip(-1)
	fmt.Println(vfp.FieldValue(1, "order"), vfp.FieldValue("oid", "curr"))
	defer SqlDisConnect(conn)

}
