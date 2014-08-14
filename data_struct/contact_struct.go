package data_struct

type Contact struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ShowMode bool `json:"show_mode"`
	Datas []ContactData `json:"datas,omitempty"`
}
