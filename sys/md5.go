package sys

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Get_md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)   //将str写入到w中

	//w.Sum(nil)将w的hash转成[]byte格式
	return fmt.Sprintf("%x", w.Sum(nil))
}

