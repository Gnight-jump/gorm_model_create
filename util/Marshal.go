/**
驼峰命名
*/
package util

import "strings"

func Marshal(pending string) string {
	if pending == "" {
		return ""
	}

	temp := strings.Split(pending, "_")
	var s string
	for _, v := range temp {
		chv := []rune(v)
		if len(chv) > 0 {
			if chv[0] >= 'a' && chv[0] <= 'z' { //首字母大写
				chv[0] -= 32
			}
			s += string(chv)
		}
	}
	return s
}
