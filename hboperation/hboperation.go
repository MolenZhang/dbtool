package hboperation

import (
	"database/sql"
	"encoding/json"
	//	"github.com/lazyshot/go-hbase"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"unsafe"
)

var (
	Client gohbase.Client
)

type DbInfoMap struct {
	DbInfo map[int]map[string]string
}

type ResMsg struct {
	resMsg map[int]map[string]string
}

type ResErrMsg struct {
	resErrMsg map[int]map[string]string
}

func ShowResult(query *sql.Rows) []byte {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	column, _ := query.Columns() //[]string error
	/*
		if 0 == len(column) {
			log.Println("操纵成功")
			results1 := make(map[int]map[string]string)
			results[0] = map[string]string{
				"Operation": "operation",
				"Status":    "successfully",
			}
			b, err := ConvertToJson(results1)
			if err != nil {
				log.Println(err)
				return nil

			}
			return b
		} else {
	*/
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))

	for i := range values {
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string)

	i := 0
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			log.Println(err)
			return nil
		}

		row := make(map[string]string)
		for k, v := range values {
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}

	log.Println("sql语句执行结果：")
	for k, v := range results {
		log.Println(k, v)
	}

	b, err := ConvertToJson(results)
	if err != nil {
		log.Println(err)
		return nil

	}
	return b
	//	}
}

func ConvertToJson(results map[int]map[string]string) ([]byte, error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	b, err := json.Marshal(results)
	if err != nil {
		log.Println("查询结果转换成json失败：", err)
		return nil, err
	}
	return b, nil
}

func PutHbaseOperation(rowkey, table, colFamily, colName, value string) (err error, result []byte) {
	log.Println("rowkey:", rowkey)
	log.Println("table:", table)
	log.Println("colFamily:", colFamily)
	log.Println("colName:", colName)
	log.Println("value:", value)
	values := map[string]map[string][]byte{colFamily: map[string][]byte{colName: Slice(value)}}
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowkey, values)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	_, err = Client.Put(putRequest)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	err, result = ResponseMsg("put", "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result

}

func GetHbaseOperation(table, rowkey, colFamily, colName string) ([]byte, error) {
	log.Println(table, rowkey, colFamily, colName)
	var result DbInfoMap
	family := map[string][]string{colFamily: []string{colName}}
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey,
		hrpc.Families(family))
	if err != nil {
		log.Println(err)
	}
	getRsp, err := Client.Get(getRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println("Value:", string(getRsp.Cells[0].Value))
	log.Println("Qualifier:", string(getRsp.Cells[0].Qualifier))
	log.Println("Rowkey:", string(getRsp.Cells[0].Row))
	log.Println("Family:", string(getRsp.Cells[0].Family))
	result.DbInfo = make(map[int]map[string]string, 0)
	info := map[string]string{
		"Table":     table,
		"RowKey":    string(getRsp.Cells[0].Row),
		"ColFamily": colFamily,
		"Qualifier": colName,
		"Value":     string(getRsp.Cells[0].Value),
	}
	result.DbInfo[0] = info

	getResult, err3 := ConvertToJson(result.DbInfo)
	if err3 != nil {
		log.Println(err3)
		return nil, err3
	}

	return getResult, nil

}

func DelHbaseOperation(table, rowkey, colFamily, colName, value string) (err error, result []byte) {
	log.Println("rowkey:", rowkey)
	log.Println("table:", table)
	log.Println("colFamily:", colFamily)
	log.Println("colName:", colName)
	log.Println("value:", value)
	values := map[string]map[string][]byte{colFamily: map[string][]byte{colName: Slice(value)}}
	delRequest, err := hrpc.NewDelStr(context.Background(), table, rowkey, values)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	_, err = Client.Delete(delRequest)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	err, result = ResponseMsg("delete", "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result
}

func ChangeHbaseOperation() error {

	return nil
}
func ResponseMsg(Hbaseoperation, status string) (error, []byte) {
	var responseMsg ResMsg
	responseMsg.resMsg = make(map[int]map[string]string, 0)
	info := map[string]string{
		"Operation": Hbaseoperation,
		"Status":    status,
	}
	responseMsg.resMsg[0] = info
	resMsgToWeb, err := ConvertToJson(responseMsg.resMsg)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, resMsgToWeb
}

/*
func ResponseErrMsg() (error, []byte) {
	var responseErrMsg ResErrMsg
	return nil, nil
}
*/

func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}
