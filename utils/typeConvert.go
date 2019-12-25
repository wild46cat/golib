package utils

import "strconv"

func SToI64(s string, defaultValue ...int64) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
	return i64
}

func SToI(s string, defaultValue ...int) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
	return res
}

func I64ToS(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IToS(i int) string {
	return strconv.Itoa(i)
}
