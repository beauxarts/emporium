package data

import (
	"github.com/boggydigital/pathways"
)

const DefaultEmporiumRootDir = "/usr/share/emporium"

const (
	Backups  pathways.AbsDir = "backups"
	Shares   pathways.AbsDir = "shares"
	Metadata pathways.AbsDir = "metadata"
)

var AllAbsDirs = []pathways.AbsDir{
	Backups,
	Shares,
	Metadata,
}
