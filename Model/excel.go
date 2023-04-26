package Model

import (
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/shakinm/xlsReader/xls"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	ExcelName string
	UserId    int
	Sheets    []Sheet
	Projects  Project
	Files     File
}

type Sheet struct {
	SheetName  string
	Title      string
	RowNum     int
	ColNum     int
	Row        []string
	Col        [][]string
	SheetFiles SheetFile
}

//总体工程
type Project struct {
	ProjectName          string
	Rows                 []string
	Cols                 []string
	IndividualProjectNum int
	IndividualProjects   []IndividualProject
}

//单项工程
type IndividualProject struct {
	IndividualProjectName string
	ProejctName           string
	UnitProjectNum        int
	SumPrice              []string
	UnitProjects          []UnitProject
	IndFileUrl            string
}

//单位工程
type UnitProject struct {
	UnitName              string
	UnitNumber            int
	ProjectName           string
	IndividualProjectName string
	Rows                  []string
	Cols                  []string
	PriceTables           []Table
}

//每个单位工程的计价表
type Table struct {
	TableName             string
	TableType             int
	UnitName              string
	IndividualProjectName string
	ProjectName           string
	TableSheet            Sheet
}

//存储的文件
type File struct {
	FileId   int `gorm:"primary_key;type:bigint"`
	FileName string
	ProId    int `gorm:"type:bigint"`
	ProName  string
	FileByte []byte `gorm:"type:mediumblob"`
	Time     time.Time
	FolderId int `gorm:"type:bigint"`
}

var ProjectRow = []string{"序号", "项目名称", "总费用（小写）", "总费用（大写）", "时间", "招标人", "造价业务类型"}
var IndividualProjectRows = []string{"序号", "单位工程名", "金额（元）", "暂估价（元）", "安全文明施工费（元）", "规费（元）"}
var UnitProjectRows = map[string][]string{
	"单位工程投标报价汇总表":  {"序号", "汇总内容", "金额（元）", "其中:暂估价（元）"},
	"单价措施项目清单与计价表": {"序号", "编码", "名称", "项目特征", "工作内容", "计量规则", "单位", "工程量", "综合单价（税前）", "人工费", "主材单价", "主材损耗", "主材费", "辅材费", "机械费", "管理费", "利润", "规费", "税金", "综合单价（含税）", "综合合价（税前）", "综合合价（含税）", "备注", "报价单位"},
	"总价措施项目清单计价表":  {"序号", " 项目编码", "项目名称", "计算基础", "费率(%)", "金额（元）", "调整费率(%)", "调整后金额(元)", "备注"},
	"其他项目清单与计价汇总表": {"序号", "项目名称", "金额(元)", "结算金额（元）", "备注"},
	"规费、税金项目计价表":   {"序号", "项目名称", "计算基础", "计算基数", "计算费率(%)", "金额(元)"},
	//" 承包人提供主要材料和工程设备一览表\r\n（适用造价信息差额调整法）": {"序号", "名称、规格、型号", "单位", "数量", "风险系数(%)", "基准单价(元)", "投标单价(元)", "发承包人确认单价(元)", "备注"},
	"主要材料价格表":      {"序号", "材料名称", "规格、型号及特殊要求", "单位", " ", "单价(元)", "备注"},
	"分部分项工程量清单计价表": {"序号", "编码", "名称", "项目特征", "工作内容", "计量规则", "供材方式", "单位", "工程量", "综合单价（税前）", "人工费", "主材单价", "主材损耗", "主材费", "辅材费", "机械费", "管理费", "利润", "规费", "税金", "综合单价（含税）", "综合合价（税前）", "综合合价（含税）", "备注", "报价单位"},
}
var TableRows = []string{
	"单位工程投标报价汇总表",
	"单价措施项目清单与计价表",
	"总价措施项目清单计价表",
	"其他项目清单与计价汇总表",
	"规费、税金项目计价表",
	" 承包人提供主要材料和工程设备一览表\r\n（适用造价信息差额调整法）",
	"分部分项工程量清单计价表",
}

func NewIndividualProject() IndividualProject {
	return IndividualProject{
		IndividualProjectName: "",
		ProejctName:           "",
		UnitProjects:          make([]UnitProject, 0),
	}
}

func NewProject() Project {
	return Project{
		Rows:               make([]string, 0),
		Cols:               make([]string, 0),
		IndividualProjects: make([]IndividualProject, 0),
	}
}

//解析分为解析xls和xlsx
func (excel *Excel) AnalyseExcel(filePath string) {
	tmp := strings.Split(filePath, ".")
	if tmp[len(tmp)-1] == "xls" {
		excel.AnalyseXls(filePath)
	} else {
		excel.AnalyseXlsx(filePath)
	}
}

func GetStrByRL(s *xls.Sheet, r, c int) string {
	row, err := s.GetRow(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	col, err := row.GetCol(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return col.GetString()
}

//解析后缀为xls的文件
func (excel *Excel) AnalyseXls(filePath string) {
	file, err := xls.OpenFile(filePath)
	if err != nil {
		log.Fatalf("open excel file err: %v", err)
	}
	sheets := []Sheet{}
	sheet0 := Sheet{
		SheetName: "封面和扉页",
		RowNum:    1,
		ColNum:    6,
		Row:       ProjectRow,
		Col:       make([][]string, 0),
	}
	s0, _ := file.GetSheet(0)
	s1, _ := file.GetSheet(1)
	sheet0.Col = append(sheet0.Col, []string{
		"",
		GetStrByRL(s0, 0, 1),
		GetStrByRL(s1, 2, 6),
		GetStrByRL(s1, 3, 6),
		GetStrByRL(s1, 11, 6),
		GetStrByRL(s1, 5, 1),
		GetStrByRL(s0, 1, 0),
	})
	sheets = append(sheets, sheet0)
	//sheet1 := Sheet{}
	for i := 2; i < file.GetNumberSheets(); i++ {
		sheet, _ := file.GetSheet(i)
		start := 0
		if strings.Contains(sheet.GetName(), "建设项目招标控制价汇总表") {
			continue
		}
		if strings.Contains(sheet.GetName(), "单项工程") {
			start = 4
		} else {
			start = 3
		}

		tmp := Sheet{
			SheetName: sheet.GetName(),
			Title:     GetStrByRL(sheet, 0, 0),
			RowNum:    int(sheet.GetNumberRows()) - start,
			ColNum:    0,
			Row:       []string{},
			Col:       make([][]string, 0),
		}
		//fmt.Println(UnitProjectRows[sheet.Row(0).Col(0)])
		if strings.Contains(tmp.SheetName, "单项工程") {
			tmp.Row = append(tmp.Row, IndividualProjectRows...)
			tmp.Title = GetStrByRL(sheet, 1, 2)
			tmp.SheetFiles = excel.GetIndFile(sheet)
		} else {
			tmp.Row = append(tmp.Row, UnitProjectRows[tmp.Title]...)
		}
		for j := start; j < int(sheet.GetNumberRows()); j++ {
			col := []string{}
			flag := strings.Contains(sheet.GetName(), "主要材料价格表")
			for k := 0; k < len(tmp.Row); k++ {
				if flag && k == 4 {
					continue
				}
				col = append(col, GetStrByRL(sheet, j, k))
			}
			tmp.Col = append(tmp.Col, col)
		}
		tmp.SheetFiles = excel.GetSheetFile(sheet)
		sheets = append(sheets, tmp)
	}
	excel.Sheets = sheets
	excel.SheetToProject()
}

//解析后缀为xlsx的文件
func (excel *Excel) AnalyseXlsx(filePath string) {

}

//获取单项工程表页
func (excel *Excel) GetIndFile(sheet *xls.Sheet) SheetFile {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"}
	sheetName := sheet.GetName()
	tmp, _ := sheet.GetRow(1)
	row := strconv.Itoa(sheet.GetNumberRows())
	f.SetSheetName("Sheet1", sheetName)
	for j := 0; j < sheet.GetNumberRows(); j++ {
		for i := 0; i < len(tmp.GetCols()); i++ {
			//fmt.Println(strconv.Itoa(i) + "=" + Model.GetStrByRL(s1, j, i))
			f.SetCellValue(sheetName, rows[i]+strconv.Itoa(j+1), GetStrByRL(sheet, j, i))
		}
	}
	f.MergeCell(sheetName, "A1", "F1")
	B2, _ := f.GetCellValue(sheetName, "B2")
	if B2 == "" {
		C2, _ := f.GetCellValue(sheetName, "C2")
		f.SetCellValue(sheetName, "B2", C2)
	}
	f.MergeCell(sheetName, "B2", "D2")
	f.MergeCell(sheetName, "D3", "f3")
	f.MergeCell(sheetName, "A"+row, "B"+row)
	f.MergeCell(sheetName, "A3", "A4")
	f.MergeCell(sheetName, "B3", "B4")
	f.MergeCell(sheetName, "C3", "C4")
	buff, err := f.WriteToBuffer()
	if err != nil {
		log.Println(err)
		return SheetFile{}
	}
	sheetFile := SheetFile{
		SId:      int(crc32.ChecksumIEEE([]byte(sheetName + time.Now().String()))),
		Name:     sheetName + ".xlsx",
		FileByte: buff.Bytes(),
	}
	sheetFile.Url = prefix + "SheetFile/" + strconv.Itoa(sheetFile.SId)
	sheetFile.SaveSheetFile()
	return sheetFile
}

func (excel *Excel) GetSheetFile(sheet *xls.Sheet) SheetFile {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB"}
	sheetName := sheet.GetName()
	tmp, _ := sheet.GetRow(2)
	//row := strconv.Itoa(sheet.GetNumberRows())
	f.SetSheetName("Sheet1", sheetName)
	for j := 0; j < sheet.GetNumberRows(); j++ {
		for i := 0; i < len(tmp.GetCols()); i++ {
			//fmt.Println(strconv.Itoa(i) + "=" + Model.GetStrByRL(s1, j, i))
			f.SetCellValue(sheetName, rows[i]+strconv.Itoa(j+1), GetStrByRL(sheet, j, i))
		}
	}
	f.MergeCell(sheetName, "A1", "F1")
	B2, _ := f.GetCellValue(sheetName, "B2")
	if B2 == "" {
		C2, _ := f.GetCellValue(sheetName, "C2")
		f.SetCellValue(sheetName, "B2", C2)
	}
	f.MergeCell(sheetName, "B2", "D2")
	buff, err := f.WriteToBuffer()
	if err != nil {
		log.Println(err)
		return SheetFile{}
	}
	sheetFile := SheetFile{
		SId:      int(crc32.ChecksumIEEE([]byte(sheetName + time.Now().String()))),
		Name:     sheetName + ".xlsx",
		FileByte: buff.Bytes(),
	}
	sheetFile.Url = prefix + "SheetFile/" + strconv.Itoa(sheetFile.SId)
	sheetFile.SaveSheetFile()
	return sheetFile
}

//显示解析结果，num表示显示的表数目
func (excel Excel) ShowExcel(num int) {
	fmt.Println("Excel name is: ", excel.ExcelName)
	for i, s := range excel.Sheets {
		if i > num {
			return
		}
		fmt.Printf("sheet%d name is: %s\n", i, s.SheetName)
		for j, _ := range s.Row {
			fmt.Printf("%10s ", s.Row[j])
		}
		fmt.Println()
		for j, _ := range s.Col {
			for k, _ := range s.Col[j] {
				fmt.Printf("%10s ", s.Col[j][k])
			}
			fmt.Println()
		}
	}
}

//将解析的Sheet表结构转换为Project结构
func (excel *Excel) SheetToProject() {
	project := NewProject()
	sheets := excel.Sheets
	project.ProjectName = sheets[0].Col[0][1]
	project.Rows = ProjectRow
	project.Cols = append(project.Cols, sheets[0].Col[0]...)
	i := 1
	for i = 1; i < len(sheets); i++ {
		if strings.Contains(sheets[i].SheetName, "单项工程") {
			individual := NewIndividualProject()
			individual.ProejctName = project.ProjectName
			individual.IndFileUrl = sheets[i].SheetFiles.Url
			tmp := []string{}
			if strings.Contains(sheets[i].Title, "\\") {
				tmp = strings.Split(sheets[i].Title, "\\")
			} else if strings.Contains(sheets[i].Title, "-") {
				tmp = strings.Split(sheets[i].Title, "-")
			} else if strings.Contains(sheets[i].Title, ")") {
				tmp = strings.Split(sheets[i].Title, ")")
			}
			individual.IndividualProjectName = tmp[1]
			individual.UnitProjectNum = len(sheets[i].Col) - 1
			fmt.Println(individual.IndividualProjectName)
			project.IndividualProjects = append(project.IndividualProjects, individual)
			project.IndividualProjectNum++
		} else {
			break
		}
	}
	for k := 0; i < len(sheets) && k < project.IndividualProjectNum; i += len(UnitProjectRows) {
		ind := project.IndividualProjects[k]
		for p := 0; p < ind.UnitProjectNum; p++ {
			unit := UnitProject{
				UnitName:              sheets[k+1].Col[p][1],
				UnitNumber:            p,
				ProjectName:           project.ProjectName,
				IndividualProjectName: ind.IndividualProjectName,
				Rows:                  IndividualProjectRows,
				Cols:                  sheets[k+1].Col[p],
			}
			for j := i; j < i+len(UnitProjectRows); j++ {
				table := Table{
					TableName:             TableRows[j-i],
					TableType:             j - i,
					UnitName:              unit.UnitName,
					ProjectName:           unit.ProjectName,
					IndividualProjectName: unit.IndividualProjectName,
					TableSheet:            sheets[j],
				}
				unit.PriceTables = append(unit.PriceTables, table)
			}
			project.IndividualProjects[k].UnitProjects = append(project.IndividualProjects[k].UnitProjects, unit)
			project.IndividualProjects[k].UnitProjectNum++
		}
		k++
	}
	excel.Projects = project
}

//显示project信息
func (excel *Excel) ShowProject() {
	fmt.Println(excel.Projects.ProjectName)
	for _, ind := range excel.Projects.IndividualProjects {
		fmt.Println(" --" + ind.IndividualProjectName)
		for _, unit := range ind.UnitProjects {
			fmt.Println("  --" + unit.UnitName)
		}
	}
}

//存储文件到mysql和Es中
func (f File) SaveFile() {
	f.Time = time.Now()
	GlobalConn.Table("files").Create(f)
	GlobalES.InsertFile(f)
}

func (sheetFile SheetFile) SaveSheetFile() {
	GlobalConn.Table("sheet_files").Create(sheetFile)
}

//存储到ElasticSearch中
func (excel *Excel) InsertElasticSearch() {
	// es := NewElasticSearch()
	// es.Init()
	GlobalES.InsertProject(excel.Projects)
	time.Sleep(time.Duration(5) * time.Second)
	GlobalES.InsertRelevance(excel.Projects)
	// es.InsertFile(excel.Files)
}
