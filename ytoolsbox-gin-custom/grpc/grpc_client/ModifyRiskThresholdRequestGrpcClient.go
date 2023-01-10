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

func ModifyThreshold(mode string, dscConfig model.POCConfig) error {
	//建立链接
	conn, err := grpc.Dial("yToolsBox-req-custom:2468", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("GRPC连接建立失败: ", err)
		return err
	}
	defer conn.Close()

	// 新建grpc
	dscServiceClient := proto.NewDSCServiceClient(conn)
	//设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// 修改阈值
	response, err := dscServiceClient.ModifyRiskThreshold(ctx, &proto.ModifyRiskThresholdRequest{
		DscIp:         dscConfig.DSCAddress,
		DscFeAccount:  dscConfig.DSCWebUserName,
		DscFePassword: dscConfig.DSCWebPassword,
		ModifyMode:    mode,
	})
	if err != nil {
		log.Println("通过GRPC修改阈值失败: ", err)
		return err
	}

	log.Println(response)
	return nil
}
