package grpc_service

import (
	"log"
	"net"

	"github.com/ahmedsharyo/AccessControlService/grpc_service/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Run() {

	grpcServer := grpc.NewServer()

	var server Server

	pb.RegisterAlertingServer(grpcServer, server)

	listen, err := net.Listen("tcp", "localhost:3000")

	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}

	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}

// Server is implementation proto interface
type Server struct{}

// alert function responsible to get the Alert information
func (Server) Alert(ctx context.Context, request *pb.AlertRequest) (*pb.AlertResponse, error) {

	var data pb.AlertResponse
	data.Alerted = true
	return &data, nil
}
