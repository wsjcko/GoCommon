package common

import (
	"github.com/speps/go-hashids"
)

/*
用来把整数生成唯一字符串（比如：通过加密解密id来隐藏真实id）
使用唯一id,生成唯一字符串订单号,唯一邀请码等唯一字符串,并可以反序列化原id
*/

var hashID *hashids.HashID

func InitHashId(minLength int, salt string) {
	hd := hashids.NewData()
	hd.Salt = salt           // 盐值,可以根据不用的业务,使用不同的盐值
	hd.MinLength = minLength // 生成唯一字符串的最小长度,注意:是最小,不是固定, >=16
	hashID = hashids.NewWithData(hd)
}

func EncodeUniqueId(uniqueId int) (string, error) {
	return hashID.Encode([]int{uniqueId}) // 参数的都是slice
}

func GetUniqueId(hashId string) (int, error) {
	result, err := hashID.DecodeWithError(hashId) //反序列化出你的原始id,也是slice类型
	return result[0], err
}
