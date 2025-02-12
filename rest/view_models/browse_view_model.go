package view_models

import (
	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/redux"
	"slices"
)

type BrowseViewModel struct {
	Shares      []string
	SharesFiles map[string][]string
}

func NewBrowseViewModel(rdx redux.Readable) *BrowseViewModel {
	shares := rdx.Keys(data.SharesFilesProperty)
	sortedShares := slices.Sorted(shares)

	sharesFiles := make(map[string][]string, len(sortedShares))
	for _, sha := range sortedShares {
		if fs, ok := rdx.GetAllValues(data.SharesFilesProperty, sha); ok {
			sharesFiles[sha] = fs
		}
	}

	return &BrowseViewModel{
		Shares:      sortedShares,
		SharesFiles: sharesFiles,
	}
}
