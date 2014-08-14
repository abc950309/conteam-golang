package data_err

import (
	"conteam/data_struct"
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
