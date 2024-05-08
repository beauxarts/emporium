package rest

import (
	"crypto/sha256"
	"embed"
	"github.com/beauxarts/emporium/data"
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/pasu"
	"html/template"
)

const DefaultRole = "default"

var (
	rdx  kvas.ReadableRedux
	tmpl *template.Template
	//go:embed "templates/*.gohtml"
	templates embed.FS
)

func SetUsername(role, u string) {
	middleware.SetUsername(role, sha256.Sum256([]byte(u)))
}

func SetPassword(role, p string) {
	middleware.SetPassword(role, sha256.Sum256([]byte(p)))
}

func Init() error {

	metadataDir, err := pasu.GetAbsDir(paths.Metadata)
	if err != nil {
		return err
	}

	if rdx, err = kvas.NewReduxReader(metadataDir, data.AllProperties()...); err != nil {
		return err
	}

	tmpl = template.Must(
		template.
			New("").
			ParseFS(templates, "templates/*.gohtml"))

	return nil
}
