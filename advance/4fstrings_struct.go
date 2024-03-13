package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

type (
	CountMeta struct {
		Id         int64  `db:"id"`
		BusinessId int64  `db:"business_id"`
		Count      int64  `db:"count"`
		Types      string `db:"types"`
	}
)

var (
	countMetaFieldNames          = builder.RawFieldNames(&CountMeta{})
	countMetaRowsWithPlaceHolder = strings.Join(
		stringx.Remove(countMetaFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

func main() {

	fmt.Println(fmt.Sprintf("update %s set %s where `id` = ?", "`users`", countMetaRowsWithPlaceHolder))
	// update `users` set `business_id`=?,`count`=?,`types`=? where `id` = ?
}
