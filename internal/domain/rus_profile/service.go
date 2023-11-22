package rus_profile

import (
	"context"
	"fmt"
	"net/http"
	"rus-profile-test/internal/config"
	common "rus-profile-test/internal/domain/common/errors"
	"rus-profile-test/internal/domain/common/errors/error_type"
	"rus-profile-test/internal/domain/http_client"
	"rus-profile-test/internal/domain/profiler"
	"strings"
)

type Service struct {
	httpClient *http_client.Client
}

func NewService(cfg *config.Config) *Service {
	return &Service{httpClient: http_client.NewClient(cfg.RusProfile.BaseUrl)}
}

func (s *Service) FindByInn(ctx context.Context, inn string) (*profiler.Response, error) {
	hasCorrectlyRedirected := false
	body, _, err := s.httpClient.GetRawBody(ctx, fmt.Sprintf("/search?query=%s&search_inactive=2", inn), nil, func(req *http.Request, via []*http.Request) error {
		if strings.Contains(req.URL.Path, "/id") {
			hasCorrectlyRedirected = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if !hasCorrectlyRedirected {
		return nil, common.NewError(error_type.NotFound, fmt.Errorf("company not found"))
	}
	_, err = s.parseBody(body)
	if err != nil {
		return nil, err
	}

	return s.parseBody(body)
}

func (s *Service) parseBody(body []byte) (*profiler.Response, error) {
	var companyName, directorName, kpp []byte

	pBody := string(body)

	innI := strings.Index(pBody, "id=\"clip_inn\"")
	inn := make([]byte, 0, 8)
	if innI == -1 {
		return nil, common.NewError(error_type.NotFound, fmt.Errorf("inn not found"))
	}
	for i := innI + 14; i < len(pBody); i++ {
		if pBody[i] == '<' {
			break
		}
		inn = append(inn, pBody[i])
	}

	companyNameI := strings.Index(pBody, "legalName")
	if companyNameI != -1 {
		companyName = make([]byte, 0, 8)
		for i := companyNameI + 12; i < len(pBody); i++ {
			if pBody[i] == '<' {
				break
			}
			companyName = append(companyName, pBody[i])
		}
		companyName = []byte(strings.ReplaceAll(string(companyName), "&quot;", "\""))
	}

	directorNameI := strings.Index(pBody, "not_masked,ul_dash_main,person_ul,link")
	if directorNameI != -1 {
		directorName = make([]byte, 0, 8)
		for i := directorNameI + 46; i < len(pBody); i++ {
			if pBody[i] == '<' {
				break
			}
			directorName = append(directorName, pBody[i])
		}
	}

	kppI := strings.Index(pBody, "id=\"clip_kpp\"") + 14

	if kppI != -1 {
		kpp = make([]byte, 0, 8)
		for i := kppI; i < len(pBody); i++ {
			if pBody[i] == '<' {
				break
			}
			kpp = append(kpp, pBody[i])
		}
	}

	return &profiler.Response{
		Inn:          byteArrayToString(inn),
		Kpp:          byteArrayToString(kpp),
		CompanyName:  byteArrayToString(companyName),
		DirectorName: byteArrayToString(directorName),
	}, nil
}

func byteArrayToString(bytes []byte) string {
	if bytes != nil {
		return string(bytes)
	}
	return ""
}
