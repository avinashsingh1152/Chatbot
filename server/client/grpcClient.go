package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"server/client/grpcProto"
	"server/providers"
)

type SlModelGrpcClient struct {
	AppConfig providers.AppConfig
	Logger    *log.Logger
}

type ISLModelClient interface {
	GetSLMGRPCClientConn() (*grpc.ClientConn, grpcProto.ChatbotServiceClient)
	GetChatbotMessage(request string) (string, error)
}

func (slm SlModelGrpcClient) GetSLMGRPCClientConn() (*grpc.ClientConn, grpcProto.ChatbotServiceClient) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		slm.Logger.Println("Error while Connecting with grpc port: " + slm.AppConfig.GrpcSLMPort)
		panic(err)
	}

	client := grpcProto.NewChatbotServiceClient(conn)

	return conn, client
}

func (slm SlModelGrpcClient) GetChatbotMessage(request string) (string, error) {
	ctx := context.Background()

	conn, client := slm.GetSLMGRPCClientConn()
	defer conn.Close()

	resp, err := client.GetBotData(ctx, &grpcProto.ChatBotRequest{Request: request})
	if err != nil {
		slm.Logger.Println("Error while calling function GetSLMGRPCClientConn")
	}

	return resp.GetResponse(), err
}

func (slm SlModelGrpcClient) GetConnectionPing() (string, error) {
	ctx := context.Background()

	conn, client := slm.GetSLMGRPCClientConn()
	defer conn.Close()

	resp, err := client.SayHello(ctx, &grpcProto.HelloRequest{Name: "success"})
	if err != nil {
		slm.Logger.Println("Error while calling function GetSLMGRPCClientConn")
	}

	return resp.GetMessage(), err
}
