package webresp

import (
	"database/sql"
)

var (
	Db *sql.DB
)

func MysqlWebResp(user, password, IP, port string) (errMsg string) {
	if user == "" || password == "" || IP == "" || port == "" {
		return "用户信息输入不完整"
	}
	return ""
}

func OracleWebResp(user, password, dbName string) (errMsg string) {
	if user == "" || password == "" || dbName == "" {
		return "用户信息输入不完整"
	}
	return ""
}

func Sqlite3WebResp(dbName string) (errMsg string) {
	if dbName == "" {
		return "请输入sqlite3数据库"
	}
	return ""
}

func HbaseWebResp(IP, Port, jsonText string) (errMsg string) {
	if IP == "" || Port == "" || jsonText == "" {
		return "IP、Port或者jsonText不能为空"
	}
	return ""

}

func RedisWebResp(IP, Port string) (errMsg string) {
	if IP == "" || Port == "" {
		return "IP或者Port不能为空"
	}
	return ""
}
