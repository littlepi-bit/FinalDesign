package routers

import (
	"FinalDesign/Controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

//跨域设置
func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}

}

//初始化gin引擎
func RouterInit() *gin.Engine {
	r := gin.Default()
	//跨域
	r.Use(cors())

	controller := Controller.NewController()
	r.GET("/files", controller.ViewAllFiles)
	// r.GET("/files/:name", controller.DownloadFile)
	r.GET("/downloadProFile", controller.GetProjectFile)
	r.GET("/downloadFile/:fileId", controller.DownloadFile)
	r.GET("/SheetFile/:SId", controller.GetSheetFile)
	r.GET("/excel", controller.ExcelDetail)
	r.GET("/getProject", controller.GetProject)
	r.GET("/getAllProject", controller.GetAllProject)
	r.GET("/getFolder/:folderId", controller.GetFolder)
	r.POST("/chunkUpload", controller.ChunkUpload)
	r.POST("/UploadFolder", controller.UploadFolder)
	r.POST("/deleteFiles", controller.DeleteFile)
	r.GET("/UploadFolder", controller.GetFolderInfo)
	r.GET("/searchProject", controller.SearchProject)
	r.GET("/searchFile", controller.SearchFile)
	r.GET("/searchMaterial", controller.SearchMaterial)
	r.GET("/searchGlobal", controller.SearchGlobal)
	r.GET("/searchMeasure", controller.SearchMeasure)
	r.GET("/getReleventDoc", controller.GetRelevantDoc)
	return r
}
