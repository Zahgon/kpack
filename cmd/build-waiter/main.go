package main

import (
	"context"
	"flag"
)

var (
	mode     = flag.String("mode", "wait", "one of: wait or copy")
	to       = flag.String("to", "", "where to copy this binary")
	waitFile = flag.String("wait-file", "", "file to wait on")
	doneFile = flag.String("done-file", "", "file to write on completion")
	execute  = flag.String("execute", "", "What to run after waiting")
	errFile  = flag.String("error-file", "", "shared error file")
)

func main() {
	flag.Parse()

	switch *mode {
	case "wait":
		if *errFile != "" {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go watchForErrors(ctx, *errFile)
		}

		err := wait(*waitFile, *doneFile, *execute)
		if err != nil {
			exitWithError(err.Error(), *errFile)
		}

	case "copy":
		err := copy(*to)
		if err != nil {
			exitWithError(err.Error(), *errFile)
		}
	}
}

func wait(waitFile, doneFile, execute string) error { _ = "STUB: not implemented"; return nil }

func copy(to string) error { _ = "STUB: not implemented"; return nil }

func waitForStep(waitPath string) { _ = "STUB: not implemented"; return }

func exists(fileName string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func exitWithError(message, errorFile string) { _ = "STUB: not implemented"; return }

func watchForErrors(ctx context.Context, errFile string) { _ = "STUB: not implemented"; return }
