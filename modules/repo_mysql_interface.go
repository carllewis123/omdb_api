package modules

type MysqlRepository interface {
	InsertLogRequest(protocol string, request interface{}, response interface{})
}
