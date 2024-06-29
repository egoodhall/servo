package options

type Options struct {
	Enabled bool   `json:"enabled" default:"false" desc:"If false, the gostruct plugin will not generate code"`
	Package string `json:"package" desc:"The name of the package to use for the generated go file"`
	File    string `json:"file" default:"gohttp.gen.go" desc:"The file to generate to generate code in"`
	Server  bool   `json:"server.enabled" default:"true" desc:"Whether to generate HTTP server code"`
	Client  bool   `json:"client.enabled" default:"true" desc:"Whether to generate HTTP client code"`
}
