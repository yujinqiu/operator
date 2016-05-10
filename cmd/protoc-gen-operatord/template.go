package main

import "github.com/sr/operator/generator"

var (
	servicesTemplate = generator.NewTemplate("services-gen.go",
		`// Code generated by protoc-gen-operatord
package main

import (
{{range .Services}}
	"{{.ImportPath}}"
{{end}}

	"github.com/sr/operator"
	"go.pedge.io/env"
	"google.golang.org/grpc"
)

func registerServices(
	server *grpc.Server,
	logger operator.Logger,
	instrumenter operator.Instrumenter,
	authorizer operator.Authorizer,
) {
{{range .Services}}
	{{.Name}}Config := &{{.PackageName}}.Env{}
	if err := env.Populate({{.Name}}Config); err != nil {
		logError(logger, "{{.Name}}", err)
	}
	{{.Name}}Server, err := {{.PackageName}}.NewAPIServer({{.Name}}Config)
	if err != nil {
		logError(logger, "{{.Name}}", err)
	}
	intercepted := &intercepted{{.PackageName}}{{.FullName}}{
		authorizer,
		instrumenter,
		{{.Name}}Server,
	}
	{{.Name}}.Register{{camelCase .FullName}}Server(server, intercepted)
	logger.Info(&operator.ServiceRegistered{&operator.Service{Name: "{{.Name}}"}})
{{end}}
}

func logError(logger operator.Logger, service string, err error) {
	logger.Error(&operator.ServiceStartupError{
		Service: &operator.Service{
			Name: service,
		},
		Message: err.Error(),
	})
}
`)

	interceptorTemplate = generator.NewTemplate("interceptor-gen.go",
		`// Code generated by protoc-gen-operatord
package main

import (
	"time"

	"github.com/sr/operator"
	"golang.org/x/net/context"

	servicepkg "{{.ImportPath}}"
)

type intercepted{{.PackageName}}{{.FullName}} struct {
	authorizer   operator.Authorizer
	instrumenter operator.Instrumenter
	server       servicepkg.{{.FullName}}Server
}

{{- range .Methods}}
// {{.Name}} intercepts the {{$.FullName}}.{{.Name}} method.
func (a *intercepted{{$.PackageName}}{{$.FullName}}) {{.Name}}(
	ctx context.Context,
	request *servicepkg.{{.Input}},
) (response *servicepkg.{{.Output}}, err error) {
	defer func(start time.Time) {
		a.instrumenter.Instrument(
			operator.NewRequest(
				request.Source,
				"{{$.PackageName}}",
				"{{.Name}}",
				"{{.Input}}",
				"{{.Output}}",
				err,
				start,
			),
		)
	}(time.Now())
	if err := a.authorizer.Authorize(request.Source); err != nil {
		return nil, err
	}
	return a.server.{{.Name}}(ctx, request)
}
{{end}}`)
)
