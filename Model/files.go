package Model

import (
	"fmt"
	"log"
	"strconv"
)

type FilesInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GeneralFile struct {
	GId      int    `json:"gId"`
	Name     string `json:"name"`
	FatherId string `json:"fatherId"`
	ProId    int    `json:"proId"`
	Time     string `json:"time"`
	Type     string `json:"type"`
	Size     string `json:"size"`
	Url      string `json:"url"`
}

type SheetFile struct {
	SId      int    `gorm:"primary_key;type:bigint;column:sId"`
	Name     string `gorm:"text"`
	Url      string
	FileByte []byte `gorm:"type:mediumblob"`
}

func GetFileByFileId(fileId string) (file File) {
	fId, _ := strconv.Atoi(fileId)
	result := GlobalConn.Table("files").Where("file_id=?", fId).Find(&file)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return
}

func GetFileByProId(proId string) File {
	id, _ := strconv.Atoi(proId)
	file := File{}
	result := GlobalConn.Table("files").Where("pro_id = ?", id).Find(&file)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return file
}

func GetFileByFolderId(proId string, folderId string) (files []File) {
	pId, _ := strconv.Atoi(proId)
	fId, _ := strconv.Atoi(folderId)
	if folderId == "0" {
		result := GlobalConn.Table("files").Where("pro_id=?", pId).Where("folder_id=?", folderId).Order("file_name").Find(&files)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	} else {
		result := GlobalConn.Table("files").Where("folder_id=?", fId).Order("file_name").Find(&files)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
	return
}

func FormatFile(folders []Folder, files []File) (gFiles []GeneralFile) {
	for _, folder := range folders {
		gTmp := FolderToGFile(folder)
		gFiles = append(gFiles, gTmp)
	}
	for _, file := range files {
		gTmp := FileToGFile(file)
		gFiles = append(gFiles, gTmp)
	}
	return
}

//文件夹转换为通用文件格式
func FolderToGFile(folder Folder) GeneralFile {
	gFile := GeneralFile{
		GId:      folder.FolderId,
		Name:     folder.FolderName,
		FatherId: strconv.Itoa(folder.FatherFolderId),
		ProId:    folder.ProId,
		Time:     folder.Time.Format("2006-01-02 15:04:05"),
		Type:     "文件夹",
		Size:     "-",
		Url:      folder.FolderUrl,
	}
	return gFile
}

//文件转换为通用文件格式
func FileToGFile(file File) GeneralFile {
	fileLen := float32(len(file.FileByte))
	unit := []string{"B", "KB", "MB", "GB"}
	index := 0
	for fileLen/1024 > 1.0 {
		fileLen /= 1024
		index++
	}
	size := fmt.Sprintf("%.2f %s", fileLen, unit[index])
	gFile := GeneralFile{
		GId:      file.FileId,
		Name:     file.FileName,
		FatherId: strconv.Itoa(file.FolderId),
		ProId:    file.ProId,
		Type:     "文件",
		Size:     size,
		Time:     file.Time.Format("2006-01-02 15:04:05"),
		Url:      fmt.Sprintf("%sdownloadFile/%d", prefix, file.FileId),
	}
	return gFile
}

func (gFile GeneralFile) PrintGFile() {
	fmt.Printf("%#v\n", gFile)
}
