package handlers

import (
	"encoding/json"
	"net/http"
	"unico/internal/core/domain"
	"unico/internal/core/ports"
	"unico/logger"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	marketsService ports.MarketsService
}

func NewHTTPHandler(marketsService ports.MarketsService) *HTTPHandler {
	return &HTTPHandler{
		marketsService: marketsService,
	}
}

func (hdl *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	searchType, e := domain.ValidateSearchType(vars["searchType"])
	if e != nil {
		writeResponse(w, http.StatusInternalServerError, "invalid search type")
		return
	}

	market, err := hdl.marketsService.Get(searchType, id)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, market)
}

func (hdl *HTTPHandler) Post(w http.ResponseWriter, r *http.Request) {
	var request domain.Market
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		logger.ErrorLog.Println("error when decoding request")
		writeResponse(w, http.StatusBadRequest, "error when decoding request")
		return
	} else {
		e := domain.ValidateStruct(request)
		if e != nil {
			writeResponse(w, http.StatusInternalServerError, "invalid body")
			return
		}
		appError := hdl.marketsService.Post(request)
		if appError != nil {
			writeResponse(w, http.StatusInternalServerError, appError.Error())
			return
		}
		writeResponse(w, http.StatusCreated, "created")
	}
}

func (hdl *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := hdl.marketsService.Delete(id)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, "Deleted")
}

func (hdl *HTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var request domain.Market
	err := json.NewDecoder(r.Body).Decode(&request)
	request.ID = id

	if err != nil {
		logger.ErrorLog.Println("error when decoding request")
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	} else {
		e := domain.ValidateStruct(request)
		if e != nil {
			writeResponse(w, http.StatusInternalServerError, "invalid body")
			return
		}
		market, appError := hdl.marketsService.Put(id, request)
		if appError != nil {
			writeResponse(w, http.StatusInternalServerError, appError.Error())
			return
		}
		writeResponse(w, http.StatusOK, market)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
