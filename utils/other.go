package utils

import (
	"strings"
	"unicode"
)

// ToCamelCaseLower 将下划线分隔的字符串转换为全小写驼峰格式
func ToCamelCaseLower(s string) string {
	parts := strings.Split(s, "_")
	var camelCaseStr strings.Builder
	camelCaseStr.WriteRune(rune(parts[0][0]))
	for i, part := range parts {
		if i > 0 {
			camelCaseStr.WriteRune(unicode.ToUpper(rune(part[0])))
		}
		camelCaseStr.WriteString(part[1:])
	}
	return camelCaseStr.String()
}
