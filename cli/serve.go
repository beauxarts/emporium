package cli

import (
	"fmt"
	"github.com/beauxarts/emporium/rest"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
	"strconv"
)

func ServeHandler(u *url.URL) error {
	portStr := u.Query().Get("port")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	username := u.Query().Get("username")
	password := u.Query().Get("password")

	rest.SetUsername(rest.DefaultRole, username)
	rest.SetPassword(rest.DefaultRole, password)

	stderr := u.Query().Has("stderr")

	return Serve(port, stderr)
}

func Serve(port int, stderr bool) error {

	if stderr {
		nod.EnableStdErrLogger()
		nod.DisableOutput(nod.StdOut)
	}

	rest.HandleFuncs()

	if err := rest.Init(); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
