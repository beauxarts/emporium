package rest

import (
	"crypto/sha256"
	"embed"
	"github.com/beauxarts/emporium/data"
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/pathways"
	"html/template"
	"path/filepath"
	"strings"
)

const DefaultRole = "default"

var (
	rdx  kevlar.ReadableRedux
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

	metadataDir, err := pathways.GetAbsDir(paths.Metadata)
	if err != nil {
		return err
	}

	if rdx, err = kevlar.NewReduxReader(metadataDir, data.AllProperties()...); err != nil {
		return err
	}

	tmpl = template.Must(
		template.
			New("").
			Funcs(FuncMap()).
			ParseFS(templates, "templates/*.gohtml"))

	return nil
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"formatFilename": formatFilename,
		"formatShare":    formatShare,
	}
}

func formatFilename(name string) template.HTML {
	ext := filepath.Ext(name)
	fnse := strings.TrimSuffix(name, ext)
	return template.HTML("<span>" + fnse + "</span><span class='subtle'>" + ext + "</span>")
}

func formatShare(name string) template.HTML {
	name = strings.TrimSuffix(name, "/")
	parts := strings.Split(name, "/")
	return template.HTML(strings.Join(parts, "<span class='subtle'> / </span>"))
}
