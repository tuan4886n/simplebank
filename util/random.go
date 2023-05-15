package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// tạo số nguyên ngẫu nhiên giữa max min
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// tạo chuỗi ngẫu nhiên độ dài n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// tạo tên Owner random
func RandomOwner() string {
	return RandomString(6)
}

// tạo số tiền random
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// tạo đơn vi tiền random
func RandomCurrency() string {
	currencies := []string{USD, EUR, VND}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// tạo Email random
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
