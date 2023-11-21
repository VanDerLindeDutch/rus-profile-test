package rus_profile

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
	common "rus-profile-test/internal/domain/common/errors"
	"rus-profile-test/internal/domain/common/errors/error_type"
	"rus-profile-test/internal/domain/http_client"
	"rus-profile-test/internal/domain/profiler"
)

type Service struct {
	httpClient *http_client.Client
}

func NewService() *Service {
	return &Service{httpClient: http_client.NewClient("https://www.rusprofile.ru")}
}

func (s *Service) FindByInn(ctx context.Context, inn string) (*profiler.Response, error) {
	hasRedirected := false
	_, _, err := s.httpClient.GetRawBody(ctx, fmt.Sprintf("/search?query=%s&search_inactive=2", inn), nil, func(req *http.Request, via []*http.Request) error {
		hasRedirected = !hasRedirected
		return nil
	})
	if err != nil {
		return nil, err
	}
	if !hasRedirected {
		return nil, common.NewError(error_type.NotFound, codes.NotFound, fmt.Errorf("company not found"))
	}

	return nil, nil
}
