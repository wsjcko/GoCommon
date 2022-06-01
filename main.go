package main

import (
	"GoCommon/common"
	// "fmt"
)

func main() {
	// TestHashId()
}

func TestHashId() {
	common.InitHashId(16, "wsjcko")
	hashId, _ := common.EncodeUniqueId(20211211)
	println(hashId) //w3nkVX2dNeGo5rq
	uniqueId, _ := common.GetUniqueId(hashId)
	println(uniqueId) //20211211
}
