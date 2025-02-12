package cli

import (
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/backups"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"net/url"
)

const daysToPreserveFiles = 30

func BackupHandler(_ *url.URL) error {
	return Backup()
}

func Backup() error {
	ea := nod.NewProgress("backing up metadata...")
	defer ea.Done()

	amp, err := pathways.GetAbsDir(paths.Metadata)
	if err != nil {
		return err
	}

	abp, err := pathways.GetAbsDir(paths.Backups)
	if err != nil {
		return err
	}

	if err := backups.Compress(amp, abp); err != nil {
		return err
	}

	cba := nod.NewProgress("cleaning up old backups...")
	defer cba.Done()

	if err := backups.Cleanup(abp, true, cba); err != nil {
		return err
	}

	return nil
}
