package main

import (
	"bytes"
	_ "embed"
	"github.com/beauxarts/emporium/cli"
	"github.com/beauxarts/emporium/paths"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"log"
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
	defer ea.Done()

	if err := pathways.Setup(dirOverridesFilename,
		paths.DefaultEmporiumRootDir,
		nil,
		paths.AllAbsDirs...); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	defs, err := clo.Load(
		bytes.NewBuffer(cliCommands),
		bytes.NewBuffer(cliHelp),
		nil)

	if err != nil {
		log.Println(err.Error())
		os.Exit(2)
	}

	clo.HandleFuncs(map[string]clo.Handler{
		"backup":  cli.BackupHandler,
		"migrate": cli.MigrateHandler,
		"scan":    cli.ScanHandler,
		"serve":   cli.ServeHandler,
		"version": cli.VersionHandler,
	})

	if err := defs.AssertCommandsHaveHandlers(); err != nil {
		log.Println(err.Error())
		os.Exit(3)
	}

	if err := defs.Serve(os.Args[1:]); err != nil {
		log.Println(err.Error())
		os.Exit(4)
	}

}
