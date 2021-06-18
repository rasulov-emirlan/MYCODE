package servers

import (
	"net/http"
)

func Start(port string) error{
	s := NewServer()
	return http.ListenAndServe(port, s.router)
}
