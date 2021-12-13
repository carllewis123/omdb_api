package repository

import (
	"database/sql"
	"fmt"
	"omdb_api/config"
	"omdb_api/modules"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

var (
	logger     = config.Logger()
	configs, _ = config.ConfigYAML()
)

func NewMysqlUsersRepository(Conn *sql.DB) modules.MysqlRepository {
	return &mysqlUserRepository{Conn}
}

func (r *mysqlUserRepository) InsertLogRequest(protocol string, request interface{}, response interface{}) {

	var err error
	_, err = r.Conn.Exec("insert  into `log` values (?,?,?,now())",
		protocol, fmt.Sprintf("%v", request), fmt.Sprintf("%v", response),
	)
	if err != nil {
		logger.Error(err)
	}
}
