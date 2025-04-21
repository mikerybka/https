package https

import (
	"net/http"

	"github.com/mikerybka/util"
)

type Server struct {
	Email        string
	CertDir      string
	AllowedHosts []string
	Handler      http.Handler
}

func (s *Server) Start() error {
	hostAllowlist := map[string]bool{}
	for _, h := range s.AllowedHosts {
		hostAllowlist[h] = true
	}
	return util.ServeHTTPS(s.Handler, s.Email, s.CertDir, func(host string) bool {
		return hostAllowlist[host]
	})
}
