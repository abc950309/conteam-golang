package main

import(
	"fmt"
	"net/http"
	"github.com/abc950309/conteam-golang/data_err"
	"encoding/json"
)

func err_output(num int) string {
	err_obj, _ := json.Marshal(data_err.GetErrObj(num))
	output := string(err_obj)
	return output
}

func err_input_get(r *http.Request, output *string) {
	*output = err_output(100)
}

func err_api1(r *http.Request, output *string) {
	*output += fmt.Sprint(r)
}

