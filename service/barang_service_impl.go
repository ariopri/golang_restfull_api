package service

import (
	"context"
	"database/sql"
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/ariopri/golang_restfull_api/models/domain"
	"github.com/ariopri/golang_restfull_api/models/web"
	"github.com/ariopri/golang_restfull_api/repository"
)

type BarangServiceImpl struct {
	BarangRepository repository.BarangRepository
	DB               *sql.DB
}

func (service *BarangServiceImpl) Create(ctx context.Context, request web.BarangCreateRequest) web.BarangResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang := domain.Barang{
		Nama:  request.Nama,
		Harga: request.Harga,
		Stok:  request.Stok,
	}

	barang = service.BarangRepository.Save(ctx, tx, barang)

	return helper.ToBarangResponse(barang)
}

func (service *BarangServiceImpl) Update(ctx context.Context, request web.BarangUpdateRequest) web.BarangResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang, err := service.BarangRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	barang.Nama = request.Nama
	barang.Harga = request.Harga
	barang.Stok = request.Stok

	barang = service.BarangRepository.Update(ctx, tx, barang)

	return helper.ToBarangResponse(barang)

}

func (service *BarangServiceImpl) Delete(ctx context.Context, barangId int) {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang, err := service.BarangRepository.FindById(ctx, tx, barangId)
	helper.PanicIfError(err)

	service.BarangRepository.Delete(ctx, tx, barang)

}

func (service *BarangServiceImpl) FindById(ctx context.Context, barangId int) (web.BarangResponse, error) {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(
		err)
	defer helper.CommitOrRollback(tx)

	barang, err := service.BarangRepository.FindById(ctx, tx, barangId)
	helper.PanicIfError(err)

	return helper.ToBarangResponse(barang), nil
}

func (service *BarangServiceImpl) FindAll(ctx context.Context) []web.BarangResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barangs := service.BarangRepository.FindAll(ctx, tx)

	return helper.ToBarangResponses(barangs)
}
