package utils

import (
	"regexp"
)

// MatchUrl 根据指定的正则匹配
func MatchUrl(pattern, str string) string {
	reg := regexp.MustCompile(pattern)
	res := reg.FindStringSubmatch(str)
	if len(res) == 2 {
		return res[1]
	}
	return ""
}

// IncludeUrl 是否包含了URL
func IncludeUrl(msg string) bool {
	reg := regexp.MustCompile(`(https:\/\/.*?\.(com|cn|tv)).*?`)
	res := reg.FindStringSubmatch(msg)
	if len(res) > 2 {
		return true
	}
	return false
}
