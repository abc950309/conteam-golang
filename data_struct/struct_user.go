package data_struct

type User struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	LoginName []string `json:"login_name"`
	Password string `json:"password"`
	Admin []string `json:"admin,omitempty"`
	TimeStamp int64 `json:"timestamp"`
}
