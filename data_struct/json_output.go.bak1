package data_struct

type JsonOutput struct {
	Code int `json:"code,omitempty"`
    ErrorDesc interface{} `json:"error_description,omitempty"` //JsonErrDesc
	Response interface{} `json:"response,omitempty"` //JsonResponse
}

type JsonErrDesc struct {
	Description string `json:"description"`
	Url string `json:"url"`
}

type JsonResponse struct {
	Method string `json:"method"`
	CliInitBack interface{} `json:"cli_init_back,omitempty"` //JsonCliInitBack
	ContactList []ContactData `json:"contact_list,omitempty"`
	MessageList []Message `json:"message_list,omitempty"`
}

type JsonCliInitBack struct{
	jPushRegistrationID string `json:"jPushRegistrationID"`
	CliID string `json:"cli_id"`
}
