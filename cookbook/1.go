package cookbook

import "strconv"

func conv(str string) (num int64, err error) {
	num, err = strconv.ParseInt(str, 10, 64)
	return
}
