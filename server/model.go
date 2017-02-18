package server

import (
	"html/template"
)

//SiteInfo 对应站点表.
type SiteInfo struct {
	ID        uint64 `db_default:"auto"`
	Status    bool
	Name      string
	URL       string
	Md5       string
	Level     int
	AreaBegin string
	AreaEnd   string
	ListID    uint64
	ContentID uint64
	StoreID   uint64
	Ctime     string `db_defualt:"now()"`
	Mtime     string `db_defualt:"now()"`
}

//WidgetType 控件类型
type WidgetType int

const (
	//WidgetText input 文本框.
	WidgetText WidgetType = iota
	//WidgetSelect 下拉列表.
	WidgetSelect
	//WidgetRadio 单选按钮.
	WidgetRadio
)

//ViewField 前端显示用的字段属性.
//Addible 在新建对话框中是否显示
//Modifiable 在修改对话框中是否显示
//Visible 在列表框里是否显示
//Readonly 是否可修改
type ViewField struct {
	Name       template.JS
	Lable      template.JS
	Widget     WidgetType
	Enum       []string
	Sortable   bool
	Addible    bool
	Visible    bool
	Modifiable bool
	Readonly   bool
}

//ViewTable 前端显示用, 对应数据库中的table.
type ViewTable struct {
	Name   template.HTML
	Lable  string
	Fields []ViewField
}
