package main

import (
	pb "FinalDesign/protos/photo"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

//定义一个结构体，作用是实现photo中的Photoer
type server struct {
	pb.UnimplementedPhotoerServer
}

//SendPhoto implements photo.Photoer
func (s *server) SendPhoto(ctx context.Context, in *pb.PhotoRequest) (*pb.PhotoReply, error) {
	return &pb.PhotoReply{
		Res: "上传图片" + in.Name + "成功",
	}, nil
}

//定义端口号，支持启动的时候输入端口号
var (
	port = flag.Int("port", 8088, "The server port")
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
	pb.RegisterPhotoerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
