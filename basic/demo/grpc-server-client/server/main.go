package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/yaolixiao/microservice_gateway/basic/demo/grpc-server-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var port = flag.Int("port", 5005, "please specify port for grpc server.")

const (
	streamingCount = 10
)

type server struct{}

func (this *server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("--- UnaryEcho ---")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("missing metadata from context.")
	}
	fmt.Println("md=", md)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func (this *server) ServerStreamingEhco(req *pb.EchoRequest, stream pb.Echo_ServerStreamingEhcoServer) error {
	fmt.Println("--- ServerStreamingEhco ---")
	for i := 0; i < streamingCount; i++ {
		fmt.Printf("send message %v\n", req.Message)
		err := stream.Send(&pb.EchoResponse{Message: req.Message})
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *server) ClientStreamingEhco(stream pb.Echo_ClientStreamingEhcoServer) error {
	fmt.Println("--- ClientStreamingEcho ---")
	var message string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("echo last received message.")
			return stream.SendAndClose(&pb.EchoResponse{Message: message})
		}
		message = req.Message
		fmt.Println("request received:", req)
		if err != nil {
			return err
		}
	}
}

func (this *server) BidirectionalStreamingecho(stream pb.Echo_BidirectionalStreamingechoServer) error {
	fmt.Println("--- BidirectionalStreamingEcho ---")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("echo last received message.")
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("request received:", req)
		if err := stream.Send(&pb.EchoResponse{Message: req.Message}); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	fmt.Println("server starting at ", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	s.Serve(lis)
}
