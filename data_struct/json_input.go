package data_struct

type JsonInput struct {
	CliInit JsonCliInit `json:"cli_init,omitempty"` //JsonCliInit
	Contact JsonContact `json:"contact,omitempty"` //JsonContact
	Message JsonMessage `json:"message,omitempty"` //JsonMessage
}

type JsonCliInit struct {
	JPushRegistrationID string `json:"jPushRegistrationID"`
	Platform string `json:"platform"`
}

type JsonContact struct {
	Act string `json:"act"`
	Read JsonContactRead `json:"read,omitempty"` //JsonContactRead
	Write JsonContactWrite `json:"write,omitempty"` //JsonContactWrite
}

type JsonContactRead struct {
	Filter JsonFilter `json:"filter,omitempty"` //JsonFilter
	List []string `json:"list,omitempty"`
}

type JsonContactWrite struct {
	UserID string `json:"user_id,omitempty"`
	Datas []ContactData `json:"datas,omitempty"`
}

type JsonMessage struct {
	Act string `json:"act"`
	Get JsonMessageGet `json:"get,omitempty"` //JsonMessageGet
	Send JsonMessageSend `json:"send,omitempty"` //JsonMessageSend
}

type JsonMessageGet JsonContactRead

type JsonMessageSend Message

type JsonFilter struct {
	Time string `json:"time,omitempty"`
	UserID []string `json:"user_id,omitempty"`
}
