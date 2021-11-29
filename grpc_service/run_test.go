package grpc_service

import (
	"log"
	"testing"

	"github.com/ahmedsharyo/AccessControlService/grpc_service/pb1"
	"golang.org/x/net/context"
)

func TestCountry(t *testing.T) {
	log.Println("hey")
	ctx := context.Background()
	request := pb1.SecutityMangersRequest{CameraId: "1"}
	server := Server{}
	response, err := server.GetSecutityMangers(ctx, &request)
	if err != nil {
		t.Error(err)
	}

	log.Println(response)
}
