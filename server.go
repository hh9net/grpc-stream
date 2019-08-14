package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	. "./pb"
)

type server struct{}

func (s *server) Start(stream Server_StartServer) error {
	//log.Println("start")
	//stream.Send(&StartResp{
	//	Id: "asdf",
	//})
	//log.Println("send resp done")

	//go func() {
		for {
			log.Println("before recv")
			ddd, err := stream.Recv()
			if err!=nil{
				fmt.Println(err)
				return err
			}
			fmt.Println(ddd)
			log.Println("recv done")
			stream.Send(&StartResp{
				Id: "asdf"+ddd.Name,
			})
			log.Println(err)
		}
	//}()

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterServerServer(grpcServer, new(server))

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(listen)
}
