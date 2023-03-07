package main

import (
	"context"
	"log"
	"net"

	// path to your proto file Hello.proto
	pb "github.com/rajeshreddyt/GrpcServer/Hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3000"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) YourMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
	// implement your logic here
	return &pb.YourResponse{Name: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	reflection.Register(s)
	print("Server started at port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// $ protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/gopher/gopher.proto
