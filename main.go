package main

import (
	"flag"
	"fmt"
	//	"io"
	"gogist/gogist"
	"os"
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

func init() {

	flag.Bool("l", false, "Authenticate to GitHub")
	flag.String("f", "", "Sets the `filename`")
	flag.String("t", "", "Sets the `type`")
	flag.Bool("p", false, "Sets the gist private")
	flag.String("d", "", "Adds a `description` to the gist")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	var username string = flag.Args()[0]
	var password string = flag.Args()[1]

	fmt.Println(username, password)

	gogist.CreateNewToken(username, password)
}
