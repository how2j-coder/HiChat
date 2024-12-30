package utils

import (
	"fmt"
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
	fmt.Println(words)
	for i, word := range words {
		words[i] = cases.Lower(language.Dutch).String(word)
	}
	return strings.Join(words, "")
}

// CamelToSnake 驼峰转下划线
func CamelToSnake(camel string) string {
	var builder strings.Builder

	for i, r := range camel {
		// 如果当前字符是大写字母并且不是第一个字符，在前面添加下划线
		if unicode.IsUpper(r) && i != 0 {
			builder.WriteByte('_')
			// 将大写字母转换为小写字母
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
