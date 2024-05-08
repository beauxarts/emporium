package rest

import (
	"github.com/boggydigital/nod"
	"net/http"
)

var (
	Log = nod.RequestLog
)

func HandleFuncs() {

	patternHandlers := map[string]http.Handler{
		"GET /browse": Log(http.HandlerFunc(GetBrowse)),
		"GET /file":   Log(http.HandlerFunc(GetFile)),

		"GET /": Log(http.RedirectHandler("/", http.StatusPermanentRedirect)),
	}

	for p, h := range patternHandlers {
		http.HandleFunc(p, h.ServeHTTP)
	}
}
