package profiler

import "context"

type Profiler interface {
	FindByInn(ctx context.Context, inn string) (*Response, error)
}

type Response struct {
	Inn          string
	Kpp          string
	CompanyName  string
	DirectorName string
}
