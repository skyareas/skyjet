package skyjet

import "path"

func Static() *Router {
	r := NewRouter()
	r.Get("*", func(req *HttpRequest, res *HttpResponse) error {
		return res.SendFile(path.Join(app.cfg.Http.ContentRoot, req.Request.URL.Path))
	})
	return r
}
