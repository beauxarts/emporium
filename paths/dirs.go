package paths

import (
	"github.com/boggydigital/pasu"
)

const DefaultEmporiumRootDir = "/usr/share/emporium"

const (
	Backups  pasu.AbsDir = "backups"
	Shares   pasu.AbsDir = "shares"
	Metadata pasu.AbsDir = "metadata"
)

var AllAbsDirs = []pasu.AbsDir{
	Backups,
	Shares,
	Metadata,
}
