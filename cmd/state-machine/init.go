package main

import "os"

type Init struct {
	Args struct {
		Type   string `required:"yes"`
		States []string
	} `positional-args:"yes"`
}

type initArgs struct {
	Name    string
	Type    string
	State   string
	Package string
}

func (init *Init) Execute([]string) error {
	file := "init"
	if config.Name != "" {
		file = config.Name + "_" + file
	}
	file += ".go"
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = initTpl.Execute(f, initArgs{
		Name:    config.Name,
		Package: config.Package,
		Type:    init.Args.Type,
	})
	if err != nil {
		return err
	}
	for i, st := range init.Args.States {
		add := Add{
			Args: AddArgs{Type: init.Args.Type, State: st, Num: i + 1},
		}
		err := add.Execute(nil)
		if err != nil {
			return err
		}
	}
	return nil
}
