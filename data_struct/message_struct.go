package data_struct

type Message struct {
	MessageID string `json:"message_id,omitempty"`
    From PublicPersonID `json:"from,omitempty"`
    To PublicPersonID `json:"to,omitempty"`
	Content string `json:"content,omitempty"`
	TimeStamp int64 `json:"timestamp,omitempty"`
}
