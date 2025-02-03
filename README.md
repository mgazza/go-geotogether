# go-geotogether

## Overview
`go-geotogether` is a Go client library for interacting with GeoTogether's API.

## Generating the Client Code
The client code is generated using [Go Swagger](https://goswagger.io/), a toolkit for generating Go client libraries from Swagger/OpenAPI specifications.

### Prerequisites
To generate the client code, you need to have Docker installed and accessible from your system. The generation process uses the `swagger` CLI from the `quay.io/goswagger/swagger` Docker image.

### Setting Up an Alias for Swagger
To simplify usage, you can set up an alias for running Swagger with Docker. Add the following line to your `~/.zshrc` or `~/.bashrc` file:

```bash
alias swagger='docker run --rm -it --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
```

Reload your shell configuration:

```bash
source ~/.zshrc  # or source ~/.bashrc
```

### Generating the Client
With the alias in place, you can generate the client code by running:

```bash
swagger generate client -f GeoTogether-swagger.yaml
```

The generated client code will be written to the current directory.

## `generate.go`
The `generate.go` file is responsible for automating the client generation process using Go's `go:generate` directive. Here is the relevant content:

To run the code generation process, use:

```bash
go generate ./generate.go
```

This will invoke the `swagger` CLI to generate the client based on the `geotogether-swagger.yaml` file.

## Contributing
If you have any improvements, feel free to fork the repository and create a pull request. Contributions are welcome!

