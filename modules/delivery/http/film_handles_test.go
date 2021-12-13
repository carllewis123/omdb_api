package http

import (
	"omdb_api/models"
	"omdb_api/modules"
	"testing"
)

var UsecaseTest struct {
	modules.UseCase
}

func TestGetDetailFilm(t *testing.T) {
	var req models.Parameter
	req.ApiKey = "faf7e5bb"
	req.ImdbID = "tt4853102"
	//a := UsecaseTest.GetDetailFilm(req)

	//expect
	// if a.ImdbID != "tt4853102" {
	// 	t.Errorf("SALAH! harusnya tt4853102")
	// }
}
