package controllers

import (
	"dbselector/dbconnect"
	"dbselector/errdeal"
	"dbselector/hboperation"
	"dbselector/webresp"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/wendal/go-oci8"
	"html/template"
	"log"
	"net/http"
	"os"
)

/*
const (
	OPEN_FAILED  string = "0"
	OPEN_SUCCESS string = "1"
)
*/
/*var (
	Db *sql.DB
)
*/
type DbInfo struct {
	DbType     string
	DbIP       string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	ZkRoot     string
}

type Relete struct {
	HbaseIP  string
	HbDomain string
}

func OpenHomePage(w http.ResponseWriter, r *http.Request) {
	t, err1 := template.ParseFiles("template/html/login/index.html")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if err1 != nil {
		log.Println(err1)
	}
	t.Execute(w, nil)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	var err error
	//	var dbInfo DbInfo
	err = r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dbInfo := &DbInfo{
		r.FormValue("DbType"),
		r.FormValue("DbIP"),
		r.FormValue("DbPort"),
		r.FormValue("DbUser"),
		r.FormValue("DbPassword"),
		r.FormValue("DbName"),
		r.FormValue("ZkRoot"),
	}
	jsonText := r.FormValue("JsonText")
	if dbInfo.DbType == "0" {
		OutputJson(w, 0, "请选择要操作的数据库", nil)
		return
	}

	//配置oracle连接所需要的tnsnames.ora文件
	if dbInfo.DbType == "oracle" {
		tnsName := "/usr/lib/oracle/11.2/client64/network/admin/tnsnames.ora"

		tnsNameFile := dbInfo.DbName + " = (DESCRIPTION = (ADDRESS_LIST = (ADDRESS = (PROTOCOL = TCP)(HOST = " + dbInfo.DbIP + ")(PORT = 1521)))(CONNECT_DATA = (SERVER=DEDICATED)(SERVICE_NAME = " + dbInfo.DbName + ")))"
		//fO, err1 := os.OpenFile(tnsName, os.O_WRONLY|os.O_CREATE, 0666)
		fO, err := os.Create(tnsName)
		errdeal.ErrDeal(err)
		fO.WriteString(tnsNameFile)
	}
	/*
		dbTypeTable := []string{"mysql", "oracle", "hbase", "redis", "sqlite3"}
		for k, v := range dbTypeTable {
			if

		}
	*/
	switch dbInfo.DbType {
	case "mysql":
		errMsg := webresp.MysqlWebResp(dbInfo.DbUser, dbInfo.DbPassword, dbInfo.DbIP, dbInfo.DbPort)
		if errMsg != "" {
			OutputJson(w, 0, errMsg, nil)
			return
		}
		err = dbconnect.OpenDbMysql(dbInfo.DbIP, dbInfo.DbPort, dbInfo.DbUser, dbInfo.DbPassword, dbInfo.DbName)
		if err != nil {
			log.Println("打开mysql数据库出错：", err)
			OutputJson(w, 0, "mysql数据库连接失败", nil)
			return
		}
	case "oracle":
		errMsg := webresp.OracleWebResp(dbInfo.DbUser, dbInfo.DbPassword, dbInfo.DbName)
		if errMsg != "" {
			OutputJson(w, 0, errMsg, nil)
			return
		}
		err = dbconnect.OpenDbOracle(dbInfo.DbUser, dbInfo.DbPassword, dbInfo.DbName)
		if err != nil {
			log.Println("打开Oracle数据库失败：", err)
			OutputJson(w, 0, "oracle数据库连接失败", nil)
			return
		}
	case "hbase":
		//	openRes := openHbase(dbInfo.DbIP, dbInfo.ZkRoot)
		errMsg := webresp.HbaseWebResp(dbInfo.DbIP, dbInfo.DbPort, jsonText)
		if errMsg != "" {
			OutputJson(w, 0, errMsg, nil)
			return
		}
		err := HbaseHostsInfo(jsonText)
		errdeal.ErrDeal(err)
		//	openRes := openHbase(dbInfo.DbIP, dbInfo.DbPort, dbInfo.ZkRoot)
		openRes := dbconnect.OpenHbase(dbInfo.DbIP, dbInfo.DbPort)
		if openRes != "1" azzazz{
			log.Println("打开Hbase数据库失败")
			OutputJson(w, 0, "hbase数据库连接失败", nil)
			return
		}
	case "redis":
		errMsg := webresp.RedisWebResp(dbInfo.DbIP, dbInfo.DbPort)
		if errMsg != "" {
			OutputJson(w, 0, errMsg, nil)
			return
		}
		err = dbconnect.OpenDbRedis(dbInfo.DbIP, dbInfo.DbPort)
		if err != nil {
			log.Println("打开redis数据库失败", err)
			OutputJson(w, 0, "redis数据库连接失败", nil)
			return
		}
	default:A
		errMsg := webresp.Sqlite3WebResp(dbInfo.DbName)
		if errMsg != "" {
			OutputJson(w, 0, errMsg, nil)
			return
		}
		err = dbconnect.OpenDbSqlites(dbInfo.DbName)
		if err != nil {
			log.Println("打开sqlite数据库失败:", err)
			OutputJson(w, 0, "sqlite3数据库连接失败", nil)
			return
		}

	}

	log.Println("当前操作的数据库类型：", dbInfo.DbType)
	log.Println("当前登录用户：", dbInfo.DbUser)
	cookie := http.Cookie{Name: "DbUser", Value: dbInfo.DbUser, Path: "/"}
	http.SetCookie(w, &cookie)

	OutputJson(w, 1, "数据库连接成功", nil)

	return
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	errdeal.ErrDeal(err)
	w.Write(b)
}

func HbaseHostsInfo(jsonText string) error {
	Webinput := hboperation.Slice(jsonText)
	hostsInfo := make([]Relete, 100)
	json.Unmarshal(Webinput, &hostsInfo)
	log.Println("Hbase Host and Domain from Webinput:", string(Webinput))

	var hbaseZk string
	for i, _ := range hostsInfo {
		hbaseZk = hostsInfo[i].HbaseIP + " " + hostsInfo[i].HbDomain
		fH, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return err
		}
		fH.WriteString("\n")
		fH.WriteString(hbaseZk)
	}
	return nil
}
