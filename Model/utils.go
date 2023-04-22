package Model

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

type Measure struct {
	ProName        string
	Total          float64
	Price          string
	MeasurePrice   float64
	RulePrice      float64
	MeasurePresent float64
	RulePresent    float64
}

type Rule struct {
	ProName   string
	RulePrice float64
}

//搜索项目通过项目名
func SearchProjectByProName(proName string) []Project {
	return GlobalES.QueryByProjectName(proName)
}

//搜索项目通过单项工程名
func SearchProjectByIndName(indName string) []Project {
	return GlobalES.QueryByIndividualProjectName(indName)
}

//搜索项目通过单体工程名
func SearchProjectByUnitName(unitName string) []Project {
	return GlobalES.QueryByUnitProjectName(unitName)
}

//通过proId精确获取项目
func SearchProjectByProId(proId string) []Project {
	Id, _ := strconv.Atoi(proId)
	project := GetPorjectByProId(Id)
	return GlobalES.QueryByProjectName(project.ProjectName)
}

//根据文件名称搜索文件
func SearchFileByFileName(proId int, fileName string) []GeneralFile {
	return GlobalES.QueryByGFileName(proId, fileName)
}

//根据indId获取单项项目
func GetIndByIndId(indId string) []Ind {
	id, _ := strconv.Atoi(indId)
	inds := make([]Ind, 0)
	result := GlobalConn.Table("individual").Where("individual_id=?", id).Find(&inds)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return inds
}

//根据SId获取SheetFile
func GetSheetFileBySId(SId string) SheetFile {
	id, _ := strconv.Atoi(SId)
	f := SheetFile{}
	result := GlobalConn.Table("sheet_files").Where("sId=?", id).First(&f)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return f
}

//根据文件名删除文件
func DeleteFileByFileName(files []GeneralFile) {
	for _, file := range files {
		GlobalES.DeleteGFileByFileId(file.ProId, file.GId)
		if file.Type == "文件" {
			folderId, _ := strconv.Atoi(file.FatherId)
			UpdataFolderTime(folderId)
			GlobalConn.Where("file_id=?", file.GId).Delete(&File{})
		} else {
			UpdataFolderTime(file.GId)
			DeleteAllFile(file.GId)
		}
	}
}

//删除文件夹下所有文件
func DeleteAllFile(folderId int) {
	var files []File
	var folders []Folder
	GlobalConn.Where("folder_id=?", folderId).Find(&files)
	GlobalConn.Table("folder").Where("father_folder_id=?", folderId).Find(&folders)
	for _, file := range files {
		GlobalES.DeleteGFileByFileId(file.ProId, file.FileId)
		GlobalConn.Delete(&file)
	}
	for _, folder := range folders {
		DeleteAllFile(folder.FolderId)
	}
	var root Folder
	GlobalConn.Table("folder").Where("folder_id=?", folderId).First(&root)
	GlobalES.DeleteGFileByFileId(root.ProId, root.FolderId)
	GlobalConn.Table("folder").Delete(&root)
}

//更新文件夹的修改日期
func UpdataFolderTime(folderId int) {
	if folderId == 0 {
		return
	}
	var tmp Folder
	GlobalConn.Table("folder").Where("folder_id=?", folderId).First(&tmp)
	fmt.Printf("更新文件夹: %v\n", tmp)
	UpdataFolderTime(tmp.FatherFolderId)
	tmp.Time = time.Now()
	GlobalConn.Exec("UPDATE folder SET time=? where folder_id=?", time.Now(), folderId)
}

//搜索材料价格
func SearchMaterialPrice(materialName string, proName string) []Sheet5 {
	return GlobalES.SearchMaterialByProName(proName, materialName)
}

//搜索综合价格
func SearchGlobalPrice(globalName string, proName string) []Sheet6 {
	return GlobalES.SearchGlobalByProName(proName, globalName)
}

//搜索措施费规费
func SearchMeasurePrice(proName string) []Measure {
	pros := SearchProjectByProName(proName)
	tmpPro := make(map[string]bool)
	for _, pro := range pros {
		tmpPro[pro.ProjectName] = true
	}
	var result []Measure
	res := GlobalConn.Table("sheet2").
		Select("sheet2.pro_name,project.total_cost_lower as price, sum(col8) as measure_price").
		Joins("left join project on project.project_name=sheet2.pro_name").
		Where("sheet2.col1=?", "合　　计").Group("sheet2.pro_name").Find(&result)
	//fmt.Printf("%v\n", result)
	fmt.Println(res.Error)
	var result2 []Rule
	res = GlobalConn.Table("unit").
		Select("unit.pro_name, sum(fees) as rule_price").
		Joins("left join project on project.project_name=unit.pro_name").
		Group("unit.pro_name").Find(&result2)
	fmt.Println(res.Error)
	tmp := make(map[string]Rule)
	for _, v := range result2 {
		tmp[v.ProName] = v
	}
	ans := make([]Measure, 0)
	for _, v := range result {
		if !tmpPro[v.ProName] {
			continue
		}
		v.Total, _ = strconv.ParseFloat(v.Price[:len(v.Price)-3], 64)
		v.MeasurePresent, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", 100*v.MeasurePrice/v.Total), 64)
		v.RulePrice = tmp[v.ProName].RulePrice
		v.RulePresent, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", 100*v.RulePrice/v.Total), 64)
		fmt.Println(v)
		ans = append(ans, v)
	}
	return ans
}

//获取相似文档
func GetRelevantDoc(proName string) []DocScore {
	_, pro := GlobalES.QueryByDocName(proName)
	sort.Slice(pro.ReleventDoc, func(i, j int) bool {
		return pro.ReleventDoc[i].Score > pro.ReleventDoc[j].Score
	})
	fmt.Printf("%v\n", pro)
	return pro.ReleventDoc
}
