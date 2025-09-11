package rest

import (
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/busan"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
)

func GetBrowse(w http.ResponseWriter, r *http.Request) {

	// GET /browse

	var err error
	rdx, err = rdx.RefreshReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := compton.Page("emporium")
	p.SetAttribute("style", "--c-rep:var(--c-background)")

	pageStack := compton.FlexItems(p, direction.Column)
	p.Append(pageStack)

	shares := rdx.Keys(data.SharesFilesProperty)
	sortedShares := slices.Sorted(shares)

	for _, sha := range sortedShares {

		trimmedSha := strings.TrimSuffix(sha, "/")

		shaHeading := compton.Heading(1)
		var headingParts []string

		if shaParts := strings.Split(trimmedSha, "/"); len(shaParts) > 1 {

			shaPartsRow := compton.FlexItems(p, direction.Row).ColumnGap(size.Small)
			shaHeading.Append(shaPartsRow)

			for ii, part := range shaParts {

				shaPartsRow.Append(compton.Fspan(p, part).ForegroundColor(color.Foreground))
				if ii != len(shaParts)-1 {
					shaPartsRow.Append(compton.Fspan(p, "/").ForegroundColor(color.Gray))
				}

				headingParts = append(headingParts, part)
			}

		} else {
			shaHeading.Append(compton.Text(trimmedSha))
			headingParts = append(headingParts, trimmedSha)
		}

		headingId := busan.Sanitize(strings.Join(headingParts, " "))
		shaHeading.SetId(headingId)
		shaHeadingLink := compton.A("#" + headingId)
		shaHeadingLink.Append(compton.Fspan(p, "#").
			ForegroundColor(color.Gray).
			FontSize(size.XXLarge))

		lineStack := compton.FlexItems(p, direction.Row).AlignContent(align.Center)
		lineStack.Append(shaHeading, shaHeadingLink)

		pageStack.Append(lineStack)

		filesStack := compton.FlexItems(p, direction.Column)
		pageStack.Append(filesStack)

		if files, ok := rdx.GetAllValues(data.SharesFilesProperty, sha); ok {

			for _, filename := range files {

				fileLineItem := compton.Fspan(p, "").
					BackgroundColor(color.Highlight).
					BorderRadius(size.XSmall).
					FontSize(size.Large).
					Padding(size.Small)

				fileLink := compton.A("/file?dir=" + sha + "&base=" + filename)

				ext := filepath.Ext(filename)
				fse := strings.TrimSuffix(filename, ext)

				if strings.HasPrefix(fse, trimmedSha) {
					fse = strings.Replace(fse, trimmedSha+" - ", "", 1)
				}

				linkFilename := compton.Fspan(p, fse).Padding(size.Unset)
				linkExt := compton.Fspan(p, ext).ForegroundColor(color.Gray).Padding(size.Unset)

				fileLink.Append(linkFilename, linkExt)

				fileLineItem.Append(fileLink)

				filesStack.Append(fileLineItem)
			}
		}
	}

	if err = p.WriteResponse(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
