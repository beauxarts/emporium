package cli

import (
	"github.com/beauxarts/emporium/data"
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/url"
	"os"
	"path/filepath"
)

func ScanHandler(u *url.URL) error {
	return Scan()
}

func Scan() error {

	sa := nod.Begin("scanning shares...")
	defer sa.End()

	metadataDir, err := pathways.GetAbsDir(paths.Metadata)
	if err != nil {
		return sa.EndWithError(err)
	}

	rdx, err := kevlar.NewReduxWriter(metadataDir, data.AllProperties()...)
	if err != nil {
		return sa.EndWithError(err)
	}

	shares, err := pathways.GetAbsDir(paths.Shares)
	if err != nil {
		return sa.EndWithError(err)
	}

	dirFiles := make(map[string][]string)

	if err = filepath.Walk(shares,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if rp, err := filepath.Rel(shares, path); err == nil {

				if rp != "." && filepath.Ext(rp) != "" {
					dir, base := filepath.Split(rp)
					dirFiles[dir] = append(dirFiles[dir], base)
				}
			} else {
				return err
			}
			return nil
		}); err != nil {
		return sa.EndWithError(err)
	}

	// clear redux before adding new values
	if err := rdx.CutKeys(data.SharesFilesProperty, rdx.Keys(data.SharesFilesProperty)...); err != nil {
		return sa.EndWithError(err)
	}

	if err := rdx.BatchAddValues(data.SharesFilesProperty, dirFiles); err != nil {
		return sa.EndWithError(err)
	}

	sa.EndWithResult("done")

	return nil
}
