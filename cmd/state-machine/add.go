package main

import (
	"os"
	"strings"
)

type Add struct {
	Args AddArgs `positional-args:"yes"`
}

type AddArgs struct {
	Type  string `required:"yes"`
	State string `required:"yes"`
	Num   int    `required:"yes"`
}

type addArgs struct {
	Name    string
	Type    string
	Package string
	State   string
	Num     int
}

func (init *Add) Execute([]string) error {
	file := strings.ToLower(init.Args.State)
	if config.Name != "" {
		file = config.Name + "_" + file
	}
	file += ".go"
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return stateTpl.Execute(f, addArgs{
		Name:    config.Name,
		Package: config.Package,
		Type:    init.Args.Type,
		Num:     init.Args.Num,
		State:   init.Args.State,
	})
}
