package grpc_service

import (
	"github.com/ahmedsharyo/AccessControlService/grpc_service/pb"

	"log"
	"testing"

	"golang.org/x/net/context"
)

func TestCountry(t *testing.T) {
	log.Println("hey")
	ctx := context.Background()
	request := pb.SecutityMangersRequest{CameraId: "1"}
	server := Server{}
	response, err := server.GetSecutityMangers(ctx, &request)
	if err != nil {
		t.Error(err)
	}

	log.Println(response)
}
