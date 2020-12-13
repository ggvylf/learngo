package main

import (
	"context"
	"fmt"

	pb "github.com/ggvylf/learngo/basic/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	//连接服务端
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn err=", err)

	}
	defer conn.Close()

	//实例化grpc客户端
	client := pb.NewUserInfoServiceClient(conn)

	//组装请求参数
	req := new(pb.UserRequest)
	req.Name = "tom"

	//调用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("resp err=", err)
	}

	//打印响应结果
	fmt.Printf("resp=%v\n", resp)
}
