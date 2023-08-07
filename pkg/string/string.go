package string

import (
	"fmt"
	"github.com/thanhpk/randstr"
)

func GenerateSerial(prefix string, n int) string {
	return fmt.Sprintf("%s%s", prefix, randstr.Hex(n))
}

