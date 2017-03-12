package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//VariablePostion 变量位置.
type VariablePostion int

//Method 请求方式.
type Method int

const (
	//URI 参数在uri里.
	URI VariablePostion = iota
	//HEADER 参数在头里.
	HEADER
	//JSON 参数在body的json里.
	JSON
	//FORM 参数在form表单中.
	FORM
)

//String 类型转字符串
func (p VariablePostion) String() string {
	switch p {
	case URI:
		return "URI"
	case HEADER:
		return "HEADER"
	case JSON:
		return "JSON"
	case FORM:
		return "FORM"
	}
	return "NIL"
}

const (
	//GET http method.
	GET Method = iota
	//POST http method.
	POST
	//PUT http method.
	PUT
	//DELETE http method.
	DELETE
	//RESTful any method, may be get,post,put or delete.
	RESTful
)

//String 类型转字符串
func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case RESTful:
		return "RESTful"
	}
	return "NIL"
}

//Response 通用返回结果
type Response struct {
	Status  int
	Message string      `json:",omitempty"`
	Data    interface{} `json:",omitempty"`
}

//SendResponse 返回结果，支持json
func SendResponse(w http.ResponseWriter, status int, f string, args ...interface{}) {
	w.Header().Add("Content-Type", "application/json")
	r := Response{Status: status, Message: f}
	if len(args) > 0 {
		r.Message = fmt.Sprintf(f, args...)
	}

	buf, _ := json.Marshal(&r)
	w.Write(buf)
}

//Abort 返回结果，支持json
func Abort(w http.ResponseWriter, f string, args ...interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, f, args...)
}

//SendResponseData 返回结果，支持json
func SendResponseData(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	buf, _ := json.Marshal(&Response{Data: data})
	w.Write(buf)
}

//SendRows 为bootstrap-talbe返回结果，根据条件查找，total是总记录数，rows是数据
func SendRows(w http.ResponseWriter, total int64, data interface{}) {
	var resp = struct {
		Total int64       `json:"total"`
		Rows  interface{} `json:"rows"`
	}{total, data}

	w.Header().Add("Content-Type", "application/json")
	buf, _ := json.Marshal(resp)
	w.Write(buf)
}
