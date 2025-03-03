package utils

import "strings"

// AdaptiveMysqlDsn adaptation of various mysql format dsn address
func AdaptiveMysqlDsn(dsn string) string {
	return strings.ReplaceAll(dsn, "mysql://", "")
}
