package data_err

type ErrOutput struct {
	Code int `json:"code"`
    ErrorDesc ErrDesc `json:"error_description"`
}

type ErrDesc struct {
	Description string `json:"description"`
	Url string `json:"url"`
}
