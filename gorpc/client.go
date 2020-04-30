package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorpc/service"
	"log"
)

func TestClient() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial grpc server: %s\n", err.Error())
	}
	defer conn.Close()

	client := service.NewUserServiceClient(conn)

	userDetail, err := client.GetUserDetail(context.Background(), &service.UserDetailRequest{UserID: "usr1"})
	if err != nil {
		log.Fatalf("failed to get user detail from server: %s\n", err.Error())
	}

	fmt.Printf("successfuly fetch user detail: %s\n", userDetail.String())
}

func main() {
	TestClient()
}