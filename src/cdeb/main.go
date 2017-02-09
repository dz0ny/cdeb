package main

import (
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
)

var version = "dev"

var (
	verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	dumb     = kingpin.Command("dumb", "Create debian file from root and control folder.").Default()
	dRoot    = dumb.Arg("root", "Path to deb root folder").Required().ExistingDir()
	dControl = dumb.Arg("control", "Path to debian control folder").Required().ExistingDir()
	dDeb     = dumb.Arg("deb", "Path to output deb file").Required().String()

	gen   = kingpin.Command("gen", "Create debian file from root folder and .cdeb template.")
	gRoot = gen.Arg("root", "Path to deb root folder").Required().ExistingDir()
	gDeb  = gen.Arg("deb", "Path to output deb file").Required().String()
	cinit = kingpin.Command("init", "Create .cdeb template.")
)

func main() {

	kingpin.Version(version).Author("Janez Troha")
	kingpin.CommandLine.Help = "Build packages for debian with great ease!"
	switch kingpin.Parse() {
	case "dumb":
		err := createDumbDeb(*dDeb, *dRoot, *dControl)
		if err != nil {
			log.Fatalln(err)
		}
	case "gen":
		fmt.Println("Not ready yet")

	case "init":
		fmt.Println("Not ready yet")
	}
}
