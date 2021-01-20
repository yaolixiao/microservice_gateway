package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"sync"
	"time"

	pb "github.com/yaolixiao/microservice_gateway/basic/demo/grpc-server-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var addr = flag.String("addr", "localhost:5005", "please specify address.")

const (
	timestampFormat = time.StampNano
	streamConunt    = 10
	AccessToken     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk2OTExMTQsImlzcyI6ImFwcF9pZF9iIn0.qb2A_WsDP_-jfQBxJk6L57gTnAzZs-SPLMSS_UO6Gkc"
)

func unaryCallWithMetadata(client pb.EchoClient, msg string) {
	fmt.Println("--- unaryCallWithMetadata ---")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		fmt.Println("client.UnaryEcho err=", err)
		return
	}
	fmt.Println("client response message:", res.Message)
}

func serverStreamingWithMetadata(client pb.EchoClient, msg string) {
	fmt.Println("--- server streaming ---")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.ServerStreamingEhco(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		fmt.Println("client.ServerStreamingEhco err=", err)
		return
	}

	var rpcStatus error
	fmt.Println("server stream begin:")
	for {
		res, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Println("- stream ", res.Message)
	}

	if rpcStatus != io.EOF {
		fmt.Println("not finished server stream. err=", rpcStatus)
	}
}

func clientStreamWithMetadata(client pb.EchoClient, msg string) {
	fmt.Println("--- client streaming ---")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.ClientStreamingEhco(ctx)
	if err != nil {
		fmt.Println("client.ClientStreamingEhco err=", err)
		return
	}

	for i := 0; i < streamConunt; i++ {
		if err := stream.Send(&pb.EchoRequest{Message: msg}); err != nil {
			fmt.Println("stream.Send err=", err)
			break
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("stream.CloseAndRecv err=", err)
		return
	}
	fmt.Println("client streaming response:", res.Message)
}

func bidirectionalWithMetadata(client pb.EchoClient, msg string) {
	fmt.Println("--- 双向 streaming ---")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer "+AccessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.BidirectionalStreamingecho(ctx)
	if err != nil {
		fmt.Println("client.BidirectionalStreamingecho err=", err)
		return
	}

	go func() {
		for i := 0; i < streamConunt; i++ {
			if err := stream.Send(&pb.EchoRequest{Message: msg}); err != nil {
				fmt.Println("stream.Send err=", err)
				break
			}
		}
		stream.CloseSend()
	}()

	var rpcStatus error
	for {
		res, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}

		fmt.Println("- stream.Recv:", res.Message)
	}

	if rpcStatus != io.EOF {
		fmt.Println("not finished stream. err=", rpcStatus)
	}
}

var message = "this is grpc client message."

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := grpc.Dial(*addr, grpc.WithInsecure())
			if err != nil {
				fmt.Println("grpc.Dial err=", err)
				return
			}
			defer conn.Close()

			client := pb.NewEchoClient(conn)

			// 调用一元方法
			// unaryCallWithMetadata(client, message)
			// time.Sleep(400 * time.Millisecond)

			// 服务端流式
			// serverStreamingWithMetadata(client, message)
			// time.Sleep(1 * time.Second)

			// 客户端流式
			// clientStreamWithMetadata(client, message)
			// time.Sleep(1 * time.Second)

			// 双向流式
			bidirectionalWithMetadata(client, message)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
