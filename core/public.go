package core

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/abc950309/conteam-golang/data_struct"
	"github.com/nu7hatch/gouuid"
	"time"
)

const (
	ConstTypeErr     = -1
	ConstTypeContact = 0
	ConstTypeMessage = 1
)

const (
	ConstMethodErr    = -1
	ConstMethodInsert = 0
	ConstMethodDelete = 1
	ConstMethodUpdate = 2
	ConstMethodGet    = 3
	ConstMethodList   = 4
)

func VerifyReqLogic(req_type int, req_method int) bool {
	if req_type < 0 || req_method < 0 {
		return false
	}

	verify_list := [2][5]bool{
		{true, true, true, true, true},
		{true, false, false, true, true},
	}

	return verify_list[req_type][req_method]
}

func link_mongo_init() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

func Controller(user_token string, dealed_type int, dealed_method int, dealed_data interface{}, source string) interface{} {
	
	/**
	code, user_id = token_deal(user_token)

	if code < 0 {
		return nil
	}
	**/
	
	fmt.Println("Controller: " , dealed_data)
	
	switch dealed_type {
	case ConstTypeContact:
		switch dealed_method {
		case ConstMethodInsert:
			value, ok := dealed_data.(data_struct.Contact)
			if ok {
				return contact_insert(value)
			}

		case ConstMethodDelete:
			value, ok := dealed_data.(data_struct.ContactFilter)
			if ok {
				return contact_delete(value)
			}

		case ConstMethodUpdate:
			value, ok := dealed_data.(data_struct.Contact)
			if ok {
				return contact_update(value)
			}

		case ConstMethodGet:
			value, ok := dealed_data.(data_struct.ContactFilter)
			if ok {
				return contact_get(value)
			}

		case ConstMethodList:
			value, ok := dealed_data.(data_struct.ContactFilters)
			if ok {
				return contact_list(value)
			}

		}
	case ConstTypeMessage:
		switch dealed_method {
		case ConstMethodInsert:
			value, ok := dealed_data.(data_struct.Message)
			if ok {
				return message_insert(value)
			}
		case ConstMethodGet:
			value, ok := dealed_data.(data_struct.MessageFilter)
			if ok {
				return message_get(value)
			}

		case ConstMethodList:
			value, ok := dealed_data.(data_struct.MessageFilters)
			if ok {
				return message_list(value)
			}

		}
	default:
	}

	return nil
}

// type methods func(value) interface{}
// contact methods
func contact_insert(value data_struct.Contact) data_struct.Contact {
	return data_struct.Contact{}
}
func contact_delete(value data_struct.ContactFilter) data_struct.Contact {
	return data_struct.Contact{}
}
func contact_update(value data_struct.Contact) data_struct.Contact {
	return data_struct.Contact{}
}
func contact_get(value data_struct.ContactFilter) data_struct.Contact {
	return data_struct.Contact{}
}
func contact_list(value data_struct.ContactFilters) data_struct.ContactList {
	return data_struct.ContactList{}
}

// message methods
func message_insert(value data_struct.Message) data_struct.Message {
	if value.Content == "" {
		return data_struct.Message{}
	}
	
	u, err := uuid.NewV4()
	if err != nil {
		return data_struct.Message{}
	}
	value.MessageID = u.String()
	value.TimeStamp = time.Now().Unix()
	
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	session.SetMode(mgo.Monotonic, true)
	
    c := session.DB("message").C("data")
    err = c.Insert(value)
    if err != nil {
        panic(err)
    }
	
	return value
}
func message_get(value data_struct.MessageFilter) data_struct.Message {
	if value.MessageID == "" {
		return data_struct.Message{}
	}
		
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	session.SetMode(mgo.Monotonic, true)
	
    c := session.DB("message").C("data")
	result := data_struct.Message{}
	err = c.Find(bson.M{"message_id": value.MessageID}).One(&result)
	if err != nil {
		return data_struct.Message{}
	}
	
	return result
}
func message_list(value data_struct.MessageFilters) data_struct.MessageList {
	return data_struct.MessageList{}
}

func token_deal(user_token string, source string) (int, string) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	
	c := session.DB("token").C("valid")
	result := data_struct.Token{}
	err = c.Find(bson.M{"token": user_token}).One(&result)
	if err != nil {
		return -1, ""
	}

	if (time.Now().Unix() > result.Expire) || (source != result.Source) {
		return -1, ""
	}

	return -1, result.UserID
}
