package server

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) ListenAndServe() error {
	return nil
}

func (s *HttpServer) Shutdown() error {
	return nil
}
