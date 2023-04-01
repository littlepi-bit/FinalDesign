package Model

import (
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"time"
)

var prefix = "http://localhost:8085/"

type Pro struct {
	Id                   int `gorm:"primary_key"`
	ProjectName          string
	IndividualProjectNum int
	OrderNumber          string
	TotalCostUpper       string
	TotalCostLower       string
	Time                 string
	Tenderee             string
	BusinessType         string
	FileUrl              string `gorm:"-"`
}

type Ind struct {
	IndividualId   int `gorm:"primary_key"`
	IndividualName string
	ProId          int
	ProName        string
	UnitNum        int
}

type Unit struct {
	UnitId            int `gorm:"primary_key"`
	UnitName          string
	IndividualId      int
	IndividualName    string
	ProId             int
	ProName           string
	OrderNumber       string
	Amount            string
	Estimate          string
	SafeCivilizedCost string
	Fees              string
}

type PriceTable struct {
	TableId        int `gorm:"primary_key"`
	TableName      string
	TableType      string
	UnitId         int
	UnitName       string
	IndividualId   int
	IndividualName string
	ProId          int
	ProName        string
}

type Sheet0 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
}

type Sheet1 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
	Col6      string
	Col7      string
	Col8      string
	Col9      string
	Col10     string
	Col11     string
	Col12     string
	Col13     string
	Col14     string
	Col15     string
	Col16     string
	Col17     string
	Col18     string
	Col19     string
	Col20     string
	Col21     string
	Col22     string
	Col23     string
	Col24     string
}

type Sheet2 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
	Col6      string
	Col7      string
	Col8      string
	Col9      string
}

type Sheet3 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
}

type Sheet4 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
	Col6      string
}

type Sheet5 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
	Col6      string
	Col7      string
	Col8      string
	Col9      string
}

type Sheet6 struct {
	SheetId   int `gorm:"primary_key"`
	SheetName string
	Title     string
	SheetType string
	TableId   int
	Col1      string
	Col2      string
	Col3      string
	Col4      string
	Col5      string
	Col6      string
	Col7      string
	Col8      string
	Col9      string
	Col10     string
	Col11     string
	Col12     string
	Col13     string
	Col14     string
	Col15     string
	Col16     string
	Col17     string
	Col18     string
	Col19     string
	Col20     string
	Col21     string
	Col22     string
	Col23     string
	Col24     string
	Col25     string
}

func TestTable() {
	files := make([]File, 0)
	result := GlobalConn.Table("files").Where("pro_id=?", 3853879939).Find(&files)
	fmt.Println("查询成功")
	fmt.Println(result)
	for _, f := range files {
		fmt.Println(f.FileName)
	}
}

func (pro *Pro) ProjectToPro(project Project) {
	pro.ProjectName = project.ProjectName
	pro.TotalCostLower = project.Cols[2]
	pro.TotalCostUpper = project.Cols[3]
	pro.IndividualProjectNum = project.IndividualProjectNum
	pro.OrderNumber = project.Cols[0]
	pro.Time = project.Cols[4]
	pro.Tenderee = project.Cols[5]
	pro.BusinessType = project.Cols[6]
	pro.Id = int(crc32.ChecksumIEEE([]byte(pro.ProjectName + time.Now().String())))
}

func (ind *Ind) IndividualProjectToInd(individual IndividualProject, proId int) {
	ind.IndividualName = individual.IndividualProjectName
	ind.ProId = proId
	ind.ProName = individual.ProejctName
	ind.UnitNum = individual.UnitProjectNum
	ind.IndividualId = int(crc32.ChecksumIEEE([]byte(ind.IndividualName + ind.ProName + time.Now().String())))
}

func (unit *Unit) UnitProjectToUnit(u UnitProject, ind Ind) {
	unit.UnitName = u.UnitName
	unit.IndividualId = ind.IndividualId
	unit.IndividualName = ind.IndividualName
	unit.ProId = ind.ProId
	unit.ProName = ind.ProName
	unit.OrderNumber = u.Cols[0]
	unit.Amount = u.Cols[2]
	unit.Estimate = u.Cols[3]
	unit.SafeCivilizedCost = u.Cols[4]
	unit.Fees = u.Cols[5]
	unit.UnitId = int(crc32.ChecksumIEEE([]byte(unit.UnitName + unit.IndividualName + unit.ProName + time.Now().String())))
}

func (p *PriceTable) TabletoPriceTable(t Table, u Unit) {
	p.TableName = t.TableName
	p.TableType = strconv.Itoa(t.TableType)
	p.UnitId = u.UnitId
	p.UnitName = u.UnitName
	p.IndividualId = u.IndividualId
	p.IndividualName = u.IndividualName
	p.ProId = u.ProId
	p.ProName = u.ProName
	p.TableId = int(crc32.ChecksumIEEE([]byte(p.TableName + u.UnitName + u.IndividualName + u.ProName + time.Now().String())))
}

func (excel *Excel) InsertDB() {
	p := Pro{}
	p.ProjectToPro(excel.Projects)
	result := GlobalConn.Table("project").Create(p)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, ind := range excel.Projects.IndividualProjects {
		tmpInd := Ind{}
		tmpInd.IndividualProjectToInd(ind, p.Id)
		result = GlobalConn.Table("individual").Create(tmpInd)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		for _, u := range ind.UnitProjects {
			tmpUnit := Unit{}
			tmpUnit.UnitProjectToUnit(u, tmpInd)
			result = GlobalConn.Table("unit").Create(tmpUnit)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			for _, t := range u.PriceTables {
				tmpT := PriceTable{}
				tmpT.TabletoPriceTable(t, tmpUnit)
				result = GlobalConn.Table("price_table").Create(tmpT)
				if result.Error != nil {
					log.Fatal(result.Error)
				}
				tmpT.InserSheet(t.TableSheet)
			}
		}
	}
	excel.Files.ProId = p.Id
	excel.Files.ProName = p.ProjectName
	excel.Files.FileId = int(crc32.ChecksumIEEE([]byte(excel.Files.FileName + p.ProjectName + time.Now().String())))
	excel.Files.SaveFile()
}

func (t PriceTable) InserSheet(s Sheet) {
	n := len(s.Col)
	switch t.TableType {
	case "0":
		for i := 0; i < n; i++ {
			s0 := Sheet0{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
			}
			result := GlobalConn.Table("sheet0").Create(s0)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	case "1":
		for i := 0; i < n; i++ {
			s1 := Sheet1{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
				Col6:      s.Col[i][5],
				Col7:      s.Col[i][6],
				Col8:      s.Col[i][7],
				Col9:      s.Col[i][8],
				Col10:     s.Col[i][9],
				Col11:     s.Col[i][10],
				Col12:     s.Col[i][11],
				Col13:     s.Col[i][12],
				Col14:     s.Col[i][13],
				Col15:     s.Col[i][14],
				Col16:     s.Col[i][15],
				Col17:     s.Col[i][16],
				Col18:     s.Col[i][17],
				Col19:     s.Col[i][18],
				Col20:     s.Col[i][19],
				Col21:     s.Col[i][20],
				Col22:     s.Col[i][21],
				Col23:     s.Col[i][22],
				Col24:     s.Col[i][23],
			}
			result := GlobalConn.Table("sheet1").Create(s1)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	case "2":
		for i := 0; i < n; i++ {
			s2 := Sheet2{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
				Col6:      s.Col[i][5],
				Col7:      s.Col[i][6],
				Col8:      s.Col[i][7],
				Col9:      s.Col[i][8],
			}
			result := GlobalConn.Table("sheet2").Create(s2)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	case "3":
		for i := 0; i < n; i++ {
			s3 := Sheet3{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
			}
			result := GlobalConn.Table("sheet3").Create(s3)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	case "4":
		for i := 0; i < n; i++ {
			s4 := Sheet4{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
				Col6:      s.Col[i][5],
			}
			result := GlobalConn.Table("sheet4").Create(s4)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	case "5":
		for i := 0; i < n; i++ {
			s5 := Sheet5{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
				Col6:      s.Col[i][5],
				Col7:      s.Col[i][6],
				Col8:      s.Col[i][7],
				Col9:      s.Col[i][8],
			}
			result := GlobalConn.Table("sheet5").Create(s5)
			if result.Error != nil {
				fmt.Println(result.Error)
			}
		}
	case "6":
		for i := 0; i < n; i++ {
			s6 := Sheet6{
				SheetId:   int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName: s.SheetName,
				Title:     s.Title,
				SheetType: t.TableType,
				Col1:      s.Col[i][0],
				Col2:      s.Col[i][1],
				Col3:      s.Col[i][2],
				Col4:      s.Col[i][3],
				Col5:      s.Col[i][4],
				Col6:      s.Col[i][5],
				Col7:      s.Col[i][6],
				Col8:      s.Col[i][7],
				Col9:      s.Col[i][8],
				Col10:     s.Col[i][9],
				Col11:     s.Col[i][10],
				Col12:     s.Col[i][11],
				Col13:     s.Col[i][12],
				Col14:     s.Col[i][13],
				Col15:     s.Col[i][14],
				Col16:     s.Col[i][15],
				Col17:     s.Col[i][16],
				Col18:     s.Col[i][17],
				Col19:     s.Col[i][18],
				Col20:     s.Col[i][19],
				Col21:     s.Col[i][20],
				Col22:     s.Col[i][21],
				Col23:     s.Col[i][22],
				Col24:     s.Col[i][23],
				Col25:     s.Col[i][24],
			}
			result := GlobalConn.Table("sheet6").Create(s6)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
		}
	}
}

func GetAllProject() (pros []Pro) {
	result := GlobalConn.Table("project").Find(&pros)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	for i, _ := range pros {
		pros[i].FileUrl = fmt.Sprintf("%sdownloadProFile?proId=%d", prefix, pros[i].Id)
	}
	return
}

func GetPorjectByProId(proId int) (pro Pro) {
	result := GlobalConn.Table("project").Where("id=?", proId).Find(&pro)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}