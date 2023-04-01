package main

import (
	"FinalDesign/Model"
	he "FinalDesign/protos/helloworld"
	pb "FinalDesign/protos/photo"
	"FinalDesign/routers"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	port = ":8088"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()

	r := routers.RouterInit()

	r.GET("/Hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/upload", func(ctx *gin.Context) {
		//单个文件
		file, err := ctx.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		dst := "./store/files/" + file.Filename

		ctx.SaveUploadedFile(file, dst)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "文件上传成功",
		})
		// f, _ := file.Open()

		// data, _ := ioutil.ReadAll(f)

		// SendImg(pb.NewPhotoerClient(conn), data, file.Filename)
		log.Println(file.Filename)
	})

	r.GET("/chunkUpload", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.POST("/uploader/mergeFile", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.GET("/rest/n/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		req := &he.HelloRequest{Name: name}
		client := he.NewGreeterClient(conn)
		res, err := client.SayHello(context.Background(), req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	//r.Run()
	Model.OpenDatabase(false)
	Model.InitElasticSearch()
	// Model.EmptyDB()
	// Model.EmptyES()
	//Model.InitFolderTree()
	defer Model.CloseDatabase()
	if err := r.Run(":8085"); err != nil {
		fmt.Printf("Could not run server: %v", err)
	}
}

//功能函数
func SendImg(client pb.PhotoerClient, data []byte, imgName string) {
	res, err := client.SendPhoto(context.Background(), &pb.PhotoRequest{Databytes: data, Name: imgName})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res.Res)
}
