package redioperation

import (
	"dbselector/hboperation"
	Molen "github.com/garyburd/redigo/redis"
	"log"
)

var RedCon Molen.Conn

func RediGet(rdCommand, rdKey string) (error, []byte) {

	var result hboperation.DbInfoMap
	value, err := Molen.String(RedCon.Do(rdCommand, rdKey))
	if err != nil {
		log.Println(err)
		return err, nil
	}
	result.DbInfo = make(map[int]map[string]string, 0)
	Getinfo := map[string]string{
		"Operation": "get",
		"Value":     value,
	}
	result.DbInfo[0] = Getinfo

	getResult, err3 := hboperation.ConvertToJson(result.DbInfo)
	if err3 != nil {
		log.Println(err3)
		return err3, nil
	}
	return nil, getResult
}

func RediSetAndDel(rdCommand, rdKey, rdValue string) (err error, result []byte) {
	_, err = RedCon.Do(rdCommand, rdKey, rdValue)
	if err != nil {
		log.Println(err)
		return err, nil
	} //1
	err, result = hboperation.ResponseMsg(rdCommand, "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result
}

func RediSmembers(rdCommand, rdKey string) (error, []byte) {
	var (
		result    hboperation.DbInfoMap
		Getinfo   map[string]string
		smemValue string
	)
	result.DbInfo = make(map[int]map[string]string, 0)
	value, err := Molen.Values(RedCon.Do(rdCommand, rdKey))
	if err != nil {
		log.Println(err)
		return err, nil
	}
	for i, v := range value {
		smemValue = string(v.([]byte))
		Getinfo = map[string]string{
			"Operation": rdCommand,
			"RediKey":   rdKey,
			"Value":     smemValue,
		}
		result.DbInfo[i] = Getinfo
		log.Println("获取set集合值为：", string(v.([]byte)))
	}

	getResult, err := hboperation.ConvertToJson(result.DbInfo)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, getResult
}

func RediLrem(rdCommand, rdKey, rdListIndex, rdValue string) (err error, result []byte) {
	_, err = RedCon.Do(rdCommand, rdKey, rdListIndex, rdValue)
	if err != nil {
		log.Println(err)
		return err, nil
	} //1
	err, result = hboperation.ResponseMsg(rdCommand, "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result
}

func RediLrange(rdCommand, rdKey, rdListFIndex, rdListLIndex string) (error, []byte) {
	var (
		lrangeValue string
		result      hboperation.DbInfoMap
		Getinfo     map[string]string
	)
	result.DbInfo = make(map[int]map[string]string, 0)
	values, err := Molen.Values(RedCon.Do(rdCommand, rdKey, rdListFIndex, rdListLIndex))
	if err != nil {
		log.Println(err)
		return err, nil
	} //1
	for i, v := range values {
		lrangeValue = string(v.([]byte))
		Getinfo = map[string]string{
			"Operation": rdCommand,
			"RediKey":   rdKey,
			"Value":     lrangeValue,
		}
		result.DbInfo[i] = Getinfo
		log.Println("获取list集合值为：", string(v.([]byte)))
	}

	getResult, err := hboperation.ConvertToJson(result.DbInfo)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, getResult

}

func RediHashSet(rdCommand, rdKey, rdKeyMember, rdKeyMemberValue string) (err error, result []byte) {
	_, err = RedCon.Do(rdCommand, rdKey, rdKeyMember, rdKeyMemberValue)
	if err != nil {
		log.Println(err)
		return err, nil
	} //1
	err, result = hboperation.ResponseMsg(rdCommand, "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result
}

func RediHashDel(rdCommand, rdKey, rdKeyMember string) (err error, result []byte) {
	_, err = RedCon.Do(rdCommand, rdKey, rdKeyMember)
	if err != nil {
		log.Println(err)
		return err, nil
	} //1
	err, result = hboperation.ResponseMsg(rdCommand, "success")
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, result
}

func RediHashGetKey(rdCommand, rdKey string) (error, []byte) {
	var result hboperation.DbInfoMap
	result.DbInfo = make(map[int]map[string]string, 0)
	values, _ := Molen.Values(RedCon.Do(rdCommand, rdKey))
	err, getResult := HashGetValues(values, rdCommand, rdKey)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, getResult
}

func RediHashGet(rdCommand, rdKey, rdKeyMember string) (error, []byte) {
	var result hboperation.DbInfoMap
	result.DbInfo = make(map[int]map[string]string, 0)
	values, _ := Molen.Values(RedCon.Do(rdCommand, rdKey, rdKeyMember))
	err, getResult := HashGetValues(values, rdCommand, rdKey)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, getResult
}

func HashGetValues(values []interface{}, rdCommand, rdKey string) (error, []byte) {
	var (
		hashValue string
		result    hboperation.DbInfoMap
		Getinfo   map[string]string
	)
	result.DbInfo = make(map[int]map[string]string, 0)
	for i, v := range values {
		hashValue = string(v.([]byte))
		Getinfo = map[string]string{
			"Operation": rdCommand,
			"RediKey":   rdKey,
			"Value":     hashValue,
		}
		result.DbInfo[i] = Getinfo
		log.Println("获取Hash值为：", string(v.([]byte)))
	}

	getResult, err := hboperation.ConvertToJson(result.DbInfo)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	return nil, getResult

}
