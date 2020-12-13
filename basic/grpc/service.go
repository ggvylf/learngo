package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ggvylf/learngo/basic/grpc/proto"
)

//定义空接口
type UserInfoService struct{}

var u = UserInfoService{}

//实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	//通过用户名查询用户信息
	name := req.Name

	//这里模拟从数据库中查询用户信息
	if name == "tom" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   12,
			Hobby: []string{"run", "play"},
		}
	}

	return

}

func main() {

	//监听端口
	addr := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen err=", err)

	}

	fmt.Println("listen ok")

	//实例化grpc
	s := grpc.NewServer()

	//在grpc上注册微服务
	pb.RegisterUserInfoServiceServer(s, &u)

	//启动服务端
	s.Serve(listener)

}
