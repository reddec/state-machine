package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"text/template"
)

var config struct {
	Package string `short:"p" description:"package" default:"main"`
	Name    string `short:"n" description:"namespace"`
	Init    Init   `description:"init machine" command:"init"`
	Add     Add    `description:"add state" command:"add"`
}

//go:generate go-bindata -pkg main templates/
var (
	initTpl  = template.Must(template.New("").Parse(string(MustAsset("templates/init.gotemplate"))))
	stateTpl = template.Must(template.New("").Parse(string(MustAsset("templates/state.gotemplate"))))
)

func main() {
	_, err := flags.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}
}
