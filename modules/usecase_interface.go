package modules

import (
	"context"
	"omdb_api/models"
)

type UseCase interface {
	SearchByParam(req models.Parameter) interface{}
	GetDetailFilm(req models.Parameter) interface{}

	//GRPC
	SearchByParamGrpc(ctx context.Context, param *models.RequestParameter) (*models.ListFilm, error)
	GetDetailFilmGrpc(ctx context.Context, in *models.RequestParameter) (*models.DetailFilm, error)
}
