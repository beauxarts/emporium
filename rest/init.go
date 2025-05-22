package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/middleware"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

const DefaultRole = "default"

var (
	rdx redux.Readable
)

func SetUsername(role, u string) {
	middleware.SetUsername(role, sha256.Sum256([]byte(u)))
}

func SetPassword(role, p string) {
	middleware.SetPassword(role, sha256.Sum256([]byte(p)))
}

func Init() error {

	metadataDir, err := pathways.GetAbsDir(data.Metadata)
	if err != nil {
		return err
	}

	if rdx, err = redux.NewReader(metadataDir, data.AllProperties()...); err != nil {
		return err
	}

	return nil
}
