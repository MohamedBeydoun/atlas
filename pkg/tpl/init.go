package tpl

func ReadmeTemplate() []byte {
	return []byte(`# {{ .Name }}

`)
}

func GitignoreTemplate() []byte {
	return []byte(`node_modules
.nyc_output
dist/
`)
}
