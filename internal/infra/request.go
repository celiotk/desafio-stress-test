package infra

import (
	"context"
	"net/http"
)

type ServiceRequest struct {
}

func NewServiceRequest() *ServiceRequest {
	return &ServiceRequest{}
}

func (s *ServiceRequest) SendRequest(ctx context.Context, url string) (status int, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
