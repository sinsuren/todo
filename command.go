package main

import "flag"

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func newCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:

	}
}
