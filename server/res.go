package server

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	w http.ResponseWriter
}

func (r *HttpResponse) Send(data interface{}) error {
	return json.NewEncoder(r.w).Encode(data)
}
