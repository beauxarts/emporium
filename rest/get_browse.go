package rest

import (
	"github.com/beauxarts/emporium/rest/view_models"
	"net/http"
)

func GetBrowse(w http.ResponseWriter, r *http.Request) {

	// GET /browse

	var err error
	rdx, err = rdx.RefreshReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	bvm := view_models.NewBrowseViewModel(rdx)

	if err := tmpl.ExecuteTemplate(w, "files", bvm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
