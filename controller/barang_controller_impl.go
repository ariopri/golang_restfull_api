package controller

import (
	"encoding/json"
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/ariopri/golang_restfull_api/models/web"
	"github.com/ariopri/golang_restfull_api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BarangControllerImpl struct {
	BarangService service.BarangService
}

func NewBarangController(barangService service.BarangService) BarangController {
	return &BarangControllerImpl{BarangService: barangService}
}

func (controller *BarangControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	crateBarangRequest := web.BarangCreateRequest{}
	helper.ReadFromRequestBody(r, &crateBarangRequest)

	barangResponse := controller.BarangService.Create(r.Context(), crateBarangRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   barangResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BarangControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	createBarangUpdate := web.BarangUpdateRequest{}
	helper.ReadFromRequestBody(r, &createBarangUpdate)

	barangId := params.ByName("BarangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)
	createBarangUpdate.Id = id

	barangResponse := controller.BarangService.Update(r.Context(), createBarangUpdate)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   barangResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BarangControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	barangId := params.ByName("BarangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)

	controller.BarangService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *BarangControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	barangId := params.ByName("BarangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)

	barangResponse, err := controller.BarangService.FindById(r.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   barangResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BarangControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	barangResponses := controller.BarangService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   barangResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}
