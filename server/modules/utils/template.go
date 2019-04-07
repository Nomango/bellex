package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

var jsVersion int64

func staticURLFor(value string) string {
	length := len(value)
	if value[length-1] == byte('/') {
		value = value[:length]
	}

	value = value + "?jsv=" + strconv.FormatInt(jsVersion, 10)
	return value
}

func init() {
	jsVersion = time.Now().Unix()

	// Register template functions.
	beego.AddFuncMap("StaticUrlFor", staticURLFor)
}
