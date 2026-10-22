package rest

import (
	"crypto/sha256"

	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/camino"
	"github.com/boggydigital/middleware"
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

	metadataDir := camino.GetAbs(data.Metadata)

	var err error
	if rdx, err = redux.NewReader(metadataDir, data.AllProperties()...); err != nil {
		return err
	}

	return nil
}
