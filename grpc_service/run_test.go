package grpc_service

import (
	"github.com/ahmedsharyo/AlertService/grpc_service/pb"

	"log"
	"testing"

	"golang.org/x/net/context"
)

func TestCountry(t *testing.T) {
	log.Println("hey")
	ctx := context.Background()
	request := pb.AlertRequest{CameraId: "1"}
	server := Server{}
	response, err := server.Alert(ctx, &request)
	if err != nil {
		t.Error(err)
	}
	if response.Alerted != true {
		t.Error("not alerted")
	}
	log.Println(response)
}
