package tpl

// ControllerTemplate ...
func ControllerTemplate() []byte {
	return []byte(`import { Request, Response } from "express";

const {{ .Name }}Controller = {

}

export { {{ .Name }}Controller };
`)
}
