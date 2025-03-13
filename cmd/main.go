package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/LootNex/Wallet_GRPC.git/internal/db"
	"github.com/LootNex/Wallet_GRPC.git/internal/service"
	"github.com/LootNex/Wallet_GRPC.git/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(grpcServer *grpc.Server, lis net.Listener) {
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("mistake when grpc server try to run %v", err)

		log.Println("GRPC server is running")

	}
}

func StartGateWayServer(ctx context.Context, grpcAddr, httpAddr string) {
	
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterWalletsHandlerFromEndpoint(ctx, mux, grpcAddr, opts)

	if err != nil {
		log.Fatalf("server isnot running %v", err)
	}
	
	if err = http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("mistake when http server try to run %v", err)
	}

}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := db.InitDB()
	
	if err != nil || conn == nil {
		log.Fatalf("something wrong with db connection %v", err)
	}
	defer conn.Close()


	grpcServer := grpc.NewServer()
	api.RegisterWalletsServer(grpcServer, service.NewServer(conn))

	reflection.Register(grpcServer)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go StartGRPCServer(grpcServer, lis)

	StartGateWayServer(ctx, ":50051", ":8080")

}
