package view_models

import (
	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/kevlar"
	"sort"
)

type BrowseViewModel struct {
	Shares      []string
	SharesFiles map[string][]string
}

func NewBrowseViewModel(rdx kevlar.ReadableRedux) *BrowseViewModel {
	shares := rdx.Keys(data.SharesFilesProperty)
	sort.Strings(shares)

	sharesFiles := make(map[string][]string, len(shares))
	for _, sha := range shares {
		if fs, ok := rdx.GetAllValues(data.SharesFilesProperty, sha); ok {
			sharesFiles[sha] = fs
		}
	}

	return &BrowseViewModel{
		Shares:      shares,
		SharesFiles: sharesFiles,
	}
}
