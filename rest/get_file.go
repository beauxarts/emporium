package rest

import (
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/http"
	"os"
	"path/filepath"
)

func GetFile(w http.ResponseWriter, r *http.Request) {

	// GET /file?dir&base

	dir := r.URL.Query().Get("dir")
	base := r.URL.Query().Get("base")

	sharesDir, err := pathways.GetAbsDir(paths.Shares)
	if err != nil {
		http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
	}

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
