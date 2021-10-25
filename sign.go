package upiot

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
)

func md5Hex(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (c client) sign(param url.Values) string {
	keys := make([]string, 0)
	for key, _ := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	str := ""
	for _, key := range keys {
		str += fmt.Sprintf("%s=%s", key, param.Get(key))
	}

	str += c.apiSecret
	return md5Hex(str)
}
