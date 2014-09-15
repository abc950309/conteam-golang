package data_err

import (
	"github.com/abc950309/conteam-golang/data_struct"
)

func GetErrObj(num int) *data_struct.JsonOutput {
	output := data_struct.JsonOutput{
		Code: num,
		ErrorDesc: data_struct.JsonErrDesc{
			Description: ErrListDesc[num],
			Url: ErrListUrl[num],
		},
	}
	
	/*
	output.Code = num
	output.ErrorDesc.Description = ErrListDesc[num]
	output.ErrorDesc.Url = ErrListUrl[num]
	*/
	
	return &output
}
