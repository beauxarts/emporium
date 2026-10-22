package rest

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/camino"
	"github.com/boggydigital/nod"
)

func GetFile(w http.ResponseWriter, r *http.Request) {

	// GET /file?dir&base

	dir := r.URL.Query().Get("dir")
	base := r.URL.Query().Get("base")

	// make sure we're working with file and not path
	base = filepath.Base(base)

	sharesDir := camino.GetAbs(data.Shares)

	absLocalFilePath := filepath.Join(sharesDir, dir, base)
	if _, err := os.Stat(absLocalFilePath); err == nil {
		_, filename := filepath.Split(absLocalFilePath)
		w.Header().Set("Cache-Control", "max-age=31536000")
		w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
		http.ServeFile(w, r, absLocalFilePath)
	} else {
		http.Error(w, nod.Error(err).Error(), http.StatusNotFound)
	}
}
