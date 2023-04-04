package Model

import (
	"fmt"
	"strconv"
	"time"
)

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
