package main

import (
	"flag"
	"fmt"
	"gogist/gogist"
	"os"
	"strings"
)

const (
	BANNER = `
                   _     _
                  (_)   | |
  __ _  ___   __ _ _ ___| |_
 / _  |/ _ \ / _  | / __| __|
| (_| | (_) | (_| | \__ \ |_
 \__, |\___/ \__, |_|___/\__|
  __/ |       __/ |
 |___/       |___/

 Command Line Tool to upload contents to Gist (https://gist.github.com/)
 Version: %s

`
	VERSION = "v0.0.1"
)

var (
	files       string
	help        bool
	login       bool
	filename    string
	kind        string
	public      bool
	description string
)

func init() {

	flag.BoolVar(&help, "help", false, "Help menu")
	flag.BoolVar(&login, "l", false, "Authenticate to GitHub")
	flag.BoolVar(&public, "p", false, "Sets the gist private")
	flag.StringVar(&description, "d", "", "Adds a `description` to the gist")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() >= 1 {
		files = strings.Join(flag.Args(), " ")
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {

	if login {
		gogist.CreateNewToken()
	}

	gogist.CreateGist(flag.Args(), filename, public, description)
}
