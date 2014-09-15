package main

import(
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
//	"github.com/abc950309/conteam-golang/data_struct"
)

var Session *mgo.Session

func init() {
	Session = link_mongo_init()
}

func main() {
	debug()
	
	http.HandleFunc("/", handler)
	
	fmt.Println("ListenAndServe: Ready!")
	
	err := http.ListenAndServe(":9876", nil)
	if err != nil {
        fmt.Println("ListenAndServe: ", err)
    }
	
	Session.Close()
}

func handler(w http.ResponseWriter, r *http.Request) {
	output := ""
	
	switch r.URL.Path[1:] {
	case "api1":
		handler_api1(r, &output)
	default:
		handler_root(r, &output)
	}
	
	fmt.Fprintf(w, output)
}

func debug() {
	
}
