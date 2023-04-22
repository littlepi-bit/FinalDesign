package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	//导入在protos文件中定义的服务
	"FinalDesign/Model"
	se "FinalDesign/protos/search"

	"google.golang.org/grpc"
)

//定义一个结构体，作用是实现helloworld中的GreeterServer
type server struct {
	//pb.UnimplementedGreeterServer
	se.UnimplementedSearcherServer
}

func (s *server) SearchProject(ctx context.Context, in *se.SearchRequest) (*se.SearchResponse, error) {
	Model.InitElasticSearch(false)
	pros := Model.SearchProjectByProName(in.Name)
	tmp, _ := json.Marshal(pros)
	return &se.SearchResponse{
		Msg:  "查询成功",
		Data: tmp,
	}, nil
}

//定义端口号，支持启动的时候输入端口号
var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	//解析输入的端口号，默认为50051
	flag.Parse()
	//TCP协议监听指定端口号
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//实例化gRPC服务
	s := grpc.NewServer()
	//服务注册
	se.RegisterSearcherServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
