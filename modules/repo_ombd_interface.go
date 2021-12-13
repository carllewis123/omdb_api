package modules

import "omdb_api/models"

type OmdbRepository interface {
	SearchMovie(req models.Parameter) interface{}
	GetDetailFilm(req models.Parameter) interface{}
}
