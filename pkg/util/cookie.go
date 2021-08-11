package util

import "strings"

func GetOptionsFromCookie(cookie string) map[string][]string {
	arr := strings.Split(cookie, ":")
	options := make(map[string][]string, 0)
	for _, v := range arr {
		if len(v) > 1 {
			options[string(v[0])] = strings.Split(v[1:], "")
		}
	}
	return options
}
