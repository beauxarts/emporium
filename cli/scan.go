package cli

import (
	"net/url"
	"os"
	"path/filepath"
	"slices"

	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/camino"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

func ScanHandler(u *url.URL) error {
	return Scan()
}

func Scan() error {

	sa := nod.Begin("scanning shares...")
	defer sa.Done()

	metadataDir := camino.GetAbs(data.Metadata)

	rdx, err := redux.NewWriter(metadataDir, data.AllProperties()...)
	if err != nil {
		return err
	}

	shares := camino.GetAbs(data.Shares)

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
		return err
	}

	// clear redux before adding new values
	if err = rdx.CutKeys(data.SharesFilesProperty, slices.Collect(rdx.Keys(data.SharesFilesProperty))...); err != nil {
		return err
	}

	if err = rdx.BatchAddValues(data.SharesFilesProperty, dirFiles); err != nil {
		return err
	}

	return nil
}
