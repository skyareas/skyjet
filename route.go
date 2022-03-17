package skyjet

type Route interface {
	HandleRequest(*HttpRequest, *HttpResponse) error
}
