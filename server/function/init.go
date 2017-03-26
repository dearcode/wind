package function

import (
	"html/template"
)

var (
	//Common 通用函数.
	Common template.FuncMap
)

func init() {
	Common = template.FuncMap{
		"TableColumn":      TableColumn,
		"TableName":        TableName,
		"TableNameToLower": TableNameToLower,
	}
}
