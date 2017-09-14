package controllers

import (
	"dbselector/errdeal"
	"dbselector/hboperation"
	"dbselector/redioperation"
	"dbselector/webresp"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type User struct {
	UserName string
}

type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

var ColName []string

func ExecAction(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	t, err1 := template.ParseFiles("template/html/admin/index.html")
	if err1 != nil {
		log.Println(err1)
	}
	cookie, err := r.Cookie("DbUser")
	if err != nil || cookie.Value == "" {
		t.Execute(w, &User{"未知用户"})
		return
	}
	t.Execute(w, &User{cookie.Value})

}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	if webresp.Db != nil {
		webresp.Db.Close()
		webresp.Db = nil
	} else if redioperation.RedCon != nil {
		redioperation.RedCon.Close()
		redioperation.RedCon = nil
	} else {
		hboperation.Client.Close()
		hboperation.Client = nil
	}
	OutputJson(w, 1, "关闭数据库", nil)
}

func ExecSqlAction(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var err error
	w.Header().Set("content-type", "application/json")
	err = r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	sql := r.FormValue("sql")
	if sql == "" {
		log.Println("请输入sql语句")
		sqlEmpty := "请输入sql语句"
		_, result := hboperation.ResponseMsg("sqlInput", sqlEmpty)
		w.Write(result)
		return
	}
	log.Println("前端输入的sql语句:", sql)

	if nil != hboperation.Client {
		//操作Hbase数据库
		tmp := strings.Split(sql, ",")
		hOperation := strings.Split(tmp[0], " ")
		switch hOperation[0] {
		case "put":
			ColName = strings.Split(tmp[2], ":")
			err, result := hboperation.PutHbaseOperation(tmp[1], hOperation[1], ColName[0], ColName[1], tmp[3]) //rowkey,table,colFamily,colName,value
			errdeal.ErrDealHbOperation(hOperation[0], err)
			w.Write(result)
		case "get":
			ColName = strings.Split(tmp[2], ":")
			result, err := hboperation.GetHbaseOperation(hOperation[1], tmp[1], ColName[0], ColName[1])
			errdeal.ErrDealHbOperation(hOperation[0], err)
			w.Write(result)

		case "delete":
			ColName = strings.Split(tmp[2], ":")
			err, result := hboperation.DelHbaseOperation(hOperation[1], tmp[1], ColName[0], ColName[1], tmp[3]) //rowkey,table,colFamily,colName,value
			errdeal.ErrDeal(err)
			w.Write(result)

		default:
			err = hboperation.ChangeHbaseOperation()
			errdeal.ErrDeal(err)
			err, result := hboperation.ResponseMsg("alter", "success")
			errdeal.ErrDeal(err)
			w.Write(result)
		}
	} else if webresp.Db != nil {
		log.Println("Db值***********:", webresp.Db)
		query, err := webresp.Db.Query(sql)
		errdeal.ErrDeal(err)
		log.Println("数据库操作返回值query:", query)
		m := hboperation.ShowResult(query)
		w.Write(m)

	} else {
		RdCommand := strings.Split(sql, " ")
		RdKey := RdCommand[1]
		switch RdCommand[0] {
		case "set", "del", "sadd", "srem", "lpush", "rpush", "zrem": //lpush 往头部加;rpush 往尾部加
			RdValue := RdCommand[2]
			err, result := redioperation.RediSetAndDel(RdCommand[0], RdKey, RdValue)
			errdeal.ErrDeal(err)
			w.Write(result)
		case "get":
			err, getResult := redioperation.RediGet(RdCommand[0], RdKey)
			errdeal.ErrDeal(err)
			w.Write(getResult)
			//set 集合
		case "smembers":
			err, getResult := redioperation.RediSmembers(RdCommand[0], RdKey)
			errdeal.ErrDeal(err)
			w.Write(getResult)

		case "lrem", "lset", "zadd":
			RdListIndex := RdCommand[2]
			RdValue := RdCommand[3]
			err, result := redioperation.RediLrem(RdCommand[0], RdKey, RdListIndex, RdValue)
			errdeal.ErrDeal(err)
			w.Write(result)

		case "lrange", "zrange":
			RdListFIndex := RdCommand[2]
			RdListLIndex := RdCommand[3]
			err, getResult := redioperation.RediLrange(RdCommand[0], RdKey, RdListFIndex, RdListLIndex)
			errdeal.ErrDeal(err)
			w.Write(getResult)

		case "hkeys", "hvals", "hgetall":
			err, getResult := redioperation.RediHashGetKey(RdCommand[0], RdKey)
			errdeal.ErrDeal(err)
			w.Write(getResult)

		case "hget":
			RdKeyMember := RdCommand[2]
			err, getResult := redioperation.RediHashGet(RdCommand[0], RdKey, RdKeyMember)
			errdeal.ErrDeal(err)
			w.Write(getResult)
		case "hset":
			RdKeyMember := RdCommand[2]
			RdKeyMemberValue := RdCommand[3]
			err, result := redioperation.RediHashSet(RdCommand[0], RdKey, RdKeyMember, RdKeyMemberValue)
			errdeal.ErrDeal(err)
			w.Write(result)
		case "hdel":
			RdKeyMember := RdCommand[2]
			err, result := redioperation.RediHashDel(RdCommand[0], RdKey, RdKeyMember)
			errdeal.ErrDeal(err)
			w.Write(result)
		}
	}
}
