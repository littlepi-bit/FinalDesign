package Model

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 创建全局连接池句柄
var GlobalConn *gorm.DB
var GBmu sync.Mutex

func OpenDatabase(remote bool) {
	tmp := ""
	if remote {
		//tmp = "root:123456@(120.77.12.35:3306)/costfile?charset=utf8mb4&parseTime=True&loc=Local"
		tmp = "root:123456@(192.168.146.131:3306)/costfile?charset=utf8mb4&parseTime=True&loc=Local"
	} else {
		tmp = "root:123456@(127.0.0.1:3306)/costfile?charset=utf8mb4&parseTime=True&loc=Local"
	}
	conn, err := gorm.Open("mysql", tmp)
	if err != nil {
		log.Fatal("failed to connect database：" + err.Error())
		return
	}
	GlobalConn = conn
}

func CloseDatabase() {
	GlobalConn.Close()
}

func InitDatebase() {
	GlobalConn.Table("project").DropTable(&Pro{})
	GlobalConn.Table("individual").DropTable(&Ind{})
	GlobalConn.Table("unit").DropTable(&Unit{})
	GlobalConn.Table("price_table").DropTable(&PriceTable{})
	GlobalConn.Table("files").DropTable(&File{})
	GlobalConn.Table("folder").DropTable(&Folder{})
	GlobalConn.Table("sheet0").DropTable(&Sheet0{})
	GlobalConn.Table("sheet1").DropTable(&Sheet1{})
	GlobalConn.Table("sheet2").DropTable(&Sheet2{})
	GlobalConn.Table("sheet3").DropTable(&Sheet3{})
	GlobalConn.Table("sheet4").DropTable(&Sheet4{})
	GlobalConn.Table("sheet5").DropTable(&Sheet5{})
	GlobalConn.Table("sheet6").DropTable(&Sheet6{})
	// GlobalConn.DropTable(&SheetFile{})

	GlobalConn.CreateTable(&SheetFile{})
	GlobalConn.Table("project").CreateTable(&Pro{})
	GlobalConn.Table("individual").CreateTable(&Ind{})
	GlobalConn.Table("unit").CreateTable(&Unit{})
	GlobalConn.Table("price_table").CreateTable(&PriceTable{})
	GlobalConn.Table("files").CreateTable(&File{})
	GlobalConn.Table("folder").CreateTable(&Folder{})
	GlobalConn.Table("sheet0").CreateTable(&Sheet0{})
	GlobalConn.Table("sheet1").CreateTable(&Sheet1{})
	GlobalConn.Table("sheet2").CreateTable(&Sheet2{})
	GlobalConn.Table("sheet3").CreateTable(&Sheet3{})
	GlobalConn.Table("sheet4").CreateTable(&Sheet4{})
	GlobalConn.Table("sheet5").CreateTable(&Sheet5{})
	GlobalConn.Table("sheet6").CreateTable(&Sheet6{})
}

func EmptyDB() {
	GlobalConn.Table("project").Delete(Pro{})
	GlobalConn.Table("individual").Delete(Ind{})
	GlobalConn.Table("unit").Delete(Unit{})
	GlobalConn.Table("price_table").Delete(PriceTable{})
	GlobalConn.Table("sheet0").Delete(Sheet0{})
	GlobalConn.Table("sheet1").Delete(Sheet1{})
	GlobalConn.Table("sheet2").Delete(Sheet2{})
	GlobalConn.Table("sheet3").Delete(Sheet3{})
	GlobalConn.Table("sheet4").Delete(Sheet4{})
	GlobalConn.Table("sheet5").Delete(Sheet5{})
	GlobalConn.Table("sheet6").Delete(Sheet6{})
	GlobalConn.Table("files").Delete(File{})
	GlobalConn.Table("folder").Delete(Folder{})
	GlobalConn.Table("sheet_files").Delete(SheetFile{})
}
