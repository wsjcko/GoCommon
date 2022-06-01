package common

import (
	"encoding/hex"
	"math/rand"
	"time"
)

/*
网上很多代码生成随机字符串的方法，反复调用 rand 随机函数。实际上。
rand 的随机函数本身也是调用的系统函数，开销根据操作系统而定。应尽量少调用。
本系列方法思路是，一次随机出足够多的数据，然后挨个填充。

比如，需要生成长度为 n 的字符串，我就直接使用 make 申请足够空间 b := make([]byte, n)， 然后使用 rand.Read(b[:]) 将此空间填满，本来这算是可以了，
但是由于 b 中包含各种奇奇怪怪的字符，我们需要控制一下输出内容。
因此可以根据随机生成的内容，再挨个替换掉。b[i] = letters[arc] 就是负责内容替换的。x 是原先随机生成的 byte 内容，对应的10进制范围是 0-255。
不过我们需要生成的字符个数，只要 32 或者 64 个，因此我们需要对 x 转化到 32 或者 64 以下的值，最快就是进行 & 或者 ^ 运算。
到底选用哪个都无所谓。这样我们就快速地替换了随机内容。
*/

var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")
var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=_")

func init() {
	rand.Seed(time.Now().Unix())
}

// RandLow 随机字符串，包含 1~9 和 a~z - [i,l,o]
func RandLow(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return b
}

// RandUp 随机字符串，包含 英文字母和数字附加=_两个符号
func RandUp(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 63
		b[i] = longLetters[arc]
	}
	return b
}

// RandHex 生成16进制格式的随机字符串
func RandHex(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var need int
	if n&1 == 0 { // even
		need = n
	} else { // odd
		need = n + 1
	}
	size := need / 2
	dst := make([]byte, need)
	src := dst[size:]
	if _, err := rand.Read(src[:]); err != nil {
		return []byte{}
	}
	hex.Encode(dst, src)
	return dst[:n]
}
