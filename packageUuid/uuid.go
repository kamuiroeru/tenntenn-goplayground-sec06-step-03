package uuid

import (
	"math/rand"
	"strings"
	"time"
)

var iToHex string = "0123456789abcdef"

func Uuid4() string {
	// 乱数の値をセット
	t := time.Now().UnixNano()
	rand.Seed(t)

	// uuid を構築する
	arr := [32]uint8{}
	for i := 0; i < 32; i++ {
		arr[i] = uint8(rand.Intn(16))
	}
	arr[16] = 4

	// string にする
	uuidStr := make([]string, 0, 32)
	for _, v := range arr {
		uuidStr = append(uuidStr, iToHex[v:v+1])
	}

	oct1 := strings.Join(uuidStr[0:8], "")
	oct2 := strings.Join(uuidStr[8:12], "")
	oct3 := strings.Join(uuidStr[12:16], "")
	oct4 := strings.Join(uuidStr[16:20], "")
	oct5 := strings.Join(uuidStr[20:32], "")

	return oct1 + "-" + oct2 + "-" + oct3 + "-" + oct4 + "-" + oct5

}
