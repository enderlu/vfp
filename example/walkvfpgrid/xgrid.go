package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	. "vfp"
)

type productModel struct {
	walk.TableModelBase
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *productModel) RowCount() int {
	return Reccount("product")
}

func (m *productModel) ResetRows() {
	m.PublishRowsReset()

}

// Called by the TableView when it needs the text to display for a given cell.
func (m *productModel) Value(row, col int) interface{} {

	GoRec(row+1, "product")

	zname := Lower(Field(col, "product").Name)
	zv := FieldValue(zname, "product")

	switch zname {
	case "pid", "pname", "categid":
		return zv
	case "price":
		//fmt.Println("zv:", zv, Val(zv))
		return Val(zv)
	}

	panic("unexpected col")
}

var conn, _ = SqlStringConnect(`
		DRIVER=SQL Server Native Client 11.0;
		SERVER=123-PC\SQLEXPRESS;UID=sa;PWD=1;APP=from vfp;
		WSID=123-PC;DATABASE=tu1;`)

func main() {
	defer SqlDisConnect(conn)
	_, ex := SqlExec(conn, `select * from dbo.product`, "product")
	Browse("product")
	if ex != nil {
		fmt.Println("can not load product:", ex)
		return
	}

	model := new(productModel)

	MainWindow{
		Title:  "VFP data and Walk table view Example",
		Size:   Size{800, 600},
		Layout: VBox{},
		Children: []Widget{
			PushButton{
				Text: "Reset Rows",
				OnClicked: func() {
					Usein("product")
					_, ex := SqlExec(conn, `select * from dbo.product`, "product")
					if ex != nil {
						fmt.Println("can not reload product:", ex)
						return
					}
					model.ResetRows()
				},
			},
			TableView{
				AlternatingRowBGColor: walk.RGB(255, 255, 224),
				CheckBoxes:            true,
				ReorderColumnsEnabled: true,
				Columns: []TableViewColumn{
					{Title: "产品编号"},
					{Title: "产品名称"},
					{Title: "产品价格", Format: "%.2f", Alignment: AlignFar},
					{Title: "类别"},
				},
				Model: model,
			},
		},
	}.Run()
}
