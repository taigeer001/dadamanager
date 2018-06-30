package xf

import (
	"os"
	"fmt"
)

type _Variable struct{
	Modules string
	Data string
}

var Dir *_Variable

func NewVariable() *_Variable {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &_Variable{
		Modules: pwd+"/modules",
		Data: pwd+"/data",
	}
}

func GetModelesDir(name string) string {
	if Dir == nil {
		Dir = NewVariable()
	}
	return Dir.Modules + "/" + name
}

func GetDataDir(name string) string {
	if Dir == nil {
		Dir = NewVariable()
	}
	return Dir.Data + "/" + name
}

