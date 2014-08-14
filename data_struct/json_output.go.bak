package data_struct

type JsonOutput struct {
	Code int `json:"code,omitempty"`
    ErrorDesc interface{} `json:"error_description,omitempty"` //JsonErrDesc
	CliInit interface{} `json:"cli_init,omitempty"` //JsonCliInit
	Contact interface{} `json:"contact,omitempty"` //JsonContact
	Message interface{} `json:"message,omitempty"` //JsonMessage
	Response interface{} `json:"response,omitempty"` //JsonResponse
}

type JsonErrDesc struct {
	Description string `json:"description"`
	Url string `json:"url"`
}

type JsonCliInit struct {
	JPushRegistrationID string `json:"jPushRegistrationID"`
	Platform string `json:"platform"`
}

type JsonContact struct {
	Act string `json:"act"`
	Read interface{} `json:"read,omitempty"` //JsonContactRead
	Write interface{} `json:"write,omitempty"` //JsonContactWrite
}

type JsonContactRead struct {
	Filter interface{} `json:"filter,omitempty"` //JsonFilter
	List []string `json:"list,omitempty"`
}

type JsonContactWrite struct {
	UserID string `json:"user_id,omitempty"` //JsonFilter
	Datas []ContactData `json:"datas,omitempty"`
}

type JsonMessage struct {
	Act string `json:"act"`
	Get interface{} `json:"get,omitempty"` //JsonMessageGet
	Send interface{} `json:"send,omitempty"` //JsonMessageSend
}

type JsonMessageGet JsonContactRead

type JsonMessageSend Message

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

type JsonFilter struct {
	Time string `json:"time,omitempty"`
	UserID []string `json:"user_id,omitempty"`
}
