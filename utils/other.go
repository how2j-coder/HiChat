package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

// ToCamel 驼峰命名与下划线命名
func ToCamel(s string, sep string) string {
	words := strings.Split(s, sep)
	for i, word := range words {
		words[i] = cases.Title(language.Dutch).String(word)
	}
	return strings.Join(words, "")
}