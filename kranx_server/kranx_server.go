/*
Приложение №1
Представляет из себя CRD сервер работающий по протоколу gRPC
*/

package main

import (
	"./storage"
	"fmt"
	"golang.org/x/net/context"
	pb "./kranxapi"

	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"

)
const(
	SERVER_PORT = ":54321"
)

type server struct{}

//Функционал сервера
func (s *server) AddData(ctx context.Context, input *pb.AddRequest) (*pb.AddResponse, error) {
	result := storage.AddToMemory(input.GetKey(), input.GetValue());
	fmt.Println("Get AddData request with key: " + input.GetKey())
	return &pb.AddResponse{Message: result}, nil
}
func (s *server) GetData(ctx context.Context, input *pb.GetRequest) (*pb.GetResponse, error) {
	result := storage.GetFromMemory(input.GetKey())
	fmt.Println("Get GetData request with key: " + input.GetKey())
	return &pb.GetResponse{Value: result}, nil
}
func (s *server) DelData(ctx context.Context, input *pb.DelRequest) (*pb.DelResponse, error) {
	result := storage.DelFromMemory(input.GetKey())
	fmt.Println("Get DelData request with key: " + input.GetKey())
	return &pb.DelResponse{Message: result}, nil

}

func main() {


	lis, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to listen(open) port: %v", err)
	} else {
		fmt.Println("Server started at " + SERVER_PORT + " port!")
	}
	s := grpc.NewServer()
	pb.RegisterKranxApiServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	/*
	storage.AddToMemory("1", "first")
	storage.AddToMemory("2", "A")
	storage.AddToMemory("3", "@")
	storage.AddToMemory("4", "@")
	storage.AddToMemory("5", "@")
	storage.AddToMemory("6", "@")
	storage.AddToMemory("7", "@")
	storage.AddToMemory("8", "@")
	storage.AddToMemory("9", "@")
	fmt.Println(storage.AddToMemory("10", "@"))
	storage.PrintMemory()
	fmt.Println(storage.AddToMemory("11", "@"))
	storage.PrintMemory()
	fmt.Println(storage.AddToMemory("12", "@"))
	storage.PrintMemory()
	fmt.Println(storage.AddToMemory("13", "@"))
	storage.PrintMemory()
	fmt.Println(storage.GetFromMemory("344"))
	*/

}
