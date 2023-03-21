package helper

import (
	"github.com/ariopri/golang_restfull_api/models/domain"
	"github.com/ariopri/golang_restfull_api/models/web"
)

func ToBarangResponse(barang domain.Barang) web.BarangResponse {
	return web.BarangResponse{
		Id:    barang.Id,
		Nama:  barang.Nama,
		Harga: barang.Harga,
		Stok:  barang.Stok,
	}
}

func ToBarangResponses(barangs []domain.Barang) []web.BarangResponse {
	var barangResponses []web.BarangResponse
	for _, barang := range barangs {
		barangResponses = append(barangResponses, ToBarangResponse(barang))
	}
	return barangResponses
}
