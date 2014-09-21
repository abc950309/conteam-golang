package data_struct

type Contact struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	ShowMode bool `json:"show_mode,omitempty"`
	Datas []ContactData `json:"datas,omitempty"`
	TimeStamp int64 `json:"timestamp"`
}

type ContactData struct {
	DataID string `json:"data_id,omitempty"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContactList struct {
	List []ContactData `json:"list"`
}

type ContactListItem struct {
	UserID string `json:"user_id"`
	TimeStamp int64 `json:"timestamp"`
}
