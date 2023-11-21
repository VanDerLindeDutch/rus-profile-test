package profiler

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	common "rus-profile-test/internal/domain/common/errors"
	"rus-profile-test/internal/domain/profiler"
	"rus-profile-test/pkg/profile_v1"
)

type serverApi struct {
	profile_v1.UnimplementedProfilerServer
	profileService profiler.Profiler
}

func Register(gRPC *grpc.Server, profileService profiler.Profiler) {
	profile_v1.RegisterProfilerServer(gRPC, &serverApi{profileService: profileService})

}

func (s *serverApi) Find(ctx context.Context, req *profile_v1.FindRequest) (*profile_v1.FindResponse, error) {
	if len([]rune(req.Inn)) != 10 {
		return nil, status.Error(codes.InvalidArgument, "inn malformed")
	}
	_, err := s.profileService.FindByInn(ctx, req.Inn)
	if err != nil {
		var domainErrType *common.Error
		if errors.As(err, &domainErrType) {
			return &profile_v1.FindResponse{
				Result: &profile_v1.FindResponse_Error{
					Error: &profile_v1.Error{
						Message: err.Error(),
						Type:    0,
					}}}, status.Error(domainErrType.GrpcCode, string(domainErrType.Code))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &profile_v1.FindResponse{
		Result: &profile_v1.FindResponse_Response{Response: &profile_v1.FindSuccessfulResponse{
			Inn:          "",
			Kpp:          "",
			CompanyName:  "",
			DirectorName: "",
		}},
	}, nil
}
