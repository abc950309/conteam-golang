package data_err

type JsonOutput struct {
	Code int `json:"code,omitempty"`
    ErrorDesc interface{} `json:"error_description,omitempty"` //JsonErrDesc
}

type JsonErrDesc struct {
	Description string `json:"description"`
	Url string `json:"url"`
}

func GetErrObj(num int) *JsonOutput {
	output := JsonOutput{
		Code: num,
		ErrorDesc: JsonErrDesc{
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