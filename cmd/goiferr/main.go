package main

import (
	"go/printer"
	"log"
	"os"

	"golang.org/x/tools/go/loader"

	"github.com/motemen/goiferr"
)

func main() {
	conf := loader.Config{
		AllowErrors: true,
	}
	conf.FromArgs(os.Args[1:], true)
	prog, err := conf.Load()

	for _, pkg := range prog.InitialPackages() {
		log.Println(pkg)
		for _, f := range pkg.Files {
			log.Println(f)
			iferr.RewriteFile(conf.Fset, f, pkg.Info)
			printer.Fprint(os.Stdout, conf.Fset, f)
		}
	}

}
