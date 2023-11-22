package grpc_profiler

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	common "rus-profile-test/internal/domain/common/errors"
	"rus-profile-test/internal/domain/common/errors/error_type"
	"rus-profile-test/internal/domain/profiler"
	"rus-profile-test/pkg/profile_v1"
	"strconv"
)

type serverApi struct {
	profile_v1.ProfilerServer
	profileService profiler.Profiler
}

func Register(gRPC *grpc.Server, profileService profiler.Profiler) {
	profile_v1.RegisterProfilerServer(gRPC, &serverApi{profileService: profileService})

}

func (s *serverApi) Find(ctx context.Context, req *profile_v1.FindRequest) (*profile_v1.FindResponse, error) {

	err := s.validateInn(req.Inn)
	if err != nil {
		return nil, err
	}

	companyProfile, err := s.profileService.FindByInn(ctx, req.Inn)
	if err != nil {
		var domainErrType *common.Error
		if errors.As(err, &domainErrType) {
			return nil, status.Error(domainErrType.Code.GetGrpcCode(), string(domainErrType.Code))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &profile_v1.FindResponse{
		Inn:          companyProfile.Inn,
		Kpp:          companyProfile.Kpp,
		CompanyName:  companyProfile.CompanyName,
		DirectorName: companyProfile.DirectorName,
	}, nil
}

func (s *serverApi) validateInn(inn string) error {
	st := status.New(codes.InvalidArgument, string(error_type.IncorrectInn))
	if len([]rune(inn)) != 10 {
		desc := "Inn length must be equal 10"
		v := &errdetails.BadRequest_FieldViolation{
			Field:       "inn",
			Description: desc,
		}
		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		st, err := st.WithDetails(br)
		if err != nil {
			return fmt.Errorf("unexpected error attaching metadata: %v", err)
		}
		return st.Err()
	}
	if _, err := strconv.Atoi(inn); err != nil {
		desc := "inn malformed"
		v := &errdetails.BadRequest_FieldViolation{
			Field:       "inn",
			Description: desc,
		}
		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)
		st, err := st.WithDetails(br)
		if err != nil {
			return fmt.Errorf("unexpected error attaching metadata: %v", err)
		}
		return st.Err()
	}
	return nil
}
