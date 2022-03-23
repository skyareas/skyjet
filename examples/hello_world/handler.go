package main

import "github.com/skyareas/skyjet"

func requestHandler(req *skyjet.HttpRequest, res *skyjet.HttpResponse) error {
	name, nameParamExists := req.LookupParam("name")
	if !nameParamExists {
		name = req.Query().Get("name")
	}
	if name != "" {
		return res.Send([]byte("Hello, " + name + "!"))
	}
	return res.Send([]byte("Please tell me your name!"))
}
