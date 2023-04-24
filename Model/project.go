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
	Id                   int `gorm:"primary_key;type:bigint"`
	UId                  int `gorm:"type:bigint;column:uid"`
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
	IndividualId   int `gorm:"primary_key;type:bigint"`
	IndividualName string
	ProId          int `gorm:"type:bigint"`
	ProName        string
	UnitNum        int
	IndFileUrl     string
}

type Unit struct {
	UnitId            int `gorm:"primary_key;type:bigint"`
	UnitName          string
	IndividualId      int `gorm:"type:bigint"`
	IndividualName    string
	ProId             int `gorm:"type:bigint"`
	ProName           string
	OrderNumber       string
	Amount            string
	Estimate          string
	SafeCivilizedCost string
	Fees              float64
}

type PriceTable struct {
	TableId        int `gorm:"primary_key;type:bigint"`
	TableName      string
	TableType      string
	UnitId         int `gorm:"type:bigint"`
	UnitName       string
	IndividualId   int `gorm:"type:bigint"`
	IndividualName string
	ProId          int `gorm:"type:bigint"`
	ProName        string
	TableFileUrl   string
}

type Sheet0 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	SheetFileUrl string
}

type Sheet1 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string
	Col6         string
	Col7         string
	Col8         string
	Col9         string
	Col10        string
	Col11        string
	Col12        string
	Col13        string
	Col14        string
	Col15        string
	Col16        string
	Col17        string
	Col18        string
	Col19        string
	Col20        string
	Col21        string
	Col22        string
	Col23        string
	Col24        string
	SheetFileUrl string
}

type Sheet2 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string
	Col6         float64
	Col7         float64
	Col8         float64
	Col9         string
	SheetFileUrl string
}

type Sheet3 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string
	SheetFileUrl string
}

type Sheet4 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string
	Col6         float64
	SheetFileUrl string
}

type Sheet5 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string
	Col6         string
	Col7         string
	Col8         string
	Col9         string
	SheetFileUrl string
}

type Sheet6 struct {
	SheetId      int `gorm:"primary_key;type:bigint"`
	SheetName    string
	Title        string
	SheetType    string
	ProId        int `gorm:"type:bigint"`
	ProName      string
	TableId      int `gorm:"type:bigint"`
	Col1         string
	Col2         string
	Col3         string
	Col4         string
	Col5         string `gorm:"type:text"`
	Col6         string
	Col7         string `gorm:"type:text"`
	Col8         string
	Col9         string
	Col10        string
	Col11        string
	Col12        string
	Col13        string
	Col14        string
	Col15        string
	Col16        string
	Col17        string
	Col18        string
	Col19        string
	Col20        string
	Col21        string
	Col22        string
	Col23        string
	Col24        string
	Col25        string
	SheetFileUrl string
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

func (pro *Pro) ProjectToPro(project Project, userId int) {
	pro.ProjectName = project.ProjectName
	pro.UId = userId
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
	ind.IndFileUrl = individual.IndFileUrl
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
	unit.Fees, _ = strconv.ParseFloat(u.Cols[5], 64)
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
	p.TableFileUrl = t.TableSheet.SheetFiles.Url
	p.TableId = int(crc32.ChecksumIEEE([]byte(p.TableName + u.UnitName + u.IndividualName + u.ProName + time.Now().String())))
}

func (excel *Excel) InsertDB() {
	p := Pro{}
	p.ProjectToPro(excel.Projects, excel.UserId)
	result := GlobalConn.Table("project").Create(p)
	GlobalFolderTree[p.Id] = make(map[string]*FolderTree)
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
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet0").Create(s0)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet0(s0)
		}
	case "1":
		for i := 0; i < n; i++ {
			s1 := Sheet1{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				Col6:         s.Col[i][5],
				Col7:         s.Col[i][6],
				Col8:         s.Col[i][7],
				Col9:         s.Col[i][8],
				Col10:        s.Col[i][9],
				Col11:        s.Col[i][10],
				Col12:        s.Col[i][11],
				Col13:        s.Col[i][12],
				Col14:        s.Col[i][13],
				Col15:        s.Col[i][14],
				Col16:        s.Col[i][15],
				Col17:        s.Col[i][16],
				Col18:        s.Col[i][17],
				Col19:        s.Col[i][18],
				Col20:        s.Col[i][19],
				Col21:        s.Col[i][20],
				Col22:        s.Col[i][21],
				Col23:        s.Col[i][22],
				Col24:        s.Col[i][23],
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet1").Create(s1)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet1(s1)
		}
	case "2":
		for i := 0; i < n; i++ {
			tmp1, _ := strconv.ParseFloat(s.Col[i][5], 64)
			tmp2, _ := strconv.ParseFloat(s.Col[i][6], 64)
			tmp3, _ := strconv.ParseFloat(s.Col[i][7], 64)
			s2 := Sheet2{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				Col6:         tmp1,
				Col7:         tmp2,
				Col8:         tmp3,
				Col9:         s.Col[i][8],
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet2").Create(s2)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet2(s2)
		}
	case "3":
		for i := 0; i < n; i++ {
			s3 := Sheet3{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet3").Create(s3)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet3(s3)
		}
	case "4":
		for i := 0; i < n; i++ {
			tmp, _ := strconv.ParseFloat(s.Col[i][5], 64)
			s4 := Sheet4{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				Col6:         tmp,
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet4").Create(s4)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet4(s4)
		}
	case "5":
		for i := 0; i < n; i++ {
			s5 := Sheet5{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				SheetFileUrl: s.SheetFiles.Url,
				// Col6:      s.Col[i][5],
				// Col7:      s.Col[i][6],
				// Col8:      s.Col[i][7],
				// Col9:      s.Col[i][8],
			}
			result := GlobalConn.Table("sheet5").Create(s5)
			if result.Error != nil {
				fmt.Println(result.Error)
			}
			GlobalES.InsertSheet5(s5)
		}
	case "6":
		for i := 0; i < n; i++ {
			s6 := Sheet6{
				SheetId:      int(crc32.ChecksumIEEE([]byte(s.SheetName + s.Title + s.Col[i][1] + time.Now().String()))),
				SheetName:    s.SheetName,
				Title:        s.Title,
				SheetType:    t.TableType,
				ProId:        t.ProId,
				ProName:      t.ProName,
				TableId:      t.TableId,
				Col1:         s.Col[i][0],
				Col2:         s.Col[i][1],
				Col3:         s.Col[i][2],
				Col4:         s.Col[i][3],
				Col5:         s.Col[i][4],
				Col6:         s.Col[i][5],
				Col7:         s.Col[i][6],
				Col8:         s.Col[i][7],
				Col9:         s.Col[i][8],
				Col10:        s.Col[i][9],
				Col11:        s.Col[i][10],
				Col12:        s.Col[i][11],
				Col13:        s.Col[i][12],
				Col14:        s.Col[i][13],
				Col15:        s.Col[i][14],
				Col16:        s.Col[i][15],
				Col17:        s.Col[i][16],
				Col18:        s.Col[i][17],
				Col19:        s.Col[i][18],
				Col20:        s.Col[i][19],
				Col21:        s.Col[i][20],
				Col22:        s.Col[i][21],
				Col23:        s.Col[i][22],
				Col24:        s.Col[i][23],
				Col25:        s.Col[i][24],
				SheetFileUrl: s.SheetFiles.Url,
			}
			result := GlobalConn.Table("sheet6").Create(s6)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			GlobalES.InsertSheet6(s6)
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

//通过项目Id获取数据库中的项目信息
func GetPorjectByProId(proId int) (pro Pro) {
	result := GlobalConn.Table("project").Where("id=?", proId).Find(&pro)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

//通过项目名获取数据库中的项目信息
func GetPorjectByProName(ProName string) (pro Pro) {
	result := GlobalConn.Table("project").Where("project_name=?", ProName).First(&pro)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

//通过项目名称获取table
func GetTableByProId(proName string) (tables []PriceTable) {
	result := GlobalConn.Table("price_table").Where("pro_name=?", proName).Find(&tables)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}
