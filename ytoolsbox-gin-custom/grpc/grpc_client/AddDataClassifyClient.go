package grpcclient

import (
	"context"
	"log"
	proto "main/grpc/protobuf"
	"main/model"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AddDataClassify(dscConfig model.POCConfig) error {
	//建立链接
	conn, err := grpc.Dial("yToolsBox-req-custom:2468", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial("0.0.0.0:2468", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("GRPC连接建立失败: ", err)
		return err
	}
	defer conn.Close()

	// 新建grpc
	grpcClient := proto.NewAddDataClassifyServiceClient(conn)
	//设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// 下发任务
	response, err := grpcClient.AddDataClassify(ctx, &proto.AddDataClassifyRequest{
		DscIp:         dscConfig.DSCAddress,
		DscFeAccount:  dscConfig.DSCWebUserName,
		DscFePassword: dscConfig.DSCWebPassword,
		PostgreSQLIp:  dscConfig.ToolBoxAddress,
	})
	if err != nil {
		log.Println("通过GRPC新增数据源失败: ", err)
		return err
	}

	log.Println(response)

	return nil
}
