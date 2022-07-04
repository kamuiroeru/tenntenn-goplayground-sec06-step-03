package uuid

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type UuidV4 [16]byte

func GenerateUuid4() UuidV4 {
	// 乱数の値をセット
	t := time.Now().UnixNano()
	rand.Seed(t)

	// 16 byte (128 bit) のランダム値を生成
	// 参考: https://stackoverflow.com/questions/35781197/generating-a-random-fixed-length-byte-array-in-go
	uuid := make([]byte, 16)
	rand.Read(uuid)

	// 特定のビットをいじる
	// 参考: https://scrapbox.io/haraheniku/UUID_v4
	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	retUuid := [16]byte{}
	for i := 0; i < 16; i++ {
		retUuid[i] = uuid[i]
	}

	// スライス -> 固定長配列、長さが足りないとエラー
	// 参考: https://stackoverflow.com/questions/30285680/how-to-convert-slice-to-fixed-size-array
	return *(*[16]byte)(uuid)
}

func (u UuidV4) String() string {
	concated := hex.EncodeToString(u[:])
	oct1 := concated[0:8]
	oct2 := concated[8:12]
	oct3 := concated[12:16]
	oct4 := concated[16:20]
	oct5 := concated[20:32]
	return oct1 + "-" + oct2 + "-" + oct3 + "-" + oct4 + "-" + oct5
}
