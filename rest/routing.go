package rest

import (
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/nod"
	"net/http"
)

var (
	Auth = middleware.BasicHttpAuth
	Log  = nod.RequestLog
)

func HandleFuncs() {

	patternHandlers := map[string]http.Handler{
		"GET /browse": Log(http.HandlerFunc(GetBrowse)),

		"GET /file": Auth(Log(http.HandlerFunc(GetFile)), DefaultRole),

		"GET /": Log(http.RedirectHandler("/browse", http.StatusPermanentRedirect)),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
