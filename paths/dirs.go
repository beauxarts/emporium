package paths

import (
	"github.com/boggydigital/pasu"
)

const DefaultEmporiumRootDir = "/usr/share/emporium"

const (
	Backups  pasu.AbsDir = "backups"
	Share    pasu.AbsDir = "share"
	Metadata pasu.AbsDir = "metadata"
)

var AllAbsDirs = []pasu.AbsDir{
	Backups,
	Share,
	Metadata,
}
