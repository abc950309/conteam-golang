package main

import (
	"fmt"
	//	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/abc950309/conteam-golang/core"
	"github.com/abc950309/conteam-golang/data_struct"
	"net/http"
)

func handler_api1(r *http.Request, output *string) {
	// process input data
	r.ParseForm()

	if r.Method == "GET" {
		err_input_get(r, output)
	} else if r.Form["type"] != nil && r.Form["method"] != nil {
		request_token := r.Form["token"][0]
		request_type := r.Form["type"][0]
		request_method := r.Form["method"][0]
		request_data := r.Form["data"][0]

		code, dealed_type, dealed_method, dealed_data := deal_request(request_type, request_method, request_data)
		if code < 0 {
			err_api1(r, output)
		} else {
			out_json, err := json.Marshal(core.Controller(request_token, dealed_type, dealed_method, dealed_data, "api1"))
			
			if err == nil {
				*output = string(out_json)
			} else {
				*output = "Error: json.Marshal Fail"
			}
		}
		fmt.Println(*output)
	} else {
		err_api1(r, output)
	}
}

func deal_request(raw_type string, raw_method string, raw_data string) (int, int, int, interface{}) {
	var request_type int
	var request_method int

	switch raw_type {
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
	
	var request_data interface{}
	
	// Json 解析
	switch request_type {
	case core.ConstTypeContact:
		switch request_method {
		case core.ConstMethodInsert:
			var request_data data_struct.Contact
		case core.ConstMethodDelete:
			var request_data data_struct.ContactFilter
		case core.ConstMethodUpdate:
			var request_data data_struct.Contact
		case core.ConstMethodGet:
			var request_data data_struct.ContactFilter
		case core.ConstMethodList:
			var request_data data_struct.ContactFilters
		}
	case core.ConstTypeMessage:
		switch request_method {
		case core.ConstMethodInsert:
			var request_data data_struct.Message
		case core.ConstMethodGet:
			var request_data data_struct.MessageFilter
		case core.ConstMethodList:
			var request_data data_struct.MessageFilters
		}
	default:
		return -1, -1, -1, -1
	}
	
	if json.Unmarshal([]byte(raw_data), &request_data) == nil {
		return 0, request_type, request_method, request_data
	}

	return -1, -1, -1, -1
}
