package main

import (
	"al.essio.dev/a/tools/internal/version"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const (
	PortExe = "/opt/local/bin/port"
)

var (
	helpMode    bool
	versionMode bool
	cwd         string
)

func init() {
	flag.BoolVar(&helpMode, "help", false, "display this help and exit.")
	flag.BoolVar(&versionMode, "version", false, "output version information and exit.")
}

func main() {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("portup: ")
	log.SetOutput(os.Stdout) // write to stdout as it could be piped
	flag.Parse()

	handleHelpAndVersionModes()

	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cwd = d

	if len(flag.Args()) > 0 {
		logfile, err := openLogFile(flag.Arg(0))
		if err != nil {
			log.Fatalf("couldn't open the file %s: %v", flag.Arg(0), err)
		}

		log.SetOutput(logfile)
	}

	if err := runPortCommand("-v", "selfupdate"); err != nil {
		log.Fatal(err)
	}

	if err := runPortCommand("list", "outdated"); err != nil {
		log.Fatal(err)
	}

	if err := runPortCommand("-v", "-u", "-c", "upgrade", "outdated"); err != nil {
		log.Fatal(err)
	}
}

func runPortCommand(args ...string) error {
	log.Println("Will run", PortExe, args)
	cmd := exec.Command(PortExe, args...)
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func openLogFile(filename string) (io.WriteCloser, error) {
	fp, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	return fp, nil
}

func usage() {
	s := `Usage: %s portup [PATH]
Make the management of the PATH environment variable
simple, fast, and predictable.

Commands:

   append, a       append a path to the end
   drop, d         drop a path
   list, l         list the paths
   prepend, p      prependPath a path to the list

Options:
`
	_, _ = fmt.Fprintln(os.Stderr, s)

	flag.PrintDefaults()
}

func handleHelpAndVersionModes() {
	if !helpMode && !versionMode {
		return
	}

	switch {
	case helpMode:
		usage()
	case versionMode:
		fmt.Println(version.Version())
	}

	os.Exit(0)
}
