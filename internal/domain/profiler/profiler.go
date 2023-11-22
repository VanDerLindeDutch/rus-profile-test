package profiler

import "context"

type Profiler interface {
	FindByInn(ctx context.Context, inn string) (*Response, error)
}

type Response struct {
	Inn          string `json:"inn"`
	Kpp          string `json:"kpp"`
	CompanyName  string `json:"companyName"`
	DirectorName string `json:"directorName"`
}
