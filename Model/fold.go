package Model

import (
	"hash/crc32"
	"log"
	"strconv"
	"time"
)

type Folder struct {
	FolderId       int `gorm:"primary_key;type:bigint"`
	FolderName     string
	ProId          int `gorm:"type:bigint"`
	FatherFolderId int `gorm:"type:bigint"`
	Level          int
	FolderUrl      string
	Time           time.Time
}

type FolderTree struct {
	FolderId   int
	FolderName string
	Next       map[string]*FolderTree
}

var GlobalFolderTree map[int]map[string]*FolderTree

func InitFolderTree() {
	ProjectFolder := map[int][]Folder{}
	pros := []Pro{}
	result := GlobalConn.Table("project").Find(&pros)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	GlobalFolderTree = make(map[int]map[string]*FolderTree)
	for _, pro := range pros {
		proId := pro.Id
		GlobalFolderTree[proId] = make(map[string]*FolderTree)
		folders := []Folder{}
		result = GlobalConn.Table("folder").Where("pro_id=?", proId).Order("level").Find(&folders)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		ProjectFolder[proId] = folders
		index := 0
		if len(folders) == 0 {
			continue
		}
		for i := 0; i <= folders[len(folders)-1].Level; i++ {
			for folders[index].Level == i {
				node := FolderTree{
					FolderId:   folders[index].FolderId,
					FolderName: folders[index].FolderName,
					Next:       map[string]*FolderTree{},
				}
				GlobalFolderTree[proId][node.FolderName] = &node
				index++
			}
		}
	}
}

func CreateFolder(path []string, proId string) File {
	n := len(path)
	pId, _ := strconv.Atoi(proId)
	i := 0
	folderTree := GlobalFolderTree[pId]
	var preFolder *FolderTree
	for ; i < n-1; i++ {
		//todo: 用文件树替代数据查询
		flag := false
		for name, node := range folderTree {
			if name == path[i] {
				flag = true
				preFolder = node
				folderTree = node.Next
				break
			}
		}
		if !flag {
			break
		}
		// folders := []Folder{}
		// GlobalConn.Table("folder").Where("pro_id=?", pId).Where("level=?", i).Find(&folders)
		// flag := false
		// for _, f := range folders {
		// 	if f.FolderName == path[i] {
		// 		flag = true
		// 		fatherId = f.FolderId
		// 		break
		// 	}
		// }
		// if !flag {
		// 	break
		// }
	}
	for ; i < n-1; i++ {
		tmpFolder := Folder{
			FolderId:   int(crc32.ChecksumIEEE([]byte(path[i] + time.Now().String()))),
			FolderName: path[i],
			ProId:      pId,
			Level:      i,
			FolderUrl:  prefix + "Folder/",
			Time:       time.Now(),
		}
		tmpNode := &FolderTree{
			FolderId:   tmpFolder.FolderId,
			FolderName: tmpFolder.FolderName,
			Next:       map[string]*FolderTree{},
		}
		if i != 0 {
			tmpFolder.FatherFolderId = preFolder.FolderId
			preFolder.Next[tmpNode.FolderName] = tmpNode
		} else {
			GlobalFolderTree[pId][tmpNode.FolderName] = tmpNode
		}
		preFolder = tmpNode
		tmpFolder.FolderUrl += strconv.Itoa(tmpFolder.FolderId)
		result := GlobalConn.Table("folder").Create(&tmpFolder)
		if result.Error != nil {
			log.Fatal(result.Error)
		}
		//将文件夹信息添加到ES中
		GlobalES.InsertFolder(tmpFolder)
	}
	fileName := path[i]
	if result := GlobalConn.Table("files").Where("folder_id=?", preFolder.FolderId).Where("file_name=?", fileName).First(&File{}); result.Error == nil {
		log.Fatal("存在同名文件")
	}
	file := File{
		FileId:   int(crc32.ChecksumIEEE([]byte(path[i] + time.Now().String()))),
		FileName: fileName,
		ProId:    pId,
		ProName:  GetPorjectByProId(pId).ProjectName,
		FolderId: preFolder.FolderId,
	}
	return file
}

func GetFolderById(proId string, folderId string) (folder []Folder) {
	pId, _ := strconv.Atoi(proId)
	fId, _ := strconv.Atoi(folderId)
	if folderId == "0" {
		GlobalConn.Table("folder").Where("pro_id=?", pId).Where("father_folder_id=?", fId).Order("folder_name").Find(&folder)
	} else {
		GlobalConn.Table("folder").Where("father_folder_id=?", fId).Order("folder_name").Find(&folder)
	}
	return
}
