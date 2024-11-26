package entity

import "context"

type ServiceRequestInterface interface {
	SendRequest(ctx context.Context, url string) (status int, err error)
}
