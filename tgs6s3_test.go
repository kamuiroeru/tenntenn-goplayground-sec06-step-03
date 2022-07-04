package tgs6s3_test

import (
	"regexp"
	"testing"

	"github.com/kamuiroeru/tgs6s3"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUuid4(t *testing.T) {
	u4 := tgs6s3.GenerateUuid4()
	// 長さがあってる
	assert.Equal(t, 32+4, len(u4.String()))
	// '-' の位置があってる
	assert.Equal(t, "-", u4.String()[8:9])
	assert.Equal(t, "-", u4.String()[13:14])
	assert.Equal(t, "-", u4.String()[18:19])
	assert.Equal(t, "-", u4.String()[23:24])
	// 第3オクテットの 一番左 [14] が必ず 4
	for i := 0; i < 100; i++ {
		u4 := tgs6s3.GenerateUuid4()
		assert.Equal(t, "4", u4.String()[14:15])
	}
	// 第4オクテットの 一番左 [19] が必ず 8 ~ f
	r1 := regexp.MustCompile(`[89abcdef]`)
	for i := 0; i < 100; i++ {
		u4 := tgs6s3.GenerateUuid4()
		assert.Regexp(t, r1, u4.String()[19:20])
	}
	// 0 ~ f のどれかになる
	r2 := regexp.MustCompile(`[0123456789abcdef]`)
	indexes := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 14, 15, 16, 17, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35}
	for i := 0; i < 100; i++ {
		u4 := tgs6s3.GenerateUuid4()
		for _, idx := range indexes {
			assert.Regexp(t, r2, u4.String()[idx:idx+1])
		}
	}
}
