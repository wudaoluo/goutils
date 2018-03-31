package servDiscover

import (
	"github.com/wudaoluo/goutils/servDiscover/pb"
	"context"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type grcpserver struct{
	softname string
	softver  string
}


// SayHello implements helloworld.GreeterServer
func (s *grcpserver) Getvers(ctx context.Context, in *pb.GPRCRequest) (*pb.GPRCReply, error) {
	return &pb.GPRCReply{Softname: s.softname,Softver:s.softver}, nil
}

func GrpcSer(port ,softname,softver string) error{
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return err
	}

	g := grpc.NewServer()
	pb.RegisterGetVerServer(g, &grcpserver{softname,softver})
	// Register reflection service on gRPC server.
	reflection.Register(g)
	if err := g.Serve(lis); err != nil {
		return err
	}
	return nil
}
