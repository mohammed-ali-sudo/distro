package main

import (
	"context"
	"log"
	"net"
	"os"

	gapi "distro/internal/gen"

	"google.golang.org/grpc"
)

type server struct {
	gapi.UnimplementedExecsServiceServer
}

func (s *server) AddDrug(ctx context.Context, in *gapi.DrugOut) (*gapi.ConfirmMessage, error) {
	log.Printf("ExecsService/AddDrug received: id=%d brand=%q dose=%q", in.Id, in.BrandName, in.Dose)
	// do whatever you need here (enqueue, persist, etc.)
	return &gapi.ConfirmMessage{Confirmation: true, Message: "add received"}, nil
}

func (s *server) UpdateDrug(ctx context.Context, in *gapi.DrugOut) (*gapi.ConfirmMessage, error) {
	log.Printf("ExecsService/UpdateDrug received: id=%d brand=%q", in.Id, in.BrandName)
	return &gapi.ConfirmMessage{Confirmation: true, Message: "update received"}, nil
}

func main() {
	addr := getenv("GRPC_ADDR", ":50051")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	gapi.RegisterExecsServiceServer(s, &server{})

	log.Println("Execs gRPC listening on", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
