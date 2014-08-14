package main

import(
//	"fmt"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func handler_api1(r *http.Request, output *string) {
	// process input data
	r.ParseForm()
	
	if r.Method == "GET" {
		err_input_get(r, output)
	} else if r.Form["method"] != nil && r.Form["data"] != nil {
		switch r.Form["method"][0] {
		case "cli_init":
			deal_cli_init(r, output)
		case "contact":
			deal_contact(r, output)
		case "message":
			deal_message(r, output)
		default:
			err_api1(r, output)
		}
	} else {
		err_api1(r, output)
	}
	// process input data end
	
	/*
	c := session.DB("golang").C("test_a")
	back := `
		<html>
		<head>
		<title></title>
		</head>
	`
	fmt.Fprintln(w, back)
	
	fmt.Fprintln(w, r.Form)

	if r.Form["name"] != nil {
		if len(r.Form["name"][0]) > 0 {
			in := Sign{Name: r.Form["name"][0], Value: r.Form["value"][0]}
			c.Insert(in)
		}
		fmt.Fprintln(w, len(r.Form["name"][0]))
	}
	
	back = `
		<body>
		<form action="" method="post">
			name:<input type="text" name="name">
			value:<input type="text" name="value">
			<input type="submit" value="登记">
		</form>
		</body>
		</html>
	`
	fmt.Fprintln(w, back)

	var results []Sign

	c.Find(bson.M{}).Sort("_id").All(&results)

	for _, v := range results {
		back = back + "<p>" + v.Name + "</p>"
		back = back + "<p>" + v.Value + "</p>"
	}

	fmt.Fprintf(w, back)
	*/
}

func deal_cli_init(r *http.Request, output *string) {
	*output += r.Form["data"][0]
	*output += "\n"
	*output += "deal_cli_init"
}

func deal_contact(r *http.Request, output *string) {
	*output += r.Form["data"][0]
	*output += "\n"
	*output += "deal_contact"
}

func deal_message(r *http.Request, output *string) {
	*output += r.Form["data"][0]
	*output += "\n"
	*output += "deal_message"
}
