package server

func (s *Server) Router() {
	NewHealthCheck(s.fw)
	NewUserController(s.fw)
}
