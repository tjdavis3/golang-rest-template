
# golang-rest-template

This includes extremely simple boilerplate and example for
- openapi v3 integration(server gen, documents)
- contextual, structured logging
- access control
- metrics
- graceful shutdown
- configuration through flag, env. (no files, but viper also can load configuration from files)
- kubernetes deploy with kustomize
- ...


Prerequisites
---
For developments, you need to install these tools.
```
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

Usage
---
- ```go mod init```
- **Write Spec**
- ```go generate ./...```
- Add metrics definitions to api/metrics.go
- write your code (api/serverImpl.go).
- write DB code (modes/models.go) -- modify server.db and InitializeDB as needed
- Update README.md

```
go run cmd/api/main.go
```


---
Define kubernetes spec with kustomize, check [deploy](deploy).
```
kustomize build prod | kubectl apply -f -
```


Structure
---
```
├── api // put all handlers here.
│   ├── api.gen.go
│   ├── api.go
│   ├── config.go
│   ├── error.go
│   ├── metrics.go
│   └── serverImpl.go // could be separated out by resource
├── cmd // put any executable here
│   └── api
├── pkg // library code.
├── models
└── spec
    └── openapi.yaml

```

Libraries
---
Here are the libraries I chose. Some of them are relatively new and may not be mature compared to the competitors. I had to make multiple patches to the listed library to satisfy my use cases. But at least for me, they really helped me to simplify the process. Hope they all have more users and contributors.

- **Openapi integration for documents, client/server codegen.** [oapi-codegen](https://github.com/deepmap/oapi-codegen), [kin-openapi](https://github.com/getkin/kin-openapi)  
  API service is meaningless without document. But managing documents and code separately gets really messy when your service grows.  

  "OpenAPI Specification" defines the standard way to manage your REST API service.
  Instead of writing code first, write openapi spec(which is extremely simple) with [APIcur.IO](https://apicur.io/) first and verify your api. Then, generate code from spec. For client code generation, you can use [openapi-generator](https://github.com/OpenAPITools/openapi-generator).

- **Web Framework.** net/http with [chi](https://github.com/go-chi/chi)  
   chi has 100% compatibility with net/http. gorilla/mux is more famous for this but, as oapi-codegen only supports chi, I just use chi.

  If you use standard net/http handler, It is extremely simple to integrate with third party middlewares like [rs/cors](https://github.com/rs/cors), [zerolog](https://github.com/rs/zerolog).


- **Contextual, structured logging.** [zerolog](https://github.com/rs/zerolog)  
  zerolog offers simplest api. It also offers helper library(hlog) that can be used with standard ```http.Handler```, and ```context.Context``` integration is amazing.
  ```
  ctx := log.With().Str("component", "module").Logger().WithContext(ctx)

  // ... somewhere in your function with context
  log.Ctx(ctx).Info().Msg("hello world")
  // Output: {"component":"module","level":"info","message":"hello world"}
  ```

- **Configuration.** [spf13/viper](https://github.com/spf13/viper)  

- **Metrics.** - [Prometheus](https://github.com/prometheus/client_golang)  
  It is extremely easy to add custom metrics to your server, check [api/metrics.go](api/metrics.go).

# Documentation

* [Overview](docs/README.md)
* [Configuration](docs/config.md)
* [Metrics](docs/metrics.md)
* [API Overview](docs/api/README.md)

