package repository

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"omdb_api/config"
	"omdb_api/models"
	"omdb_api/modules"
	"os"
)

type omdbRepositoryServices struct {
	Conn *sql.DB
}

var (
	logger       = config.Logger()
	configs, _   = config.ConfigYAML()
	configUrl, _ = config.ConfigUrl()
)

func OmdbRepository() modules.OmdbRepository {
	return &omdbRepositoryServices{}
}

func (r *omdbRepositoryServices) SearchMovie(param models.Parameter) interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", configUrl.UrlOmdbSearchByParam, nil)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("apikey", param.ApiKey)
	q.Add("s", param.Title)
	q.Add("page", param.Page)

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}

	var respSeach interface{}
	//var respErrorShowBussinesInfo models.RespErrorShowBusinessInfo

	json.Unmarshal(body, &respSeach)

	return respSeach
}

func (r *omdbRepositoryServices) GetDetailFilm(param models.Parameter) interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", configUrl.UrlOmdbSearchByParam, nil)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("apikey", param.ApiKey)
	q.Add("i", param.ImdbID)

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}

	var respSeach interface{}
	json.Unmarshal(body, &respSeach)

	return respSeach

}
