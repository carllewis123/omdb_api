package main

import (
	"net"
	_AppDelivery "omdb_api/modules/delivery/http"

	_AppMysqlRepo "omdb_api/modules/repository/databases"
	_AppOmdbRepo "omdb_api/modules/repository/rest"

	"net/http"
	"omdb_api/config"

	_AppUseCase "omdb_api/modules/usecase"

	pb "omdb_api/models"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	logger = config.Logger()
)

// type FilmManagementServer struct {
// 	pb.UnimplementedFilmManagementServer
// }

func main() {
	cfg, _ := config.ConfigUrl()
	db := config.Connection()

	muxHandler := mux.NewRouter().StrictSlash(true)

	logger.Info("HTTP Service Started")

	mysqlRepo := _AppMysqlRepo.NewMysqlUsersRepository(db)
	omdbRepo := _AppOmdbRepo.OmdbRepository()
	usecase := _AppUseCase.Usecase(mysqlRepo, omdbRepo)
	_AppDelivery.OmdbHandler(muxHandler, usecase)

	go RunServerGrpc()

	logger.Fatal(http.ListenAndServe(":"+cfg.ServicePortHttp, muxHandler))
}

func RunServerGrpc() {
	cfg, _ := config.ConfigUrl()
	db := config.Connection()
	mysqlRepo := _AppMysqlRepo.NewMysqlUsersRepository(db)
	omdbRepo := _AppOmdbRepo.OmdbRepository()

	lis, err := net.Listen("tcp", cfg.ServicePortGrpc)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	useceseParam := _AppUseCase.AppRepository{MysqlRepo: mysqlRepo, OmdbRepo: omdbRepo}
	pb.RegisterFilmManagementServer(s, &useceseParam)
	logger.Printf("GRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
