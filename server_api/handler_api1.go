package main

import(
//	"fmt"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/abc950309/conteam-golang/core"
	"github.com/abc950309/conteam-golang/data_struct"
	"encoding/json"
)

func handler_api1(r *http.Request, output *string) {
	// process input data
	r.ParseForm()
	
	if r.Method == "GET" {
		err_input_get(r, output)
	} else if r.Form["type"] != nil && r.Form["method"] != nil {
		request_type := r.Form["type"][0]
		request_method := r.Form["method"][0]
		request_data := r.Form["data"][0]
		code, dealed_type, dealed_method, dealed_data := deal_request(request_type, request_method, request_data)
		print(code, dealed_type, dealed_method, dealed_data)
	} else {
		err_api1(r, output)
	}
}

func deal_request(raw_type string, raw_method string, raw_data string) (int, int, int, interface{}) {
	var request_type int
	var request_method int
	
	switch raw_type{
	case "contact":
		request_type = core.ConstTypeContact
	case "message":
		request_type = core.ConstTypeMessage
	default:
		request_type = core.ConstTypeErr
	}
	
	switch raw_method {
	case "insert":
		request_method = core.ConstMethodInsert
	case "delete":
		request_method = core.ConstMethodDelete
	case "update":
		request_method = core.ConstMethodUpdate
	case "get":
		request_method = core.ConstMethodGet
	case "list":
		request_method = core.ConstMethodList
	default:
		request_method = core.ConstMethodErr
	}
	
	if core.VerifyReqLogic(request_type, request_method) == false {
		return -1, -1, -1, -1
	}
	
	switch request_type {
	case core.ConstTypeContact:
		var request_data data_struct.Contact
		json.Unmarshal([]byte(raw_data), &request_data)
	case core.ConstTypeMessage:
		var request_data data_struct.Message
		json.Unmarshal([]byte(raw_data), &request_data)
	default:
		return -1, -1, -1, -1
		var request_data interface{}
	}
	
	return 0, request_type, request_method, request_data
}
