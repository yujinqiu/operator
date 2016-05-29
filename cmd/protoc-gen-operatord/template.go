package main

import "github.com/sr/operator/generator"

var servicesTemplate = generator.NewTemplate("services-gen.go",
	`// Code generated by protoc-gen-operatord
package main

import (
	"errors"
	"flag"
	"os"

	"github.com/sr/operator"
	"google.golang.org/grpc"

{{range .Services}}
	"{{.ImportPath}}"
{{end}}
)

func registerServices(
	server *grpc.Server,
	logger operator.Logger,
	flags *flag.FlagSet,
) error {
{{range .Services}}
	{{.Name}}Config := &{{.PackageName}}.{{.FullName}}Config{}
{{- end}}
{{range .Services}}
	{{- $serviceName := .Name }}
	{{- range .Config}}
	flags.StringVar(&{{$serviceName}}Config.{{camelCase .Name}}, "{{$serviceName}}-{{.Name}}", "", "")
	{{- end}}
{{- end}}
	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}
	errs := make(map[string][]error)
{{range .Services}}
	{{- $serviceName := .Name }}
	{{- range .Config}}
	if {{$serviceName}}Config.{{camelCase .Name}} == "" {
		errs["{{$serviceName}}"] = append(errs["{{$serviceName}}"], errors.New("{{.Name}}"))
	}
	{{- end }}
{{- end }}
{{range .Services}}
	if len(errs["{{.Name}}"]) != 0 {
		logError(logger, "{{.Name}}", errors.New("TODO"))
	} else {
		{{.Name}}Server, err := {{.PackageName}}.NewAPIServer({{.Name}}Config)
		if err != nil {
			logError(logger, "{{.Name}}", err)
		} else {
			{{.Name}}.Register{{camelCase .FullName}}Server(server, {{.Name}}Server)
			logger.Info(&operator.ServiceRegistered{&operator.Service{Name: "{{.Name}}"}})
		}
	}
{{- end}}
	return nil
}

func logError(logger operator.Logger, service string, err error) {
	logger.Error(&operator.ServiceStartupError{
		Service: &operator.Service{
			Name: service,
		},
		Message: err.Error(),
	})
}`)
