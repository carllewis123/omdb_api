package usecase

import (

	//"encoding/base64"

	"context"
	"omdb_api/config"
	"omdb_api/models"
	"omdb_api/modules"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

var (
	logger = config.Logger()
)

type AppRepository struct {
	MysqlRepo modules.MysqlRepository
	OmdbRepo  modules.OmdbRepository
	models.UnimplementedFilmManagementServer
}

type FilmManagementServer struct {
	models.FilmManagementServer
}

func Usecase(a modules.MysqlRepository, b modules.OmdbRepository) modules.UseCase {
	return &AppRepository{
		MysqlRepo: a,
		OmdbRepo:  b,
	}
}

func (r *AppRepository) SearchByParam(req models.Parameter) interface{} {
	logger.Info(req)

	var listFilm models.ListFilmJSON
	resp := r.OmdbRepo.SearchMovie(req)

	mResp, _ := resp.(map[string]interface{})

	err := mapstructure.Decode(mResp, &listFilm)
	if err != nil {
		logger.Panic(err)
	}

	go r.MysqlRepo.InsertLogRequest("http", req, resp)

	if listFilm.Response == "True" {
		return listFilm
	} else {
		return resp
	}

}

func (r *AppRepository) GetDetailFilm(req models.Parameter) interface{} {

	var detailFilm models.DetailFilmJSON

	resp := r.OmdbRepo.GetDetailFilm(req)

	mResp, _ := resp.(map[string]interface{})

	err := mapstructure.Decode(mResp, &detailFilm)
	if err != nil {
		logger.Panic(err)
	}

	go r.MysqlRepo.InsertLogRequest("http", req, resp)
	if detailFilm.Response == "True" {
		return detailFilm
	} else {
		return resp
	}

	return detailFilm
}

func (r *AppRepository) GetDetailFilmGrpc(ctx context.Context, param *models.RequestParameter) (*models.DetailFilm, error) {
	var respDetailFilm models.DetailFilm
	var detailFilm models.DetailFilmJSON

	var request models.Parameter
	request.ApiKey = param.ApiKey
	request.ImdbID = param.ImdbID

	resp := r.OmdbRepo.GetDetailFilm(request)

	mResp, _ := resp.(map[string]interface{})

	err := mapstructure.Decode(mResp, &detailFilm)
	if err != nil {
		logger.Panic(err)
	}

	respDetailFilm.Title = detailFilm.Title
	respDetailFilm.Year = detailFilm.Year
	go r.MysqlRepo.InsertLogRequest("grpc", param, respDetailFilm)

	return &respDetailFilm, nil

}

func (r *AppRepository) SearchByParamGrpc(ctx context.Context, param *models.RequestParameter) (*models.ListFilm, error) {

	var request models.Parameter
	var respListFilm models.ListFilm
	var responseOmdb models.ListFilmJSON

	request.ApiKey = param.ApiKey
	request.ImdbID = param.ImdbID
	request.Page = param.Page
	request.Title = param.Title

	resp := r.OmdbRepo.SearchMovie(request)
	mResp, _ := resp.(map[string]interface{})

	err := mapstructure.Decode(mResp, &responseOmdb)
	if err != nil {
		logger.Panic(err)
	}

	var resultSearch []*models.DetailFilm

	for _, f := range responseOmdb.Search {
		resultSearch = append(resultSearch, &models.DetailFilm{Title: f.Title, Year: f.Year, ImdbID: f.ImdbID})
	}

	a, _ := strconv.Atoi(responseOmdb.TotalResults)
	respListFilm.TotalResult = int32(a)
	respListFilm.Response = responseOmdb.Response
	respListFilm.Search = resultSearch

	go r.MysqlRepo.InsertLogRequest("grpc", param, respListFilm)
	return &respListFilm, nil
}
