package main

import (
	"github.com/egoodhall/servo/cliutil"
	"github.com/egoodhall/servo/plugin"
)

type Options struct {
	Enabled      bool              `name:"enabled" default:"false" desc:"If false, the gostruct plugin will not generate code"`
	Package      string            `name:"package" desc:"The name of the package to use for the generated go file"`
	File         string            `name:"file" default:"gotype.gen.go" desc:"The file to generate to generate code in"`
	Tags         []string          `name:"tags" default:"json" desc:"The tags to add to the generated struct"`
	OptionalTags map[string]string `name:"optional_tags" default:"json=omitempty" desc:"The tag value to add to the generated struct, if the field is optional"`
}

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest(new(GoTypePlugin))
	})
}

type GoTypePlugin struct{}
