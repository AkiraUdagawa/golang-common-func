package command

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func Echo(data string) {
	fmt.Println(data)
}

func Md5(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
