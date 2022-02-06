package server

import "io"

// HttpRequestBody struct represents an Http request body.
type HttpRequestBody struct {
	raw []byte
	// this allows users to parse the body to any value and use it across routes
	Value interface{}
}

// Read reads all value of r into HttpRequestBody.
func (ref *HttpRequestBody) Read(r io.ReadCloser) (err error) {
	ref.raw, err = io.ReadAll(r)
	return
}

// Bytes return the Http request body as []byte
func (ref *HttpRequestBody) Bytes() []byte {
	return ref.raw
}

// Bytes return the Http request body as string
func (ref *HttpRequestBody) String() string {
	return string(ref.raw)
}
