package http

import (
	"encoding/json"
	"net/http"
	"omdb_api/config"
	"omdb_api/models"
	"omdb_api/modules"

	"github.com/gorilla/mux"
)

type OmdbUsecase struct {
	UseCase modules.UseCase
}

var (
	configs, _  = config.ConfigYAML()
	urlcfgEn, _ = config.ConfigUrl()
	logger      = config.Logger()
	err         error
)

func OmdbHandler(m *mux.Router, u modules.UseCase) {
	handler := &OmdbUsecase{
		UseCase: u,
	}

	m.HandleFunc(urlcfgEn.UrlSearchByParam, handler.SearchByParam).Methods("POST")
	m.HandleFunc(urlcfgEn.UrlGetDetailFilm, handler.GetDetailFilm).Methods("POST")
}

func (e *OmdbUsecase) SearchByParam(w http.ResponseWriter, r *http.Request) {
	var request models.Parameter
	w.Header().Set("Content-type", "application/json")
	r.ParseForm()
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Panic(err)
	}
	logger.Info(request)
	val := e.UseCase.SearchByParam(request)

	json.NewEncoder(w).Encode(val)

}

func (e *OmdbUsecase) GetDetailFilm(w http.ResponseWriter, r *http.Request) {
	var request models.Parameter
	w.Header().Set("Content-type", "application/json")
	r.ParseForm()

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Panic(err)
		json.NewEncoder(w).Encode(err)
	}
	logger.Info(request)
	val := e.UseCase.GetDetailFilm(request)

	json.NewEncoder(w).Encode(val)

}
