package main

import (
	//	"fmt"
	"net/http"
)

func handler_root(r *http.Request, output *string) {
	*output += `
		Wrong Way!!!
	`
}
