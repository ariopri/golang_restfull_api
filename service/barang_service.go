package service

import (
	"context"
	"github.com/ariopri/golang_restfull_api/models/web"
)

// Service merupakan business logic
type BarangService interface {
	Create(ctx context.Context, request web.BarangCreateRequest) web.BarangResponse
	Update(ctx context.Context, request web.BarangUpdateRequest) web.BarangResponse
	Delete(ctx context.Context, barangId int)
	FindById(ctx context.Context, barangId int) (web.BarangResponse, error)
	FindAll(ctx context.Context) []web.BarangResponse
}
