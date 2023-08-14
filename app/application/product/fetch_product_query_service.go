package product

import "context"

type FetchProductQueryServiceDto struct {
	ID        string
	Name      string
	Price     int64
	Stock     int
	OwnerID   string
	OwnerName string
}

type FetchProductQueryService interface {
	Run(ctx context.Context) ([]*FetchProductQueryServiceDto, error)
}