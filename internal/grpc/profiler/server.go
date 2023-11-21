package profiler

import (
	"context"
	"google.golang.org/grpc"
	"rus-profile-test/pkg/profile_v1"
)

type serverApi struct {
	profile_v1.UnimplementedProfilerServer
}

func Register(gRPC *grpc.Server) {
	profile_v1.RegisterProfilerServer(gRPC, &serverApi{})

}

func (s *serverApi) Find(ctx context.Context, req *profile_v1.FindRequest) (*profile_v1.FindResponse, error) {
	panic("")
}
