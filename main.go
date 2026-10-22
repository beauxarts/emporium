package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"

	"github.com/beauxarts/emporium/cli"
	"github.com/beauxarts/emporium/data"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/nod"
)

var (
	//go:embed "cli-commands.txt"
	cliCommands []byte
	//go:embed "cli-help.txt"
	cliHelp []byte
)

func main() {
	nod.EnableStdOutPresenter()

	ea := nod.Begin("emporium is serving your sharing needs")
	defer ea.Done()

	if err := data.InitEmporiumCamino(); err != nil {
		log.Fatalln(nod.Error(err))
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
		"scan":    cli.ScanHandler,
		"serve":   cli.ServeHandler,
		"version": cli.VersionHandler,
	})

	if err = defs.AssertCommandsHaveHandlers(); err != nil {
		log.Println(err.Error())
		os.Exit(3)
	}

	if err = defs.ParseServe(os.Args[1:]); err != nil {
		log.Println(err.Error())
		os.Exit(4)
	}

}
