/**
field字段处理
*/
package generate

import (
	"strings"
)

// 获取字段json描述
func getFieldJson(field Field) string {
	return `json:"` + field.Field + `" ` + `gorm:"column:` + field.Field + `"`
}

// 获取字段说明
func getFieldComment(field Field) string {
	if len(field.Comment) > 0 {
		return "// " + field.Comment
	}
	return ""
}

// 获取字段类型
func getFiledType(field Field) string {
	typeArr := strings.Split(field.Type, "(")
	switch typeArr[0] {
	case "int":
		return "int"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float32"
	case "double":
		return "float32"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "time":
		return "time.Time"
	default:
		return "string"
	}
}
