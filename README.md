
# golang-rest-template

This includes extremely simple boilerplate and example for
- openapi v3 integration(server gen, documents)
- contextual, structured logging
- access control
- metrics
- graceful shutdown
- configuration through flag, env. 
- kubernetes deploy with kustomize
- ...

![Language: GO](https://img.shields.io/badge/language-golang-blue)


Usage
---
- Create a new repo from this template
- Initialize the new repo
```bash
./initialize.sh {module}
```
eg. 
```bash
./initialize.sh github.com/ringsq/newrepo
```
You can then remove intialize.sh
```bash
git rm initialize.sh
```

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


[//]: # (Remove everything from this line to the top and edit the text below)

# Project: A new Project

A short description of the project needs to go here.

![Language: GO](https://img.shields.io/badge/language-golang-blue)

## Getting Started

### Prerequisites

oapi-codegen is used to keep the code up to date with the openapi specification.

```
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

### Installation


## Useage


## Documentation

* [Overview](docs/README.md)
* [Metrics](docs/metrics.md)
* [API Overview](docs/api/README.md)

