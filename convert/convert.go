package convert

import "strconv"

/**
 string 转换
 */
func StrToInt64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 64)
	return i, err
}

func StrToInt32(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 32)
	return i, err
}

func StrToInt(s string) (i int, err error)  {
	i, err = strconv.Atoi(s)
	return i, err
}

func StrToFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	return f, err
}

func StrToByte(s string) []byte  {
	return []byte(s)
}

/**
 int 转换
 */
func IntToStr(i int) string  {
	return strconv.Itoa(i)
}

func IntToInt32(i int) int32 {
	return int32(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

/**
 int32 转换
 */
func Int32ToInt(i int32) int  {
	return int(i)
}

func Int32ToInt64(i int32) int64 {
	return int64(i)
}

/**
 int64 转换
 */
func Int64ToInt(i int64) int {
	return int(i)
}

func Int64ToInt32 (i int64) int32 {
	return int32(i)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}





