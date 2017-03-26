package view

import (
	"html/template"
	"reflect"
)

//FilterInfo 自动修改规则
type FilterInfo struct {
	ID   uint64
	Name string
	Key1 string
	Key2 string
	Type int
}

//SiteInfo 对应站点表.
type SiteInfo struct {
	ID     uint64 `db_default:"auto"`
	Status bool
	Name   string
	List   struct {
		ID   uint64
		Name string
	}
	Content struct {
		ID   uint64
		Name string
	}
	Filter []struct {
		ID   uint64
		Name string
	}
	URL       string
	Md5       string
	Level     int
	AreaBegin string
	AreaEnd   string
	ListID    uint64
	ContentID uint64
	StoreID   uint64
	Ctime     string `db_default:"now()"`
	Mtime     string `db_default:"now()"`
}

//ListInfo 对应列表.
type ListInfo struct {
	ID         uint64 `db_default:"auto"`
	Name       string
	BodyBegin  string
	BodyEnd    string
	ItemBegin  string
	ItemEnd    string
	URLBegin   string `db:"url_begin"`
	URLEnd     string `db:"url_end"`
	TitleBegin string
	TitleEnd   string
	PageBegin  string
	PageEnd    string
	NextLabel  string
	Ctime      string `db_default:"now()"`
	Mtime      string `db_default:"now()"`
}

//ContentInfo 对应内容.
type ContentInfo struct {
	ID           uint64 `db_default:"auto"`
	Name         string
	BodyBegin    string
	BodyEnd      string
	TitleBegin   string
	TitleEnd     string
	DateBegin    string
	DateEnd      string
	ContentBegin string
	ContentEnd   string
	Ctime        string `db_default:"now()"`
	Mtime        string `db_default:"now()"`
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
	//WidgetTextArea 多行文本.
	WidgetTextArea
	//WidgetSelectMore 下拉列表.
	WidgetSelectMore
)

//Field 前端显示用的字段属性.
//Addible 在新建对话框中是否显示
//Modifiable 在修改对话框中是否显示
//Visible 在列表框里是否显示
//Readonly 是否可修改
type Field struct {
	Name       template.JS
	Lable      template.JS
	Relation   template.JS
	Widget     WidgetType
	Sortable   bool
	Addible    bool
	Visible    bool
	Modifiable bool
	Readonly   bool
}

//Table 前端显示用, 对应数据库中的table.
type Table struct {
	Name     template.HTML
	ID       template.JS
	Lable    string
	Fields   []Field
	Selector interface{}
	Object   interface{}
}

//GetObject 生成新Selector对象
func (t *Table) GetObject() interface{} {
	return reflect.New(reflect.TypeOf(t.Object)).Interface()
}

//GetObjectSlice 生成新Selector对象
func (t *Table) GetObjectSlice() interface{} {
	return reflect.New(reflect.SliceOf(reflect.TypeOf(t.Object))).Interface()
}
