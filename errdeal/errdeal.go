package errdeal

import (
	"log"
)

func ErrDeal(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func ErrDealHbOperation(hbOperation string, err error) {
	if err != nil {
		log.Println(hbOperation+"操作失败：", err)
		return
	}
}
