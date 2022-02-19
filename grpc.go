package go_gopher_grpc

import (
	"context"
	"fmt"
)

type Server struct{
}

func (s *Server) Gopher(ctx context.Context, gopherRequest *GopherRequest) (*GopherReply, error){
	fmt.Println("hola El servidor a recibido la peticion")
	return &GopherReply{ Message: "holaaa" }, nil
}

func (s *Server) mustEmbedUnimplementedGopherServer(){
	fmt.Println("Received message")
}