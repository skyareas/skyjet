package skyjet

type RouteHandler = func(req *HttpRequest, res *HttpResponse) error
