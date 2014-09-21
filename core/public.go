package core

import(
//	"fmt"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//	"github.com/abc950309/conteam-golang/data_struct"
)

const(
	ConstTypeErr		= -1
    ConstTypeContact	= 0
    ConstTypeMessage	= 1
)

const(
	ConstMethodErr		= -1
    ConstMethodInsert	= 0
    ConstMethodDelete	= 1
	ConstMethodUpdate	= 2
	ConstMethodGet		= 3
	ConstMethodList		= 4
)

func VerifyReqLogic(req_type int, req_method int) bool {
	if req_type < 0 || req_method < 0 {
		return false
	}
	
	verify_list := [2][5]bool{
	{	true,	true,	true,	true,	true	},
	{	true,	false,	false,	true,	true	},
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

func Controller() {
}


