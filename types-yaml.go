package main

type Yaml struct {
	ApiVersion    string
	Metadata      metadata
	Specification Specification
}

type metadata struct {
	Name  string
	Label string
}

type Specification struct {
	Endpoint  []endpoint
	Transport []transport
}

type endpoint struct {
	Name   string
	Input  input
	Output []interface{}
}

type input struct {
	Method   string
	Variable []string
}

type transport struct {
	Type string
}
