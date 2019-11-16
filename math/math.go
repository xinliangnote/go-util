package math

import (
	"fmt"
	m "math"
	"strconv"
	"strings"
)

// 小数点后 n 位 - 四舍五入
func RoundedFixed(val float64, n int) float64 {
	shift := m.Pow(10, float64(n))
	fv := 0.0000000001 + val //对浮点数产生.xxx999999999 计算不准进行处理
	return m.Floor(fv*shift+.5) / shift
}

// 小数点后 n 位 - 舍去
func TruncRound(val float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", val)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	inst, _ := strconv.ParseFloat(newFloat, 64)
	return inst
}
