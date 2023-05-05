package Controller

import (
	"FinalDesign/Model"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ProId  string
	UserId int
}

func NewController() *Controller {
	return &Controller{}
}

//返回文件信息
func (controller *Controller) ViewAllFiles(c *gin.Context) {
	f, err := os.Open("./store/files")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatal(err)
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatal(err)
	}
	prefix := "http://localhost:8085/files/"
	filesInfo := make([]Model.FilesInfo, 0)
	for _, file := range files {
		filesInfo = append(filesInfo, Model.FilesInfo{Name: file.Name(), URL: prefix + file.Name()})
	}
	c.JSON(200, filesInfo)
}

//下载文件
// func (controller *Controller) DownloadFile(c *gin.Context) {
// 	name := c.Param("name")
// 	c.Header("Content-Length", "-1")
// 	c.File("./store/files/" + name)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "下载文件成功",
// 	})
// }

func (controller *Controller) AnalyseFile(c *gin.Context) {

}

//返回Excel信息
func (controller *Controller) ExcelDetail(c *gin.Context) {
	e := Model.Excel{
		ExcelName: "test3.xls",
		Sheets:    []Model.Sheet{},
	}
	e.AnalyseExcel("./client/" + e.ExcelName)
	c.JSON(http.StatusOK, e.Sheets)
}

//上传文件后返回解析信息
func (controller *Controller) ChunkUpload(c *gin.Context) {
	// fmt.Println(ctx.FormFile("upfile"))
	fmt.Println(c.PostForm("folder"))
	// f, _ := c.FormFile("upfile")
	files, fileHeader, _ := c.Request.FormFile("upfile")
	_, err := os.Stat("./store")
	if os.IsNotExist(err) {
		err = os.Mkdir("store", 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	PrefixPath := "./store/files"
	_, err = os.Stat(PrefixPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(PrefixPath, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer os.RemoveAll(PrefixPath)
	path := c.PostForm("relativePath")
	c.SaveUploadedFile(fileHeader, "./store/files/"+path)
	byteData := make([]byte, fileHeader.Size)
	files.Read(byteData)
	e := Model.Excel{
		ExcelName: fileHeader.Filename,
		UserId:    controller.UserId,
		Sheets:    []Model.Sheet{},
		Files: Model.File{
			FileName: fileHeader.Filename,
			FileByte: byteData,
		},
	}
	e.AnalyseExcel("./store/files/" + e.ExcelName)
	//e.ShowExcel(6)
	go func() {
		e.InsertDB()
		fmt.Println("插入Mysql成功")
		e.InsertElasticSearch()
		fmt.Println("插入ES成功")
	}()
	c.JSON(http.StatusOK, gin.H{
		"analyse": true,
		"Sheets":  e.Sheets,
		"Project": e.Projects,
	})
}

func (controller *Controller) GetAllProject(c *gin.Context) {
	pros := Model.GetAllProject()
	c.JSON(http.StatusOK, pros)
}

func (controller *Controller) GetAllProjectAndInfo(c *gin.Context) {
	pros, info := Model.GetAllProjectAndInfo()
	c.JSON(http.StatusOK, gin.H{
		"projects": pros,
		"info":     info,
	})
}

func (controller *Controller) GetInfo(c *gin.Context) {
	proName := c.Query("proName")
	info := Model.GlobalES.QueryInfoByProName(proName)
	c.JSON(http.StatusOK, info[0])
}

//获取工程信息文件
func (controller *Controller) GetProjectFile(c *gin.Context) {
	proId := c.Query("proId")
	//通过项目id获取文件
	file := Model.GetFileByProId(proId)
	c.Header("Content-Disposition", "attachment; filename="+file.FileName)
	c.Writer.Write(file.FileByte)
	c.JSON(http.StatusOK, gin.H{
		"filename": file.FileName,
	})
}

//获取单项工程文件
func (controller *Controller) GetSheetFile(c *gin.Context) {
	SId := c.Param("SId")
	file := Model.GetSheetFileBySId(SId)
	c.Header("Content-Disposition", "attachment; filename="+file.Name)
	c.Writer.Write(file.FileByte)
	c.JSON(http.StatusOK, gin.H{
		"filename": file.Name,
	})
}

//接收上传的文件夹
func (controller *Controller) UploadFolder(c *gin.Context) {
	proId := controller.ProId
	fmt.Printf("proId=%s\n" + proId)
	// f, _ := c.FormFile("upfile")
	files, fileHeader, _ := c.Request.FormFile("upfile")
	path := c.PostForm("relativePath")
	fmt.Println("相关路径为：" + path)
	// c.SaveUploadedFile(f, "./store/files/"+path)
	byteData := make([]byte, fileHeader.Size)
	files.Read(byteData)
	paths := strings.Split(path, "/")
	tmp := Model.CreateFolder(paths, proId)
	tmp.FileByte = byteData
	tmp.SaveFile()
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//获取上传文件夹信息
func (controller *Controller) GetFolderInfo(c *gin.Context) {
	proId := c.Query("proId")
	controller.ProId = proId
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//显示文件
func (controller *Controller) GetFolder(c *gin.Context) {
	proId := c.Query("proId")
	folderId := c.Param("folderId")
	if folderId == "" {
		folderId = "0"
	}
	folders := Model.GetFolderById(proId, folderId)
	files := Model.GetFileByFolderId(proId, folderId)
	gFiles := Model.FormatFile(folders, files)
	c.JSON(http.StatusOK, gFiles)
}

//下载文件
func (controller *Controller) DownloadFile(c *gin.Context) {
	fileId := c.Param("fileId")
	file := Model.GetFileByFileId(fileId)
	c.Header("Content-Disposition", "attachment; filename="+file.FileName)
	c.Writer.Write(file.FileByte)
	c.JSON(http.StatusOK, gin.H{
		"filename": file.FileName,
	})
}

//搜索项目
func (controller *Controller) SearchProject(c *gin.Context) {
	proName := c.Query("proName")
	indName := c.Query("indName")
	unitName := c.Query("unitName")
	var projects []Model.Project
	if unitName != "" {
		projects = Model.SearchProjectByUnitName(unitName)
	} else if indName != "" {
		projects = Model.SearchProjectByIndName(indName)
	} else if proName != "" {
		projects = Model.SearchProjectByProName(proName)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求为空",
		})
	}
	pros := []Model.Pro{}
	for i, _ := range projects {
		pro := Model.GetPorjectByProName(projects[i].ProjectName)
		pros = append(pros, pro)
	}
	c.JSON(http.StatusOK, pros)
}

//搜索文件
func (controller *Controller) SearchFile(c *gin.Context) {
	fileName := c.Query("fileName")
	proID, _ := strconv.Atoi(c.Query("proId"))
	gFile := Model.SearchFileByFileName(proID, fileName)
	// for _, g := range gFile {
	// 	g.PrintGFile()
	// }
	c.JSON(http.StatusOK, gFile)
}

//删除文件
func (controller *Controller) DeleteFile(c *gin.Context) {
	// gFile := []Model.GeneralFile{}
	// c.Bind(&gFile)
	// fmt.Printf("%#v\n", gFile)
	var tmp Model.Msg
	err := c.ShouldBindJSON(&tmp)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%v\n", tmp)
	gFile := tmp.FileInfo
	Model.DeleteFileByFileName(gFile)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

//通过proId获取project
func (controller *Controller) GetProject(c *gin.Context) {
	proId := c.Query("proId")
	fmt.Println("proId=" + proId)
	projects := Model.SearchProjectByProId(proId)
	c.JSON(http.StatusOK, projects[0])
}

//搜索材料价格
func (controller *Controller) SearchMaterial(c *gin.Context) {
	materialName := c.Query("materialName")
	proName := c.Query("proName")
	s := Model.SearchMaterialPrice(materialName, proName)
	c.JSON(http.StatusOK, s)
}

//搜索综合价格
func (controller *Controller) SearchGlobal(c *gin.Context) {
	globalName := c.Query("globalName")
	proName := c.Query("proName")
	s := Model.SearchGlobalPrice(globalName, proName)
	c.JSON(http.StatusOK, s)
}

//搜索措施费规费
func (controller *Controller) SearchMeasure(c *gin.Context) {
	proName := c.Query("proName")
	indName := c.Query("indName")
	s := Model.SearchMeasurePrice(proName, indName)
	c.JSON(http.StatusOK, s)
}

//获取关联项目
func (controller *Controller) GetRelevantDoc(c *gin.Context) {
	proName := c.Query("proName")
	fmt.Println("proName=" + proName)
	d := Model.GetRelevantDoc(proName)
	c.JSON(http.StatusOK, d)
}

//登录验证
func (controller *Controller) LoginCheck(c *gin.Context) {
	user := &Model.User{}
	c.Bind(&user)
	if user.Name == "" {
		log.Println("用户名不能为空")
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名不能为空",
		})
		return
	} else if user.Password == "" {
		log.Println("密码不能为空")
		c.JSON(http.StatusOK, gin.H{
			"msg": "密码不能为空",
		})
		return
	}
	result := user.CheckUser()
	fmt.Println(user)
	if !result {
		log.Println("密码或用户名错误！")
		c.JSON(http.StatusOK, gin.H{
			"msg": "密码或用户名错误！",
		})
		return
	} else {
		token := Model.GenerateToken(&Model.JWTClaims{
			UserID:   user.UId,
			Username: user.Name,
			Password: user.Password,
		})
		// Model.SetHash(
		// 	token,
		// 	Model.JsontoString(gin.H{
		// 		"userId":   user.UId,
		// 		"userName": user.Name,
		// 		"password": user.Password,
		// 	}),
		// 	time.Minute*5)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token", //你的cookie的名字
			Value:    token,   //cookie值
			Path:     "/",
			Domain:   "",
			MaxAge:   604800,
			Secure:   false,
			HttpOnly: false,
			// SameSite: 4, //下面是详细解释
		})
		controller.UserId = user.UId
		c.JSON(http.StatusOK, gin.H{
			"token":    token,
			"msg":      "ok",
			"userName": user.Name,
		})
	}
}

func (controller *Controller) SendInfo(c *gin.Context) {
	info := Model.Info{}
	c.Bind(&info)
	fmt.Println(info)
	Model.GlobalES.InsertInfo(info)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
