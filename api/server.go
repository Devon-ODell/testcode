package api

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserbyID)
	return http.ListenAndServe(s.listenAssr, nil)
}

func (s *Server) handleGetUserbyID(w http.ResposeWriter, r *http.Request) {} 

