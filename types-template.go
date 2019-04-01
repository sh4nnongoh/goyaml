package main

type File struct {
	Comment    string
	Package    string
	Imports    []NamedTypeValue
	Constants  []NamedTypeValue
	Vars       []NamedTypeValue
	Interfaces []Interface
	Structs    []Struct
	Methods    []Method
}

type Struct struct {
	Name    string
	Comment string
	Vars    []NamedTypeValue
}
type Interface struct {
	Name    string
	Comment string
	Methods []Method
}
type Method struct {
	Comment    string
	Name       string
	Struct     NamedTypeValue
	Body       string
	Parameters []NamedTypeValue
	Results    []NamedTypeValue
}
type NamedTypeValue struct {
	Name     string
	Type     string
	Value    string
	HasValue bool
}
