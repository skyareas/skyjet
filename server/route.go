package server

type Route interface {
	HandleRequest(*HttpRequest, *HttpResponse) error
}
