package options

type Options struct {
	Enabled bool   `name:"enabled" default:"false" desc:"If false, the gostruct plugin will not generate code"`
	Package string `name:"package" desc:"The name of the package to use for the generated go file"`
	File    string `name:"file" default:"gohttp.gen.go" desc:"The file to generate to generate code in"`
	Server  bool   `name:"server.enabled" default:"true" desc:"Whether to generate HTTP server code"`
	Client  bool   `name:"client.enabled" default:"true" desc:"Whether to generate HTTP client code"`
}
