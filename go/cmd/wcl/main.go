package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"bitbucket.org/jsxqf/tool-kit/go/pkg"
)

var (
	mkansiblePath  string
	randompassSize int
)

const (
	mkansible  = "mkansible"
	randompass = "randompass"
)

func main() {

	mkansibleCmd := flag.NewFlagSet(mkansible, flag.ExitOnError)
	mkansibleCmd.StringVar(&mkansiblePath, "path", ".", "where to generate ansible scaffold")

	randompassCmd := flag.NewFlagSet(randompass, flag.ExitOnError)
	randompassCmd.IntVar(&randompassSize, "n", 10, "how many letters you want")

	if len(os.Args) == 1 {
		// not subcommand
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case mkansible:
		if err := mkansibleCmd.Parse(os.Args[2:]); err != nil {
			log.Fatalf("E! parse flag set failed: %v", err)
		}
		if err := pkg.Mkansible(mkansiblePath); err != nil {
			log.Fatalf("E! mkansible failed: %v", err)

		}

	case randompass:
		if err := randompassCmd.Parse(os.Args[2:]); err != nil {
			log.Fatalf("E! parse flag set failed: %v", err)
		}
		// TODO
	default:
		flag.PrintDefaults()
	}

}

func usage() {

	msg := `
Usage : wcl [command] [args]

Use wcl [command] -h to get more details
`
	fmt.Printf("%s", msg)
}
