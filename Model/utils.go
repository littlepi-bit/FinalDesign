package Model

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

//根据文件名称搜索文件
func SearchFileByFileName(proId int, fileName string) []GeneralFile {
	return GlobalES.QueryByGFileName(proId, fileName)
}

//根据文件名删除文件
func DeleteFileByFileName(files []GeneralFile) {
	for _, file := range files {
		GlobalES.DeleteGFileByFileId(file.ProId, file.GId)
		if file.Type == "文件" {
			GlobalConn.Where("file_id=?", file.GId).Delete(&File{})
		} else {
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
