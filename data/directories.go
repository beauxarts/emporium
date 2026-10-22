package data

import (
	"os"

	"github.com/boggydigital/camino"
)

const emporiumRootDir = "/usr/share/emporium"

const (
	directoriesFilename = "directories.txt"
)

const (
	Backups camino.AbsDir = iota
	Shares
	Metadata
)

var absDirNames = map[camino.AbsDir]string{
	Backups:  "backups",
	Shares:   "shares",
	Metadata: "metadata",
}

func InitEmporiumCamino() error {

	var overrides map[string]string

	if _, err := os.Stat(directoriesFilename); err == nil {
		if overrides, err = camino.ReadOverrides(directoriesFilename); err != nil {
			return err
		}
	}

	emporiumAbsPaths := camino.ResolveAbsPaths(emporiumRootDir, absDirNames, overrides)

	return camino.Register(emporiumAbsPaths, nil, nil, len(overrides) == 0)

}
