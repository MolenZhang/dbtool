package dbconnect

import (
	"database/sql"
	"dbselector/hboperation"
	"dbselector/redioperation"
	"dbselector/webresp"
	Molen "github.com/garyburd/redigo/redis"
	"github.com/tsuna/gohbase"
	"log"
)

const (
	OPEN_FAILED  string = "0"
	OPEN_SUCCESS string = "1"
)

func OpenDbMysql(ip, port, user, password, dbname string) (err error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("mysql starting...")
	webresp.Db, err = sql.Open("mysql", user+":"+password+"@tcp("+ip+":"+port+")/"+dbname+"?charset=utf8")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func OpenDbOracle(user, password, dbname string) (err error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("oracle starting...")
	webresp.Db, err = sql.Open("oci8", user+"/"+password+"@"+dbname)
	if err != nil {
		return err
	}
	return nil
}

func OpenDbSqlites(dbName string) (err error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("sqlite3 starting...")
	webresp.Db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

//func openHbase(IP, Port, zkRoot string) string {
func OpenHbase(IP, Port string) string {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("hbase starting...")
	//	zkroot := zk.ResourceName(zkRoot)
	//hboperation.Client = gohbase.NewClient(IP+":"+Port, zkroot)
	hboperation.Client = gohbase.NewClient(IP + ":" + Port)
	if hboperation.Client == nil {
		log.Println("Hbase数据库打开失败")
		return OPEN_FAILED
	}

	return OPEN_SUCCESS
}

func OpenDbRedis(IP, Port string) (err error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("redis starting...")
	redioperation.RedCon, err = Molen.Dial("tcp", IP+":"+Port)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
