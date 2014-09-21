package data_struct

type Message struct {
	MessageID string `json:"message_id,omitempty"`
    From PublicPersonID `json:"from,omitempty"`
    To PublicPersonID `json:"to,omitempty"`
	Content string `json:"content,omitempty"`
	TimeStamp int64 `json:"timestamp,omitempty"`
}

type PublicPersonID struct {
    UserID string `json:"user_id,omitempty"`
	GroupID string `json:"group_id,omitempty"`
}

type MessageList struct {
	List []ContactData `json:"list"`
}

type MessageListItem struct {
	MessageID string `json:"message_id"`
	TimeStamp int64 `json:"timestamp"`
}
