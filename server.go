package https

import (
	"fmt"
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

func (s *Server) StartDev(addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		first, rest, isRoot := util.PopPath(r.URL.Path)
		if isRoot {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, "<pre>\n")
			for _, h := range s.AllowedHosts {
				fmt.Fprintf(w, "<a href=\"/%s\">%s\n</a>", h, h)
			}
			fmt.Fprintf(w, "</pre>\n")
			return
		}
		r.URL.Path = rest
		r.Host = first
		s.Handler.ServeHTTP(w, r)
	})
	return http.ListenAndServe(addr, nil)
}
