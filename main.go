package main

import (
	"bytes"
	_ "embed"
	"github.com/beauxarts/emporium/cli"
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pasu"
	"os"
)

var (
	//go:embed "cli-commands.txt"
	cliCommands []byte
	//go:embed "cli-help.txt"
	cliHelp []byte
)

const (
	dirOverridesFilename = "directories.txt"
)

func main() {
	nod.EnableStdOutPresenter()

	ea := nod.Begin("emporium is serving your sharing needs")
	defer ea.End()

	if err := pasu.Setup(dirOverridesFilename,
		paths.DefaultEmporiumRootDir,
		nil,
		paths.AllAbsDirs...); err != nil {
		_ = ea.EndWithError(err)
		os.Exit(1)
	}

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		nil)

	if err != nil {
		_ = ea.EndWithError(err)
		os.Exit(1)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":  cli.BackupHandler,
		"serve":   cli.ServeHandler,
		"version": cli.VersionHandler,
	})

	if err := defs.AssertCommandsHaveHandlers(); err != nil {
		_ = ea.EndWithError(err)
		os.Exit(1)
	}

	if err := defs.Serve(os.Args[1:]); err != nil {
		_ = ea.EndWithError(err)
		os.Exit(1)
	}

}
